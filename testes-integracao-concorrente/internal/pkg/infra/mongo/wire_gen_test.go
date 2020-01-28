// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package mongo

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
)

// Injectors from wire.go:

func initializeConfigTest() (config.Configuration, error) {
	string2 := newIntegrationConfigFile()
	configuration, err := config.NewConfig(string2)
	if err != nil {
		return config.Configuration{}, err
	}
	return configuration, nil
}

func initializeDBTest(config2 config.Configuration) (*DB, error) {
	loggingLevels := config.NewLoggingLevels(config2)
	loggerMongo := logging.NewMongo(loggingLevels)
	db, err := NewDB(config2, loggerMongo)
	if err != nil {
		return nil, err
	}
	return db, nil
}