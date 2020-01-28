# Aplicação de exemplo Aluguel de Ferramentas

Descrição da aplicação de exemplo da apresentação Escalando Testes de Integração da Aplicação com Docker.

- [O Problema](#O-Problema)
- [Dependências](#Dependências)
- [Ambiente Desenvolvimento](#Ambiente-Desenvolvimento)
    - [Dependências Desenvolvimento](#Dependências-Desenvolvimento)
    - [Configurações](#Configurações)
    - [Infra Testes](#Infra-Testes)
- [Testes](#Testes)
    - [Unitários](#Unitários)
    - [Integração](#Integração)
    - [Unitários e Integração](#Unitários-e-Integração)
- [Comandos Make](#comandos-make)

## O Problema

Verificar na descrição do [Projeto Aluguel de Ferramentas](../README.md).

Exemplificamos como executar testes de integração em paralelo, mesmo que envolvam serviços externos, como Bases de Dados.

As técnicas apresentadas também podem ser usadas para executar testes em paralelo de aplicações que integragem com mensageria [AMQP(RabbitMQ)](https://github.com/kimprado/rabbeasy/blob/951f41960bd92149b1ec100b2ef6426c2389c27e/pkg/rabbeasy/clientsIT_test.go#L11).

Definimos as seguintes estratégias para diferentes tipos de Banco de Dados.

- **Postgres**

    Não exige alterações na camada Repository da aplicação.

    Definimos um esquema com dados(estado) padrão em forma de [template](https://www.postgresql.org/docs/9.4/manage-ag-templatedbs.html). Cada teste cria seu próprio esquema baseado em seu nome.

- **MongoDB**

    Não exige alterações na camada Repository da aplicação.

    Criamos um "esquema" com seu respectivo estado para cada teste na inicialização do container.

- **Redis**

    Exige alterações na camada Repository da aplicação que faz comunicação com Redis.

    As chaves do Redis passam a ter um prefixo parametrizado por configuração. Cada teste personaliza este prefixo com seu próprio nome.

    A parametrização de configuração para cada teste é possível por meio de Injeção de Dependências.

## Dependências

- **[Go](http://golang.org/)** >= 1.13 - Liguagem usada na implementação da API. 
- **[Wire](http://github.com/google/wire)** >= 0.4.0 - Framework de Injeção de Dependências.
- **[mgo](http://gopkg.in/mgo.v2)** - Popular driver de MongoDB.
- **[pgx](http://github.com/jackc/pgx)** - Driver performárico de Postgres.
- **[Redigo](http://github.com/gomodule/redigo)** - Driver performárico, que mantém API do Redis.
- **[slLog](http://github.com/kimprado/sllog)** - Escrevi esta lib para configurar logging como no Spring Boot.
- **[Configor](http://github.com/jinzhu/configor)** - Lib flexível para carregar configuração, via Variáveis de Ambiente e outros.
- **[HttpRouter](http://github.com/julienschmidt/httprouter)** - HTTP mux performático e flexível.
- **[Testify](http://github.com/stretchr/testify)** - Lib que simplifica assertions.

## Ambiente Desenvolvimento

Segue descrição das ferramentas e configurações do ambiente de desenvolvimento.

### Dependências Desenvolvimento

As seguintes ferramentas devem ser instaladas.

- **Docker** - Ferramenta usada para containerização.
- **Docker Compose** - Ferramenta usada para orquestração em ambiende de dev.
- **Go** >= 1.13 - Linguagem de programação.
- **Wire** >= 0.4.0 - Framework de Injeção de Dependências(provisionado pelo script [`configure`](configure))

Execute script [`configure`](configure) presente na raiz do repositório. 

O script cria configuração de IDE.

```sh
./configure
```

Os seguintes arquivos são criados, caso necessário.

- .vscode/settings.json - Arquivo da IDE VSCode
- .vscode/launch.json - Arquivo da IDE VSCode

### Configurações

Para execução dos testes são usados os seguintes arquivos:
 
 - [`config-integration-container.json`](configs/config-integration-container.json) - Execução de testes em ambiente containerizado.

 - [`config-integration.json`](configs/config-integration.json) - Execução de testes em ambiente local.

### Infra Testes

- Iniciar infra de Testes

    ```sh
    make infra-test-start
    ```

- Interromper infra de Testes

    ```sh
    make infra-test-stop
    ```

## Testes

Estão dividos em *[Unitários](#Unitários)* e *[Integração](#Integração)*. Os grupos de testes são separados em arquivos de testes diferentes. Usei o conceito de [Build Constraints ou Build Tag](http://golang.org/pkg/go/build/#hdr-Build_Constraints) para selecionar quais testes queremos executar.

Para especificar um grupo de teste executamos o comando `go test` com o parâmetro `-tags`.

### Unitários

Testes unitários que não dependem da infra para executar, são mais rápidos, podendo conter Mock Objects conforme necessário.

Use os seguintes comandos para executar os testes unitários.

- Ambiente Containerizado.

    ```sh
    make test-unit-container
    ```

- Ambiente Local.

    ```sh
    make test-unit
    ```

### Integração

Testes de integração dependem do [deploy da infra de testes](#Infra-Testes). Acessam os serviços de dependência sem Mock Objects. 

Procuramos acelerar sua execução habilitando o paralelismo com *`t.Parallel()`*.

- Ambiente Containerizado.

    ```sh
    make test-integration-container
    ```

- Ambiente Local.

    ```sh
    make test-integration
    ```

### Unitários e Integração

Permite executar ao mesmo tempo testes de [Unidade](#Unitários) e de [Integração](#Integração). O benefício é ter maior cobertura e estatística unificada.

Use os seguintes comandos para executar os testes.

- Ambiente Containerizado.

    ```sh
    make test-all-container
    ```

- Ambiente Local.

    ```sh
    make test-all
    ```
## Comandos Make

Todos comandos para facilitar o desenvolvimento estão no [Makefile](Makefile).

Para listar comandos disponíveis use o seguinte comando.

```sh
make help
```