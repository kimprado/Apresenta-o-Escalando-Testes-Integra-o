package mongo

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/google/wire"
)

var pkgSetConfigTest = wire.NewSet(
	newIntegrationConfigFile,
	config.PkgSet,
)

var pkgSetTest = wire.NewSet(
	PkgSet,
	config.NewLoggingLevels,
	logging.PkgSet,
)

func newIntegrationConfigFile() string {
	return "../../../../configs/config-integration.json"
}
