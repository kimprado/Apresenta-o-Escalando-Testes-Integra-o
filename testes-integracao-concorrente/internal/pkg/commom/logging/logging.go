package logging

import (
	"github.com/example/equipment-rental/internal/pkg/commom/config"
	l "github.com/kimprado/sllog/pkg/logging"
)

// Logger para logar ROOT
type Logger struct {
	l.Logger
}

//LoggerMongo para logar infra.database.mongo
type LoggerMongo struct {
	l.Logger
}

// LoggerRedisDB para logar infra.redis.db
type LoggerRedisDB struct {
	l.Logger
}

//LoggerEquipamentoRepository para logar equipamento.repository
type LoggerEquipamentoRepository struct {
	l.Logger
}

// LoggerWebServer para logar webserver
type LoggerWebServer struct {
	l.Logger
}

// NewLogger cria Logger ""(ROOT)
func NewLogger(configLevels config.LoggingLevels) (log Logger) {
	log = Logger{l.NewLogger("", configLevels)}
	return
}

//NewMongo cria Logger "infra.database.mongo"
func NewMongo(configLevels config.LoggingLevels) (log LoggerMongo) {
	log = LoggerMongo{l.NewLogger("infra.database.mongo", configLevels)}
	return
}

// NewRedisDB cria Logger "infra.redis.db"
func NewRedisDB(configLevels config.LoggingLevels) (log LoggerRedisDB) {
	log = LoggerRedisDB{l.NewLogger("infra.redis.db", configLevels)}
	return
}

//NewEquipamentoRepository cria Logger "equipamento.repository"
func NewEquipamentoRepository(configLevels config.LoggingLevels) (log LoggerEquipamentoRepository) {
	log = LoggerEquipamentoRepository{l.NewLogger("equipamento.repository", configLevels)}
	return
}

// NewWebServer cria Logger "webserver"
func NewWebServer(configLevels config.LoggingLevels) (log LoggerWebServer) {
	log = LoggerWebServer{l.NewLogger("webserver", configLevels)}
	return
}
