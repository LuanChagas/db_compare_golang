package schemas

import "database/sql"

type MapTabelas map[string]DadosCompareMysql
type MapViews map[string]string

type DadosMap struct {
	Tabelas MapTabelas
}
type DadosCompareMysql struct {
	Engine    string
	Collation string
	Colunas   map[string]DadosColunasMysql
	Chaves    map[string]DadosChavesMysql
}

type DadosSchemaTabelaMysql struct {
	Tabela    string
	Engine    sql.NullString
	Collation sql.NullString
}

type DadosColunasMysql struct {
	Tabela       string
	Campo        string
	ValorDefault sql.NullString
	Nulo         string
	TipoCampo    string
	Caracteres   sql.NullString
	Collation    sql.NullString
	Extra        sql.NullString
}

type DadosChavesMysql struct {
	Tabela     string
	Campo      string
	Tipo       string
	Referencia sql.NullString
}
