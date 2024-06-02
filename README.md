# DB Compare Golang

DB Compare é uma ferramenta desenvolvida em Go para comparar bancos de dados MySQL e PostgreSQL. Ele permite identificar diferenças entre dois bancos de dados e garantir a consistência entre eles.


## Instalação

### Pré-requisitos

- [Go](https://golang.org/doc/install) instalado.
- MySQL ou PostgreSQL configurado.

### Compilação

Clone o repositório e compile o projeto:

```bash
git clone https://github.com/seuusuario/dbcompare.git
cd dbcompare
go build -o dbcompare .
```

## Como Usar

### Argumentos

- `-connPrimaria`: Conexão primária. Exemplo:
  - MySQL: `--connPrimaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1"`
  - PostgreSQL: `--connPrimaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1"`
- `-connSecundaria`: Conexão secundária. Exemplo:
  - MySQL: `--connSecundaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1"`
  - PostgreSQL: `--connSecundaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1"`
- `-mysql`: Indica que o banco de dados é MySQL. Exemplo: `--mysql`
- `-postgres`: Indica que o banco de dados é PostgreSQL. Exemplo: `--postgres`

### Exemplos

#### No Linux

##### Banco PostgreSQL

```bash
./dbcompare --postgres \
--connPrimaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1" \
--connSecundaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres2"
```

##### Banco MySQL

```bash
./dbcompare --mysql \
--connPrimaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1" \
--connSecundaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql2"
```

#### No Windows

##### Banco PostgreSQL

```bash
./dbcompare.exe --postgres \
--connPrimaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres1" \
--connSecundaria "user=postgres password=123456 host=localhost port=5432 dbname=db_compare_postgres2"
```

##### Banco MySQL

```bash
./dbcompare.exe --mysql \
--connPrimaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql1" \
--connSecundaria "root:123456@tcp(127.0.0.1:3306)/dbcompare_mysql2"
```
#### Saída
![enter image description here](https://github.com/LuanChagas/db_compare_golang/blob/f29fb33ff9af1cb386bd1caf51c5bb09b82f90f2/inicio.png?raw=true)
![enter image description here](https://raw.githubusercontent.com/LuanChagas/db_compare_golang/f29fb33ff9af1cb386bd1caf51c5bb09b82f90f2/fim.png)
## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Para mais informações, entre em contato:

- Email: luanchagas@hotmail.com
- LinkedIn: [Luan Chagas](https://www.linkedin.com/in/luanchagas/)










