// +build wireinject

package mongo

import (
	"github.com/google/wire"
	"github.com/example/equipment-rental/internal/pkg/commom/config"
)

func initializeConfigTest() (config config.Configuration, err error) {
	panic(wire.Build(pkgSetConfigTest))
}

func initializeDBTest(config config.Configuration) (m *DB, err error) {
	panic(wire.Build(pkgSetTest))
}
