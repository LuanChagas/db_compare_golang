package databases

import (
	"database/sql"
	"dbcompare/schemas"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func BuscarTabelasPostgres(conn *sql.DB, banco string) ([]schemas.DadosSchemaTabelaMysql, error) {

	rows, err := conn.Query("select table_name from information_schema.tables where table_schema = 'public' and table_catalog = $1", banco)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dados := []schemas.DadosSchemaTabelaMysql{}
	for rows.Next() {
		queryDados := schemas.DadosSchemaTabelaMysql{}
		if err := rows.Scan(&queryDados.Tabela); err != nil {

			return nil, err
		}
		dados = append(dados, queryDados)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return dados, nil
}

func BuscarColunasPostgres(conn *sql.DB, banco string) ([]schemas.DadosColunasMysql, error) {
	rows, err := conn.Query(`
	select 
	table_name, column_name, column_default,is_nullable,
	data_type,character_maximum_length,collation_name
	from information_schema.columns where table_schema = 'public' and table_catalog = $1`, banco)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dados := []schemas.DadosColunasMysql{}

	for rows.Next() {
		queryDados := schemas.DadosColunasMysql{}

		if err := rows.Scan(&queryDados.Tabela, &queryDados.Campo,
			&queryDados.ValorDefault, &queryDados.Nulo, &queryDados.TipoCampo, &queryDados.Caracteres,
			&queryDados.Collation); err != nil {
			return nil, err
		}
		dados = append(dados, queryDados)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dados, nil

}

func BuscarChavesPostgres(conn *sql.DB, banco string) ([]schemas.DadosChavesMysql, error) {
	rows, err := conn.Query(`
	select 
	k.table_name, k.column_name, k.constraint_name,cu.table_name
	from information_schema.key_column_usage as k
    JOIN information_schema.constraint_column_usage as cu on cu.CONSTRAINT_NAME =  k.constraint_name
    where k.table_schema = 'public' and  k.table_catalog = $1`, banco)
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
