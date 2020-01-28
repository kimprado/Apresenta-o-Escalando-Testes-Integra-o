package app

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/google/wire"
)

// AppSet define providers do pacote
var AppSet = wire.NewSet(
	config.AppSet,
	logging.PkgSet,

	NewEquipmentRentalApp,
)
