// +build test integration

package aluguel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestConsultarMartelo(t *testing.T) {
	t.Parallel()

	var id bson.ObjectId
	id = bson.ObjectIdHex("555b6e830850536438063761")

	var err error
	c, err := initializeConfigTest()
	if err != nil {
		t.Errorf("Erro ao criar Configuração: %+v\n", err)
		return
	}

	c.MongoDB.Database = c.MongoDB.Database + "-" + t.Name()

	p, err := initializeEquipamentoRepositoryTest(c)
	if err != nil {
		t.Errorf("Erro ao criar Repository: %+v\n", err)
		return
	}

	assert.NotNil(t, p)

	v, err := p.ConsultarPorID(id)
	assert.Nil(t, err)
	assert.NotNil(t, v)

	assert.Equal(t, "Martelo", v.Descricao)
	assert.Equal(t, "e-001", v.Codigo)

}
