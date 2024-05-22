package schemas

import "database/sql"

type MapTabelas map[string]DadosCompareMysql

type DadosCompareMysql struct {
	Engine    string
	Collation string
	Colunas   map[string]DadosColunasMysql
	Chaves    map[string]DadosChavesMysql
	Views     map[string]DadosViewMysql
}

type DadosSchemaTabelaMysql struct {
	Tabela    string
	Engine    string
	Collation string
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

type DadosViewMysql struct {
	Tabela string
	View   string
}
