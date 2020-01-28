# Self-Documenting Makefiles https://swcarpentry.github.io/make-novice/08-self-doc/index.html
## help				: Exibe comandos make disponíveis.
.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $<

## build				: Compila aplicação. Gera arquivo './equipment-rental-api.bin'.
build:
	@./scripts/compile.sh build

## wire				: Gera/Atualiza códigos(wire_gen*.go) do framework de Injeção de Dependências.
wire:
	@./scripts/compile.sh wire
	@./scripts/compile.sh wire-testes

## test-unit			: Testes de unidade.
test-unit:
	@./scripts/test.sh unit

## test-integration		: Testes de integração.
test-integration:
	@./scripts/test.sh integration

## test-all			: Executa testes de unidade e integração.
test-all:
	@./scripts/test.sh all

## test-unit-container		: Executa testes de unidade em ambiente containerizado.
test-unit-container:
	@docker-compose up --build test-unit

## test-integration-container	: Executa testes de integração em ambiente containerizado.
test-integration-container: infra-test-stop
	@docker-compose up --build test-integration

## test-all-container		: Executa testes de unidade e integração em ambiente containerizado.
test-all-container:
	@docker-compose up --build test-all

## infra-test-start		: Inicia serviços de dependência de testes containerizados.
infra-test-start:
	@docker-compose up -d --build mongo-test redis-test

## infra-test-stop		: Interrompe containers de testes.
infra-test-stop:
	@docker-compose rm -fsv test-unit test-integration test-all redis-test mongo-test 
