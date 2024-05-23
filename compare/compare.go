package compare

import (
	"dbcompare/schemas"
	"fmt"
)

type ResultadoCompareTabela struct {
	Tabela       string
	Diferenca    bool
	TabelaExiste bool
	DadosSchema  []CompareCampos
	Colunas      []CompareCampos
	Chaves       []CompareCampos
	Views        []CompareCampos
}

type ResultadoCompareViews struct {
	Tabela      string
	IsDiferenca bool
	View        string
	Diferenca   string
}

type CompareCampos struct {
	Nome      string
	Original  string
	Diferenca string
}

func Comparar(dados1 schemas.DadosMap, dados2 schemas.DadosMap) {
	resultadoTabelas := compareLoop(dados1.Tabelas, dados2.Tabelas)
	resultadoViews := compareViews(dados1.Views, dados2.Views)
	fmt.Printf("%v", resultadoTabelas)
	fmt.Printf("%v", resultadoViews)
}

func compareLoop(dados1 schemas.MapTabelas, dados2 schemas.MapTabelas) []ResultadoCompareTabela {
	var resultado []ResultadoCompareTabela
	var resultadoTemp ResultadoCompareTabela
	for tabela, campo := range dados1 {
		resultadoTemp = ResultadoCompareTabela{Tabela: tabela}
		if _, ok := dados2[tabela]; !ok {
			resultadoTemp.Diferenca = true
			resultadoTemp.TabelaExiste = false
			resultado = append(resultado, resultadoTemp)
			continue
		}
		resultadoTemp.TabelaExiste = true
		resultadoTemp.DadosSchema, resultadoTemp.Diferenca = compareDadosTabela(campo, dados2, tabela)
		resultadoTemp.Colunas = compareColunas(campo.Colunas, dados2, tabela, &resultadoTemp.Diferenca)
		resultadoTemp.Colunas = compareChaves(campo.Chaves, dados2, tabela, &resultadoTemp.Diferenca)
		resultado = append(resultado, resultadoTemp)
	}
	return resultado
}

func compareDadosTabela(campo schemas.DadosCompareMysql, dados2 schemas.MapTabelas, tabela string) ([]CompareCampos, bool) {
	compareCamposTemp := []CompareCampos{}
	diferenca := false

	if campo.Collation != dados2[tabela].Collation {
		diferenca = true
		compareCamposTemp = append(compareCamposTemp, CompareCampos{
			Nome:      "Collation",
			Original:  campo.Collation,
			Diferenca: dados2[tabela].Collation,
		})
	}
	if campo.Engine != dados2[tabela].Engine {
		diferenca = true
		compareCamposTemp = append(compareCamposTemp, CompareCampos{
			Nome:      "Engine",
			Original:  campo.Engine,
			Diferenca: dados2[tabela].Engine,
		})
	}

	return compareCamposTemp, diferenca
}

func compareColunas(campos map[string]schemas.DadosColunasMysql, dados2 schemas.MapTabelas, tabela string, diferenca *bool) []CompareCampos {
	var compareCamposTemp []CompareCampos
	comparacaoDeCampos := func(nome, original, diferente string) {
		if original != diferente {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:      nome,
				Original:  original,
				Diferenca: diferente,
			})
		}
	}
	for campo, coluna := range campos {
		if _, ok := dados2[tabela].Colunas[campo]; !ok {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:      "Campo",
				Original:  campo,
				Diferenca: "Não existe",
			})
			continue
		}

		coluna2 := dados2[tabela].Colunas[campo]
		comparacaoDeCampos("Caracteres", coluna.Caracteres.String, coluna2.Caracteres.String)
		comparacaoDeCampos("Default", coluna.ValorDefault.String, coluna2.ValorDefault.String)
		comparacaoDeCampos("Nulo", coluna.Nulo, coluna2.Nulo)
		comparacaoDeCampos("Tipo do Campo", coluna.TipoCampo, coluna2.TipoCampo)
		comparacaoDeCampos("Collation", coluna.Collation.String, coluna2.Collation.String)
		comparacaoDeCampos("Extra", coluna.Extra.String, coluna2.Extra.String)

	}
	return compareCamposTemp
}
func compareChaves(campos map[string]schemas.DadosChavesMysql, dados2 schemas.MapTabelas, tabela string, diferenca *bool) []CompareCampos {
	var compareCamposTemp []CompareCampos
	comparacaoDeCampos := func(nome, original, diferente string) {
		if original != diferente {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:      nome,
				Original:  original,
				Diferenca: diferente,
			})
		}
	}
	for campo, chave := range campos {
		if _, ok := dados2[tabela].Chaves[campo]; !ok {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:      "Campo",
				Original:  campo,
				Diferenca: "Não existe",
			})
			continue
		}

		coluna2 := dados2[tabela].Chaves[campo]
		comparacaoDeCampos("Tipo", chave.Tipo, coluna2.Tipo)
		comparacaoDeCampos("Referencia", chave.Referencia.String, coluna2.Referencia.String)

	}
	return compareCamposTemp
}

func compareViews(views schemas.MapViews, dados2 schemas.MapViews) []ResultadoCompareViews {
	var resultado []ResultadoCompareViews
	for tabela, view := range views {
		if _, ok := dados2[tabela]; !ok {
			resultado = append(resultado, ResultadoCompareViews{
				Tabela:      tabela,
				IsDiferenca: true,
				View:        view,
				Diferenca:   "Não existe",
			})
			continue
		}

		if view != dados2[tabela] {
			resultado = append(resultado, ResultadoCompareViews{
				Tabela:      tabela,
				IsDiferenca: true,
				View:        view,
				Diferenca:   dados2[tabela],
			})
		}

	}
	return resultado
}
