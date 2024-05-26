package compare

import (
	"dbcompare/schemas"
	"fmt"
)

type ResultadoGeral struct {
	Tabelas         ResultadoCompareTabela
	BancoPrimario   []ResultadoCompareCamposTabela
	BancoSecundario []ResultadoCompareCamposTabela
}

type ResultadoCompareTabela struct {
	Tabela map[string]ExisteTabela
}

type ExisteTabela struct {
	Primario   bool `json:"primario"`
	Secundario bool `json:"secundario"`
}

type ResultadoCompareCamposTabela struct {
	Tabela       string
	Diferenca    bool
	TabelaExiste bool
	DadosSchema  []CompareCampos
	Colunas      []CompareCampos
	Chaves       []CompareCampos
}

type CompareCampos struct {
	Nome           string
	TipoComparacao string
	Primario       string
	Secundario     string
}

func Comparar(primario schemas.DadosMap, secundario schemas.DadosMap) (ResultadoGeral, error) {

	if primario.Tabelas == nil || secundario.Tabelas == nil {
		return ResultadoGeral{}, fmt.Errorf("Dados primarios ou secundarios estão vazios")
	}
	resultadoTabelas := compareTabelas(primario.Tabelas, secundario.Tabelas)
	resultadoCamposTabelas := compareLoop(primario.Tabelas, secundario.Tabelas)
	resultadoCamposTabelasSecundaria := compareLoop(secundario.Tabelas, primario.Tabelas)

	return ResultadoGeral{
		Tabelas:         resultadoTabelas,
		BancoPrimario:   resultadoCamposTabelas,
		BancoSecundario: resultadoCamposTabelasSecundaria,
	}, nil
}

func compareTabelas(primario schemas.MapTabelas, secundario schemas.MapTabelas) ResultadoCompareTabela {
	resultadoCompareTabela := ResultadoCompareTabela{
		Tabela: make(map[string]ExisteTabela),
	}
	existeTabelas(primario, resultadoCompareTabela, "primario")
	existeTabelas(secundario, resultadoCompareTabela, "secundario")
	return resultadoCompareTabela
}

func existeTabelas(dados schemas.MapTabelas, mapTabelas ResultadoCompareTabela, banco string) {

	for tabela := range dados {
		if _, ok := mapTabelas.Tabela[tabela]; !ok {
			mapTabelas.Tabela[tabela] = ExisteTabela{}
		}
		mapTabela := mapTabelas.Tabela[tabela]
		if banco == "primario" {
			mapTabela.Primario = true
		} else {
			mapTabela.Secundario = true
		}
		mapTabelas.Tabela[tabela] = mapTabela
	}
}

func compareLoop(primario schemas.MapTabelas, secundario schemas.MapTabelas) []ResultadoCompareCamposTabela {
	var resultado []ResultadoCompareCamposTabela
	var resultadoTemp ResultadoCompareCamposTabela
	for tabela, campo := range primario {
		resultadoTemp = ResultadoCompareCamposTabela{Tabela: tabela}
		if _, ok := secundario[tabela]; !ok {
			resultadoTemp.Diferenca = true
			resultadoTemp.TabelaExiste = false
			resultado = append(resultado, resultadoTemp)
			continue
		}
		resultadoTemp.TabelaExiste = true
		resultadoTemp.DadosSchema, resultadoTemp.Diferenca = compareDadosTabela(campo, secundario[tabela])
		resultadoTemp.Colunas = compareColunas(campo.Colunas, secundario[tabela].Colunas, &resultadoTemp.Diferenca)
		resultadoTemp.Chaves = compareChaves(campo.Chaves, secundario[tabela].Chaves, &resultadoTemp.Diferenca)
		resultado = append(resultado, resultadoTemp)
	}

	return resultado
}

func compareDadosTabela(campoPrimario, campoSecundario schemas.DadosCompareMysql) ([]CompareCampos, bool) {
	compareCamposTemp := []CompareCampos{}
	diferenca := false

	comparacaoDeCampos := func(nome, original, diferente string) {
		if original != diferente {
			diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:       nome,
				Primario:   original,
				Secundario: diferente,
			})
		}
	}
	comparacaoDeCampos("Collation", campoPrimario.Collation, campoSecundario.Collation)
	comparacaoDeCampos("Engine", campoPrimario.Engine, campoSecundario.Engine)

	return compareCamposTemp, diferenca
}

func compareColunas(camposPrimario, secundario map[string]schemas.DadosColunasMysql, diferenca *bool) []CompareCampos {
	var compareCamposTemp []CompareCampos
	comparacaoDeCampos := func(nome, original, diferente, campoPrimario string) {
		if original != diferente {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:           campoPrimario,
				TipoComparacao: nome,
				Primario:       original,
				Secundario:     diferente,
			})
		}
	}
	for campoPrimario, coluna := range camposPrimario {
		if _, ok := secundario[campoPrimario]; !ok {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:           campoPrimario,
				TipoComparacao: "Campo",
				Primario:       campoPrimario,
				Secundario:     "Não existe",
			})
			continue
		}

		coluna2 := secundario[campoPrimario]
		comparacaoDeCampos("Caracteres", coluna.Caracteres.String, coluna2.Caracteres.String, campoPrimario)
		comparacaoDeCampos("Default", coluna.ValorDefault.String, coluna2.ValorDefault.String, campoPrimario)
		comparacaoDeCampos("Nulo", coluna.Nulo, coluna2.Nulo, campoPrimario)
		comparacaoDeCampos("Tipo do Campo", coluna.TipoCampo, coluna2.TipoCampo, campoPrimario)
		comparacaoDeCampos("Collation", coluna.Collation.String, coluna2.Collation.String, campoPrimario)
		comparacaoDeCampos("Extra", coluna.Extra.String, coluna2.Extra.String, campoPrimario)

	}
	return compareCamposTemp
}
func compareChaves(camposPrimario, secundario map[string]schemas.DadosChavesMysql, diferenca *bool) []CompareCampos {
	var compareCamposTemp []CompareCampos
	comparacaoDeCampos := func(nome, original, diferente, campoPrimario string) {
		if original != diferente {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:           campoPrimario,
				TipoComparacao: nome,
				Primario:       original,
				Secundario:     diferente,
			})
		}
	}
	for campoPrimario, chave := range camposPrimario {
		if _, ok := secundario[campoPrimario]; !ok {
			*diferenca = true
			compareCamposTemp = append(compareCamposTemp, CompareCampos{
				Nome:           campoPrimario,
				TipoComparacao: "Chave",
				Primario:       campoPrimario,
				Secundario:     "Não existe",
			})
			continue
		}

		colunaSecundaria := secundario[campoPrimario]
		comparacaoDeCampos("Tipo", chave.Tipo, colunaSecundaria.Tipo, campoPrimario)
		comparacaoDeCampos("Referencia", chave.Referencia.String, colunaSecundaria.Referencia.String, campoPrimario)

	}
	return compareCamposTemp
}
