package config

import (
	"fmt"
	"regexp"
	"strconv"
)

type ConfiguracaoDB struct {
	Host          string
	Porta         string
	Usuario       string
	Senha         string
	Banco         string
	StringConexao string
	TipoBanco     string
}

func (config *ConfiguracaoDB) StringConexaoToConfiguracaoDB() error {

	var regex *regexp.Regexp
	if config.TipoBanco == "MYSQL" {
		regex = regexp.MustCompile(`^(.*?):(.*?)@tcp\((.*?):(.*?)\)\/(.*?)$`)
	}

	if config.TipoBanco == "POSTGRESQL" {
		regex = regexp.MustCompile(`^user=(.*?) password=(.*?) host=(.*?) port=(.*?) dbname=(.*?)( sslmode=(.*?))?$`)
	}

	matches := regex.FindStringSubmatch(config.StringConexao)
	if matches == nil || len(matches) < 6 {
		return fmt.Errorf("formato invalido da string")
	}

	if _, err := strconv.Atoi(matches[4]); err != nil {
		return fmt.Errorf("formato invalido da porta")
	}

	if err := validacaoExisteVazio(matches); err != nil {
		return err
	}

	config.Usuario = matches[1]
	config.Senha = matches[2]
	config.Host = matches[3]
	config.Porta = matches[4]
	config.Banco = matches[5]

	return nil
}

func validacaoExisteVazio(matches []string) error {

	campo := [6]string{"usuario", "senha", "host", "porta", "banco"}

	for i, value := range matches[1:6] {
		if value == "" {
			return fmt.Errorf("campo %s está vazio", campo[i])
		}
	}
	return nil
}
