package task

import (
	"dbcompare/compare"
	"dbcompare/config"
	"dbcompare/databases"
	"dbcompare/schemas"
	"dbcompare/stats"
	"fmt"

	"github.com/fatih/color"
)

type ResultTask struct {
	Tipo  string
	Dados schemas.DadosMap
	Err   error
}

func PrepararDados(dadosBD config.ConfiguracaoDB, tipo string, dadosStats *stats.Stats, resultChan chan<- ResultTask) {
	color.Blue("Conectando banco %s ...", tipo)
	conn, err := databases.Conectar(dadosBD)
	if err != nil {
		resultChan <- ResultTask{Tipo: tipo, Err: fmt.Errorf("✗ Erro ao conectar a conexão  %s: %v", tipo, err)}
		return
	}
	color.Green("Conexão do banco  %s    ✓", tipo)
	defer conn.Close()
	color.Blue("Buscando dados do banco %s ...", tipo)
	var tempoProcessamento *string
	if tipo == "Primário" {
		tempoProcessamento = &dadosStats.TempoProcessamentoDadosPrimarios
	} else {
		tempoProcessamento = &dadosStats.TempoProcessamentoDadosSecundarios
	}

	dadosPreparado, err := compare.PrepararDadosMysql(conn, dadosBD, tempoProcessamento)
	if err != nil {
		resultChan <- ResultTask{Tipo: tipo, Err: fmt.Errorf("✗ Erro ao preparar dados %s: %v", tipo, err)}
		return
	}
	color.Green("Dados do banco %s buscados    ✓", tipo)

	resultChan <- ResultTask{Tipo: tipo, Dados: dadosPreparado}
}
