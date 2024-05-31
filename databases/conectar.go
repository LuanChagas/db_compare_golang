package databases

import (
	"database/sql"
	"dbcompare/config"
	"dbcompare/schemas"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Conectar(configuracao config.ConfiguracaoDB) (*sql.DB, error) {

	var db *sql.DB
	var err error

	switch configuracao.TipoBanco {
	case "MYSQL":
		db, err = sql.Open("mysql", configuracao.StringConexao)
	case "POSTGRESQL":
		db, err = sql.Open("postgres", configuracao.StringConexao)
	default:
		return nil, fmt.Errorf("tipo de banco de dados desconhecido: %s", configuracao.TipoBanco)
	}

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func BuscarTabelas(conn *sql.DB, dadosBanco config.ConfiguracaoDB) ([]schemas.DadosSchemaTabelaMysql, error) {
	switch dadosBanco.TipoBanco {
	case "MYSQL":
		return BuscarTabelasMysql(conn, dadosBanco.Banco)
	case "POSTGRESQL":
		return BuscarTabelasPostgres(conn, dadosBanco.Banco)
	default:
		return nil, fmt.Errorf("tipo de banco de dados desconhecido: %s", dadosBanco.TipoBanco)
	}
}

func BuscarColunas(conn *sql.DB, dadosBanco config.ConfiguracaoDB) ([]schemas.DadosColunasMysql, error) {
	switch dadosBanco.TipoBanco {
	case "MYSQL":
		return BuscarColunasMysql(conn, dadosBanco.Banco)
	case "POSTGRESQL":
		return BuscarColunasPostgres(conn, dadosBanco.Banco)
	default:
		return nil, fmt.Errorf("tipo de banco de dados desconhecido: %s", dadosBanco.TipoBanco)
	}
}

func BuscarChaves(conn *sql.DB, dadosBanco config.ConfiguracaoDB) ([]schemas.DadosChavesMysql, error) {
	switch dadosBanco.TipoBanco {
	case "MYSQL":
		return BuscarChavesMysql(conn, dadosBanco.Banco)
	case "POSTGRESQL":
		return BuscarChavesPostgres(conn, dadosBanco.Banco)
	default:
		return nil, fmt.Errorf("tipo de banco de dados desconhecido: %s", dadosBanco.TipoBanco)
	}
}
