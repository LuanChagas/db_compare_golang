package config

import (
	"fmt"
)

type ConfiguracaoDB struct {
	Host    string
	Porta   string
	Usuario string
	Senha   string
	Banco   string
}

func (config *ConfiguracaoDB) StringConexaoMysql() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Usuario, config.Senha, config.Host, config.Porta, config.Banco)
}
