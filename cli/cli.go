package cli

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

type DadosInput struct {
	ConnPrimaria   string
	ConnSecundaria string
	TipoBanco      string
}

func ParseFlags() (*DadosInput, error) {
	var connPrimaria, connSecundaria, tipoBanco string
	var mysql, postgres bool
	flag.StringVar(&connPrimaria, "connPrimaria", "",
		"Conexão primaria Example.: MYSQL-> \"--connPrimaria root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1\" , POSTGRESQL -> \"--connPrimaria user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1\"")
	flag.StringVar(&connSecundaria, "connSecundaria",
		"",
		"Conexão secundária Example.: MYSQL-> \"--connSecundaria root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1\" , POSTGRESQL -> \"--connSecundaria user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1\"")
	flag.BoolVar(&mysql, "mysql", false, "Banco do tipo mysql Example.: --mysql")
	flag.BoolVar(&postgres, "postgres", false, "Banco do tipo Postgresql Example.: --postgres")
	flag.Usage = func() {
		fmt.Fprintf(color.Output, "%s:\n", color.YellowString("Como usar:"))
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(color.Output, "  -%s\n", color.GreenString(f.Name))
			fmt.Fprintf(color.Output, "        %s\n", color.BlueString(f.Usage))
		})
	}
	flag.Parse()

	if !mysql && !postgres {
		return nil, fmt.Errorf("tipo de banco não foi passado nos argumentos")
	}

	if mysql && postgres {
		return nil, fmt.Errorf("escolha apenas um tipo de banco nos argumentos")
	}

	if mysql {
		tipoBanco = "MYSQL"
	}

	if postgres {
		tipoBanco = "POSTGRESQL"
	}

	if connPrimaria == "" {
		return nil, fmt.Errorf("dados da conexão primária não foi passada")
	}

	if connSecundaria == "" {
		return nil, fmt.Errorf("dados da conexão secundária não foi passada")
	}

	return &DadosInput{
		ConnPrimaria:   connPrimaria,
		ConnSecundaria: connSecundaria,
		TipoBanco:      tipoBanco,
	}, nil
}
