package stats

import (
	"dbcompare/schemas"
	"encoding/json"
	"fmt"
	"time"
)

type Stats struct {
	TempoGeral                         string
	TempoProcessamentoDadosPrimarios   string
	TempoProcessamentoDadosSecundarios string
	TempoComparacao                    string
	TamanhoDadosPrimarios              string
	TamanhoDadosSecundarios            string
}

type Contagem struct {
	Inicio time.Time
}

func (contagem *Contagem) IniciarContagem() {
	contagem.Inicio = time.Now()
}

func (contagem *Contagem) TempoDecorrido() string {
	tempo := time.Since(contagem.Inicio)
	return fmt.Sprintf("%d ms", tempo.Milliseconds())
}

func CalcularTamanho(dados schemas.DadosMap) (string, error) {
	bytes, err := json.Marshal(dados)
	if err != nil {
		return "", err
	}
	tamanho := len(bytes)
	if tamanho > 1024 {
		return fmt.Sprintf("%.2f KB", float64(tamanho/1024)), nil
	}
	return fmt.Sprintf("%d bytes", tamanho), nil
}
