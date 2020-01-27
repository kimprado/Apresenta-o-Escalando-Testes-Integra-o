# Apresentação Escalando Testes de Integração da Aplicação com Docker

Código fonte dos projetos de exemplo usados na apresentação.

## Objetivo

Apresentar implementação de arquitetura de aplicação que permite executar testes de integração com colaboradores reais, em paralelo, mas sem [race condition](https://en.wikipedia.org/wiki/Race_condition).

## Projeto Aluguel de Ferramentas

Projeto implementado como API HTTP.

Aplicação representa um simples e parcial processo de Aluguel de Ferramentas.

Usuários(Locatário) são capazes de alugar diversos equipamentos, conforme disponibilidade de estoque.

Temos seguinte tabela que mapeia entidades e suas respectivas bases de dados.

Entidade | Base de Dados | Objeto | Motivação
--- | ---: | :---: | :---
Locatário | Postgres | Table | Permite Guardar dados estruturados bem conhecidos
Equipamento | MongoDB | Colection |  Entidades com diveras estruturas entre si
Estoque | Redis | Key-Value | Totalização de valor


## Versões da Aplicação de Exemplo

Para fins de comparação temos duas versões de implementação do projeto. Uma versão é flexível e permite rodar testes em paralelo, outra é rígida e os testes devem ser executados um por vez.

 - [Versão concorrente](./testes-integracao-concorrente)
 - [Versão não concorrente](./testes-integracao-nao-concorrente)
