package mongo

import (
	"fmt"
	"time"

	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	mgo "gopkg.in/mgo.v2"
)

//DB mantém referência para pool de conexões mongo
type DB struct {
	mongoSession *mgo.Session
	inicializado bool
	config       config.Configuration
	logger       logging.LoggerMongo
}

//NewDB cria instância DB
func NewDB(config config.Configuration, l logging.LoggerMongo) (m *DB, err error) {
	var mongoDB = new(DB)
	mongoDB.config = config
	mongoDB.logger = l
	mongoDB.inicializado = true

	s, err := mongoDB.connect()
	if err != nil {
		return
	}

	mongoDB.mongoSession = s

	mongoDB.logger.Infof("Conectado ao DB\n")
	return mongoDB, err
}

//connect realiza a conexão com o DB
// Promove criação interna do pool do mgo.
// Retorno *mgo.Session funciona como proxy do pool.
// Conexões são obtidas posteriormente com *mgo.Session.Copy()
func (m DB) connect() (s *mgo.Session, err error) {
	var session *mgo.Session

	if m.config.MongoDB.Cluster {
		mongoDBDialInfo := &mgo.DialInfo{
			Direct:   false,
			Addrs:    []string{m.config.MongoDB.Host},
			Timeout:  m.config.MongoDB.Timeout * time.Second,
			Database: m.config.MongoDB.Database,
			Username: m.config.MongoDB.User,
			Password: m.config.MongoDB.Password,
		}
		session, err = mgo.DialWithInfo(mongoDBDialInfo)
	} else if m.config.MongoDB.User != "" {
		connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", m.config.MongoDB.User, m.config.MongoDB.Password, m.config.MongoDB.Host, m.config.MongoDB.Port, m.config.MongoDB.Database)
		session, err = mgo.Dial(connectionString)
	} else if m.config.MongoDB.User == "" {
		connectionString := fmt.Sprintf("mongodb://%s:%d/%s", m.config.MongoDB.Host, m.config.MongoDB.Port, m.config.MongoDB.Database)
		session, err = mgo.Dial(connectionString)
	}

	if err != nil {
		m.logger.Errorf("Erro ao conectar com o mongo %s\n", err)
		return
	}
	s = session
	return
}

//GetSession retorna conexão com o DB
func (m *DB) GetSession() (s *mgo.Session) {
	s = m.mongoSession.Copy()
	return
}
