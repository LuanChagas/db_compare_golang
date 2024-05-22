package main

import (
	"dbcompare/compare"
	"dbcompare/config"
	"dbcompare/databases"
	"log"
)

func main() {
	configuracaoBd1 := config.ConfiguracaoDB{
		Host:    "localhost",
		Porta:   "3306",
		Usuario: "root",
		Senha:   "123456",
		Banco:   "dbcompare_mysql1",
	}
	configuracaoBd2 := config.ConfiguracaoDB{
		Host:    "localhost",
		Porta:   "3306",
		Usuario: "root",
		Senha:   "123456",
		Banco:   "dbcompare_mysql2",
	}
	conn1, err := databases.ConectarMysql(configuracaoBd1)

	if err != nil {
		log.Fatalf("Erro:%v", err)
	}

	conn2, err := databases.ConectarMysql(configuracaoBd2)

	if err != nil {
		log.Fatalf("Erro:%v", err)
	}

	dadosPreparado1 := compare.PrepararDadosMysql(conn1, configuracaoBd1)
	dadosPreparado2 := compare.PrepararDadosMysql(conn2, configuracaoBd2)
	compare.Comparar(dadosPreparado1, dadosPreparado2)

}
