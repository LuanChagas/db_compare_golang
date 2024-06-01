package main

import (
	"dbcompare/cli"
	"dbcompare/compare"
	"dbcompare/config"
	"dbcompare/output"
	"dbcompare/schemas"
	"dbcompare/stats"
	"dbcompare/task"
	"os"

	"github.com/fatih/color"
)

const version = "0.3"

var dadosStats stats.Stats
var contagem stats.Contagem

func main() {

	contagem.IniciarContagem()
	dadosInput, err := cli.ParseFlags()
	if err != nil {
		color.Red("✗ Erro na entrada de dados: %v", err)
		os.Exit(1)
	}
	dadosBdPrimario, err := configurarConexao(dadosInput.ConnPrimaria, dadosInput.TipoBanco, "Primária")
	if err != nil {
		color.Red("✗ Erro ao configurar a conexão Primária: %v", err)
		os.Exit(1)
	}
	color.Green("Configuração da conexão do banco Primário    ✓")
	dadosBdSecundario, err := configurarConexao(dadosInput.ConnSecundaria, dadosInput.TipoBanco, "Secundário")
	if err != nil {
		color.Red("✗ Erro ao configurar a conexão Secundário: %v", err)
		os.Exit(1)
	}
	color.Green("Configuração da conexão do banco Secundário    ✓")

	resultChan := make(chan task.ResultTask, 2)
	go task.PrepararDados(dadosBdPrimario, "Primário", &dadosStats, resultChan)
	go task.PrepararDados(dadosBdSecundario, "Secundário", &dadosStats, resultChan)
	var dadosPreparadoPrimario, dadosPreparadoSecundario schemas.DadosMap

	for i := 0; i < 2; i++ {
		result := <-resultChan
		if result.Err != nil {
			color.Red("✗ Erro: %v", result.Err)
			os.Exit(1)
		}
		if result.Tipo == "Primário" {
			dadosPreparadoPrimario = result.Dados
		} else {
			dadosPreparadoSecundario = result.Dados
		}
	}

	color.Blue("Inicio da comparação  ...")
	resultado, err := compare.Comparar(dadosPreparadoPrimario, dadosPreparadoSecundario, &dadosStats.TempoComparacao)
	if err != nil {
		color.Red("✗ Erro ao comparar: %v", err.Error())
		os.Exit(1)
	}
	color.Green("Comparação realizada    ✓")

	err = gerarStats(dadosPreparadoPrimario, dadosPreparadoSecundario, &dadosStats, &contagem)
	if err != nil {
		color.Red("✗ Erro ao criar stats dos dados: %v", err.Error())
		os.Exit(1)
	}
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

func configurarConexao(conexao string, tipoBanco string, tipoConexao string) (config.ConfiguracaoDB, error) {
	dadosBD := config.ConfiguracaoDB{
		StringConexao: conexao,
		TipoBanco:     tipoBanco,
	}
	color.Blue("Configurando da conexão do banco %s ...", tipoConexao)
	if err := dadosBD.StringConexaoToConfiguracaoDB(); err != nil {
		return config.ConfiguracaoDB{}, err
	}
	color.Green("Configuração da conexão do banco %s    ✓", tipoConexao)
	return dadosBD, nil
}

func gerarStats(dadosPreparadoPrimario, dadosPreparadoSecundario schemas.DadosMap, dadosStats *stats.Stats, contagem *stats.Contagem) error {
	color.Blue("Gerando Stats  ...")
	var err error
	dadosStats.TamanhoDadosPrimarios, err = stats.CalcularTamanho(dadosPreparadoPrimario)
	if err != nil {
		return err
	}
	dadosStats.TamanhoDadosSecundarios, err = stats.CalcularTamanho(dadosPreparadoSecundario)
	if err != nil {
		return err
	}
	color.Green("Stats gerado    ✓")
	dadosStats.TempoGeral = contagem.TempoDecorrido()
	return nil

}
