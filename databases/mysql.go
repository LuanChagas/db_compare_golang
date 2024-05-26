package databases

import (
	"database/sql"
	"dbcompare/config"
	"dbcompare/schemas"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarMysql(configuracao config.ConfiguracaoDB) (*sql.DB, error) {
	stringConexao := configuracao.StringConexaoMysql()

	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func BuscarTabelas(conn *sql.DB, configuracao config.ConfiguracaoDB) ([]schemas.DadosSchemaTabelaMysql, error) {
	rows, err := conn.Query("select table_name,engine, table_collation from information_schema.tables where table_schema = ?", configuracao.Banco)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dados := []schemas.DadosSchemaTabelaMysql{}
	for rows.Next() {
		queryDados := schemas.DadosSchemaTabelaMysql{}
		if err := rows.Scan(&queryDados.Tabela, &queryDados.Engine, &queryDados.Collation); err != nil {

			return nil, err
		}
		dados = append(dados, queryDados)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("ola")
		return nil, err
	}
	return dados, nil
}

func BuscarColunas(conn *sql.DB, banco string) ([]schemas.DadosColunasMysql, error) {
	rows, err := conn.Query(`
	select 
	table_name, column_name, column_default,is_nullable,
	column_type,character_set_name,collation_name,extra 
	from information_schema.columns where table_schema = ?`, banco)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dados := []schemas.DadosColunasMysql{}

	for rows.Next() {
		queryDados := schemas.DadosColunasMysql{}

		if err := rows.Scan(&queryDados.Tabela, &queryDados.Campo,
			&queryDados.ValorDefault, &queryDados.Nulo, &queryDados.TipoCampo, &queryDados.Caracteres,
			&queryDados.Collation, &queryDados.Extra); err != nil {
			return nil, err
		}
		dados = append(dados, queryDados)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dados, nil

}

func BuscarChaves(conn *sql.DB, banco string) ([]schemas.DadosChavesMysql, error) {
	rows, err := conn.Query(`
	select 
	table_name, column_name, constraint_name,referenced_table_name
	from information_schema.key_column_usage where table_schema = ?`, banco)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dados := []schemas.DadosChavesMysql{}

	for rows.Next() {
		queryDados := schemas.DadosChavesMysql{}

		if err := rows.Scan(
			&queryDados.Tabela, &queryDados.Campo, &queryDados.Tipo, &queryDados.Referencia); err != nil {
			return nil, err
		}
		dados = append(dados, queryDados)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dados, nil
}
