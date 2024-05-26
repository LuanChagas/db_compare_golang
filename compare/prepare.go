package compare

import (
	"database/sql"
	"dbcompare/config"
	"dbcompare/databases"
	"dbcompare/schemas"
	"log"
)

func PrepararDadosMysql(conn *sql.DB, configuracao config.ConfiguracaoDB) (schemas.DadosMap, error) {
	mapTabelas := make(schemas.MapTabelas)
	if err := agruparDadosTabela(mapTabelas, conn, configuracao); err != nil {
		return schemas.DadosMap{}, err
	}
	if err := agruparDadosColuna(mapTabelas, conn, configuracao); err != nil {
		return schemas.DadosMap{}, err
	}
	if err := agruparDadosChaves(mapTabelas, conn, configuracao); err != nil {
		return schemas.DadosMap{}, err
	}
	mapDados := schemas.DadosMap{
		Tabelas: mapTabelas,
	}
	return mapDados, nil
}

func agruparDadosTabela(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) error {
	dadosTabelas, err := databases.BuscarTabelas(conn, configuracao)
	if err != nil {
		log.Fatal(err)
	}
	for _, valor := range dadosTabelas {
		mapTabelas[valor.Tabela] = schemas.DadosCompareMysql{
			Engine:    valor.Engine.String,
			Collation: valor.Collation.String,
			Colunas:   make(map[string]schemas.DadosColunasMysql),
			Chaves:    make(map[string]schemas.DadosChavesMysql),
		}

	}
	return nil
}

func agruparDadosColuna(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) error {
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
	return nil
}

func agruparDadosChaves(mapTabelas schemas.MapTabelas, conn *sql.DB, configuracao config.ConfiguracaoDB) error {
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
	return nil
}
