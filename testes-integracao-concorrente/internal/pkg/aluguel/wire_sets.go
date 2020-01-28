package aluguel

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/example/equipment-rental/internal/pkg/infra/mongo"
	"github.com/example/equipment-rental/internal/pkg/infra/redis"
	"github.com/google/wire"
)

// PkgSet define providers do pacote
var PkgSet = wire.NewSet(
	NewEquipamentoRepository,
)

var pkgSetConfigTest = wire.NewSet(
	newIntegrationConfigFile,
	config.PkgSet,
)

var pkgSetTest = wire.NewSet(
	PkgSet,
	config.NewLoggingLevels,
	config.NewRedisDB,
	logging.PkgSet,
	mongo.PkgSet,
	redis.PkgSet,
)

func newIntegrationConfigFile() string {
	return "../../../configs/config-integration.json"
}
