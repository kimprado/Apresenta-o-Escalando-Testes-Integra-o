package app

import (
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
)

// EquipmentRentalApp representa instância da aplicação
type EquipmentRentalApp struct {
	logger logging.Logger
}

// NewEquipmentRentalApp cria app
func NewEquipmentRentalApp(l logging.Logger) (a *EquipmentRentalApp) {
	a = new(EquipmentRentalApp)
	a.logger = l
	return
}

// Bootstrap é responsável por iniciar a aplicação
func (a *EquipmentRentalApp) Bootstrap() {
	a.logger.Infof("Iniciando serviços da aplicação...\n")

}
