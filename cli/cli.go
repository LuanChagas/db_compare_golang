package cli

import (
	"flag"
	"fmt"
)

type DadosInput struct {
	ConnPrimaria   string
	ConnSecundaria string
}

func ParseFlags() (*DadosInput, error) {
	var connPrimaria, connSecundaria string
	flag.StringVar(&connPrimaria, "connPrimaria", "", "Conexão primaria")
	flag.StringVar(&connSecundaria, "connSecundaria", "", "Conexão Secundaria")

	flag.Parse()

	if connPrimaria == "" {
		return nil, fmt.Errorf("dados da conexão primária não foi passada")
	}

	if connSecundaria == "" {
		return nil, fmt.Errorf("dados da conexão secundária não foi passada")
	}

	return &DadosInput{
		ConnPrimaria:   connPrimaria,
		ConnSecundaria: connSecundaria,
	}, nil
}
