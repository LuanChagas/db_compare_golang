package compare

import (
	"database/sql"
	"dbcompare/config"
	"dbcompare/databases"
	"dbcompare/schemas"
	"log"
)

func PrepararDadosMysql(conn *sql.DB, configuracao config.ConfiguracaoDB) schemas.MapTabelas {
	mapTabelas := schemas.MapTabelas{}
	agruparDadosTabela(mapTabelas, conn, configuracao)
	agruparDadosColuna(mapTabelas, conn, configuracao)
	agruparDadosChaves(mapTabelas, conn, configuracao)
	agruparDadosViews(mapTabelas, conn, configuracao)

	return mapTabelas

}

func agruparDadosTabela(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) {
	dadosTabelas, err := databases.BuscarTabelas(conn, configuracao)
	if err != nil {
		log.Fatal(err)
	}
	for _, valor := range dadosTabelas {
		mapTabelas[valor.Tabela] = schemas.DadosCompareMysql{
			Engine:    valor.Engine,
			Collation: valor.Collation,
			Colunas:   make(map[string]schemas.DadosColunasMysql),
			Chaves:    make(map[string]schemas.DadosChavesMysql),
			Views:     make(map[string]schemas.DadosViewMysql),
		}
	}
}

func agruparDadosColuna(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) {
	dadosColunas, err := databases.BuscarColunas(conn, configuracao.Banco)
	if err != nil {
		log.Fatal(err)
	}
	for _, valor := range dadosColunas {
		if tabelaExiste, ok := mapTabelas[valor.Tabela]; ok {
			tabelaExiste.Colunas[valor.Campo] = valor
			mapTabelas[valor.Tabela] = tabelaExiste
		}
	}
}

func agruparDadosChaves(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) {
	dadosChaves, err := databases.BuscarChaves(conn, configuracao.Banco)
	if err != nil {
		log.Fatal(err)
	}
	for _, valor := range dadosChaves {
		if tabelaExiste, ok := mapTabelas[valor.Tabela]; ok {
			tabelaExiste.Chaves[valor.Campo] = valor
			mapTabelas[valor.Tabela] = tabelaExiste
		}
	}
}

func agruparDadosViews(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) {
	dadosViews, err := databases.BuscarViews(conn, configuracao.Banco)
	if err != nil {
		log.Fatal(err)
	}
	for _, valor := range dadosViews {
		if tabelaExiste, ok := mapTabelas[valor.Tabela]; ok {
			tabelaExiste.Views[valor.View] = valor
			mapTabelas[valor.Tabela] = tabelaExiste
		}
	}
}
