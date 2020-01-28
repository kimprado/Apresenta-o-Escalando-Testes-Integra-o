// +build wireinject

package main

import (
	app "github.com/example/equipment-rental/internal/app"
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/google/wire"
)

// initializeConfig inicializa Configuration
func initializeConfig(path string) (configuration config.Configuration, err error) {
	panic(wire.Build(config.NewConfig))
}

// initializeAppender inicializa FileAppender
func initializeAppender(configuration config.Configuration) (appender logging.FileAppender, err error) {
	panic(wire.Build(app.AppSet))
}

// initializeApp inicializa EquipmentRentalApp
func initializeApp(configuration config.Configuration) (a *app.EquipmentRentalApp, err error) {
	panic(wire.Build(app.AppSet))
}
