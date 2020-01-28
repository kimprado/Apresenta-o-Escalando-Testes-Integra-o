// +build wireinject

package aluguel

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/google/wire"
)

func initializeConfigTest() (config config.Configuration, err error) {
	panic(wire.Build(pkgSetConfigTest))
}

func initializeEquipamentoRepositoryTest(config config.Configuration) (pr *EquipamentoRepository, err error) {
	panic(wire.Build(pkgSetTest))
}
