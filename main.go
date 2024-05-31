package main

import (
	"dbcompare/cli"
	"dbcompare/compare"
	"dbcompare/config"
	"dbcompare/databases"
	"dbcompare/output"
	"dbcompare/stats"
	"os"

	"github.com/fatih/color"
)

const version = "0.2"

var dadosStats stats.Stats
var contagem stats.Contagem

func main() {

	contagem.IniciarContagem()
	dadosInput, err := cli.ParseFlags()
	if err != nil {
		color.Red("✗ Erro na entrada de dados: %v", err)
		os.Exit(1)
	}

	dadosBdPrimario := config.ConfiguracaoDB{
		StringConexao: dadosInput.ConnPrimaria,
		TipoBanco:     dadosInput.TipoBanco,
	}

	dadosBdSecundario := config.ConfiguracaoDB{
		StringConexao: dadosInput.ConnSecundaria,
		TipoBanco:     dadosInput.TipoBanco,
	}

	color.Blue("Configurando da conexão do banco Primário ...")
	if err := dadosBdPrimario.StringConexaoToConfiguracaoDB(); err != nil {
		color.Red("✗ Erro ao configurar a conexão primária: %v", err)
		return
	}
	color.Green("Configuração da conexão do banco Primário    ✓")

	color.Blue("Configurando conexão do banco Secundário ...")
	if err := dadosBdSecundario.StringConexaoToConfiguracaoDB(); err != nil {
		color.Red("✗ Erro ao configurar a conexão secundária: %v", err)
		os.Exit(1)
	}
	color.Green("Configuração da conexão do banco Secundário    ✓")

	color.Blue("Conectando banco Primário ...")
	connPrimaria, err := databases.Conectar(dadosBdPrimario)
	if err != nil {
		color.Red("✗ Erro ao conectar a conexão primária: %v", err)
		os.Exit(1)
	}
	color.Green("Conexão do banco Primário    ✓")
	defer connPrimaria.Close()

	color.Blue("Conectando banco Secundário ...")
	connSecundaria, err := databases.Conectar(dadosBdSecundario)
	if err != nil {
		color.Red("✗ Erro ao conectar a conexão secundaria: %v", err)
		os.Exit(1)
	}
	color.Green("Conexão do banco Secundário    ✓")
	defer connSecundaria.Close()
	color.Blue("Buscando dados do banco Primário ...")
	dadosPreparadoPrimario, err := compare.PrepararDadosMysql(connPrimaria, dadosBdPrimario, &dadosStats.TempoProcessamentoDadosPrimarios)
	if err != nil {
		color.Red("✗ Erro ao preparar dados primários: %v", err)
		os.Exit(1)
	}
	color.Green("Dados do banco Primário buscados    ✓")

	color.Blue("Buscando dados do banco Secundário ...")
	dadosPreparadoSecundario, err := compare.PrepararDadosMysql(connSecundaria, dadosBdSecundario, &dadosStats.TempoProcessamentoDadosSecundarios)
	if err != nil {
		color.Red("✗ Erro ao preparar dados secundários: %v", err)
		os.Exit(1)
	}
	color.Green("Dados do banco Secundário buscados    ✓")

	color.Blue("Inicio da comparação  ...")
	resultado, err := compare.Comparar(dadosPreparadoPrimario, dadosPreparadoSecundario, &dadosStats.TempoComparacao)
	if err != nil {
		color.Red("✗ Erro ao comparar: %v", err.Error())
		os.Exit(1)
	}
	color.Green("Comparação realizada    ✓")

	color.Blue("Gerando Stats  ...")

	dadosStats.TamanhoDadosPrimarios, err = stats.CalcularTamanho(dadosPreparadoPrimario)
	if err != nil {
		color.Red("✗ Erro ao criar stats dos dados Primários: %v", err.Error())
		os.Exit(1)
	}
	dadosStats.TamanhoDadosSecundarios, err = stats.CalcularTamanho(dadosPreparadoSecundario)
	if err != nil {
		color.Red("✗ Erro ao criar stats dos dados Secundários: %v", err.Error())
		os.Exit(1)
	}
	color.Green("Stats gerado    ✓")
	dadosStats.TempoGeral = contagem.TempoDecorrido()
	dadosOutPut := output.DadosOutput{
		Resultado:       resultado,
		BancoPrimario:   dadosBdPrimario.Banco,
		BancoSecundario: dadosBdSecundario.Banco,
		Version:         version,
		Stats:           dadosStats,
	}

	color.Blue("Gerando output HTML  ...")
	if err := output.GerarHtml(dadosOutPut); err != nil {
		color.Red("✗ Erro ao gerar output: %v", err.Error())
		os.Exit(1)
	}

}
