package aluguel

import (
	"time"

	"github.com/example/equipment-rental/internal/pkg/commom/config"
	"github.com/example/equipment-rental/internal/pkg/commom/logging"
	"github.com/example/equipment-rental/internal/pkg/infra/mongo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Equipamento representa equipamento
type Equipamento struct {
	ID        bson.ObjectId `bson:"_id"`
	Codigo    string        `bson:"codigo"`
	Descricao string        `bson:"descricao"`
	RPM       string        `bson:"descricao"`
	criacao   time.Time     `bson:"criacao"`
	Ativo     bool          `bson:"ativo"`
}

var projecaoEquipamento = bson.M{
	"_id":       1,
	"codigo":    1,
	"descricao": 1,
	"rpm":       1,
	"criacao":   1,
	"ativo":     1,
}

//EquipamentoRepository permite manter dados de paredão
type EquipamentoRepository struct {
	config  config.Configuration
	mongoDB *mongo.DB
	logger  logging.LoggerEquipamentoRepository
}

//NewEquipamentoRepository cria instância de EquipamentoRepository
func NewEquipamentoRepository(mongoDB *mongo.DB, config config.Configuration, l logging.LoggerEquipamentoRepository) *EquipamentoRepository {
	p := new(EquipamentoRepository)
	p.mongoDB = mongoDB
	p.config = config
	p.logger = l
	return p
}

//Ativos retorna equipamentos ativos
func (p *EquipamentoRepository) Ativos() (lista []*Equipamento, err error) {

	session := p.mongoDB.GetSession()
	defer session.Close()

	lista = []*Equipamento{}

	query := bson.M{
		"ativo": true,
	}

	collection := session.DB(p.config.MongoDB.Database).C("Equipamento")
	var q *mgo.Query
	q = collection.Find(query)
	q.Select(projecaoEquipamento)

	err = q.All(&lista)

	return
}

// ConsultarPorID retorna paredão de id informado
func (p *EquipamentoRepository) ConsultarPorID(id bson.ObjectId) (equipamento *Equipamento, err error) {

	session := p.mongoDB.GetSession()
	defer session.Close()

	collection := session.DB(p.config.MongoDB.Database).C("Equipamento")
	var q *mgo.Query
	q = collection.FindId(id)
	q.Select(projecaoEquipamento)

	var equipamentoEncontrado Equipamento
	err = q.One(&equipamentoEncontrado)

	if err == mgo.ErrNotFound {
		err = nil
		return
	}

	if err != nil {
		return
	}

	equipamento = &equipamentoEncontrado

	return
}
