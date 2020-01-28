package main

import (
	"fmt"
	"log"
	"os"

	app "github.com/example/equipment-rental/internal/app"
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
)

func main() {
	fmt.Println("iniciando serviço")
	defer fmt.Println("serviço encerrado")

	var err error
	var config config.Configuration
	var appender logging.FileAppender
	var app *app.EquipmentRentalApp

	config, err = initializeConfig("")
	if err != nil {
		fmt.Printf("Erro ao carregar configurações %v\n", err)
		return
	}

	log.Println("Configurações carregadas!")

	if config.Logging.File != "" {
		log.Printf("Arquivo de logging %q\n", config.Logging.File)
		appender, _ = initializeAppender(config)
		appender.Configure()
	} else {
		log.SetOutput(os.Stdout)
	}

	app, err = initializeApp(config)
	if err != nil {
		fmt.Printf("Erro ao iniciar aplicação %v\n", err)
		return
	}

	app.Bootstrap()
}
