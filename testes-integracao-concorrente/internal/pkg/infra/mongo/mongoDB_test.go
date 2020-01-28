// +build test integration

package mongo

import (
	"testing"
)

func TestConectarDB(t *testing.T) {
	t.Parallel()

	var err error
	c, err := initializeConfigTest()
	if err != nil {
		t.Errorf("Erro ao criar Configuração: %+v\n", err)
		return
	}

	c.MongoDB.Database = c.MongoDB.Database + "-" + t.Name()

	db, err := initializeDBTest(c)
	if err != nil {
		t.Errorf("Conexão banco de dados %v\n", err)
		return
	}

	s := db.GetSession()
	if s == nil {
		t.Errorf("Session não pode ser nula\n")
	}

	err = s.Ping()
	if err != nil {
		t.Errorf("Validação Conexão banco de dados %s\n", err)
		return
	}
}
