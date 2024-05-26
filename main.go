package main

import (
	"dbcompare/compare"
	"dbcompare/config"
	"dbcompare/databases"
	"dbcompare/output"
	"log"
)

func main() {
	configuracaoBdPrimaria := config.ConfiguracaoDB{
		Host:    "localhost",
		Porta:   "3306",
		Usuario: "root",
		Senha:   "123456",
		Banco:   "dbcompare_mysql1",
	}
	configuracaoBdSecundaria := config.ConfiguracaoDB{
		Host:    "localhost",
		Porta:   "3306",
		Usuario: "root",
		Senha:   "123456",
		Banco:   "dbcompare_mysql2",
	}
	connPrimaria, err := databases.ConectarMysql(configuracaoBdPrimaria)

	if err != nil {
		log.Panic(err)
	}

	connSecundaria, err := databases.ConectarMysql(configuracaoBdSecundaria)

	if err != nil {
		log.Panic(err)
	}

	dadosPreparadoPrimario, err := compare.PrepararDadosMysql(connPrimaria, configuracaoBdPrimaria)

	if err != nil {
		log.Panic(err)
	}

	dadosPreparadoSecundario, err := compare.PrepararDadosMysql(connSecundaria, configuracaoBdSecundaria)

	if err != nil {
		log.Panic(err)
	}

	resultado, err := compare.Comparar(dadosPreparadoPrimario, dadosPreparadoSecundario)
	if err != nil {
		log.Panic(err)
	}
	output.GerarHtml(resultado)

}
