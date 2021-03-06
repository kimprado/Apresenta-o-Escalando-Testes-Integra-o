// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package aluguel

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/example/equipment-rental/internal/pkg/infra/mongo"
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

func initializeEquipamentoRepositoryTest(config2 config.Configuration) (*EquipamentoRepository, error) {
	loggingLevels := config.NewLoggingLevels(config2)
	loggerMongo := logging.NewMongo(loggingLevels)
	db, err := mongo.NewDB(config2, loggerMongo)
	if err != nil {
		return nil, err
	}
	loggerEquipamentoRepository := logging.NewEquipamentoRepository(loggingLevels)
	equipamentoRepository := NewEquipamentoRepository(db, config2, loggerEquipamentoRepository)
	return equipamentoRepository, nil
}
