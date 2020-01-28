// +build test integration

package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var now = time.Now()

const configTemplate = `
{
	"environment": {
		"name": "test-%s"
	},
    "server": {
        "port": "3080"
    },
    "redisDB": {
        "host": "host-IT-test",
        "port": 6379
	},
    "logging": {
        "level": {
            "ROOT": "INFO"
        }
    }
}
`

// func TestCreateNewInvalidConfig(t *testing.T) {
// 	c, err := NewConfig("./configs/config-dev-inexistente.json")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, c)
// }

func TestLoadConfig(t *testing.T) {
	f, d, err := createTmpFile()

	if err != nil {
		t.Fatalf("Falha ao criar arquivo temporário para teste %v\n", err)
	}
	defer cleanResources(f, d)

	dateTime := now.Format("2006-01-02 15:04:05")
	writeFile(f, fmt.Sprintf(configTemplate, dateTime))

	expect := struct {
		environment string
		serverPort  string
		redisDbHost string
		redisDbPort int
		logging     map[string]string
	}{
		environment: "test-" + dateTime,
		serverPort:  "3080",
		redisDbHost: "host-IT-test",
		redisDbPort: 6379,
		logging: map[string]string{
			"ROOT": "INFO",
		},
	}

	var c Configuration

	c, err = loadConfig(f.Name())

	if err != nil {
		t.Errorf("Erro ao carregar configurações %v", err)
		return
	}

	assert.Equal(t, expect.environment, c.Environment.Name)
	assert.Equal(t, expect.serverPort, c.Server.Port)
	assert.Equal(t, expect.redisDbHost, c.RedisDB.Host)
	assert.Equal(t, expect.redisDbPort, c.RedisDB.Port)

	for k, v := range expect.logging {
		z, ok := c.Logging.Level[k]
		if !ok {
			t.Errorf("Log de nível %q não encontrado na lista\n", k)
		}
		if ok && v != z {
			t.Errorf("Log esperado %q[%s] é diferente de %q[%s]\n", k, v, k, z)
		}
	}

	if err := f.Close(); err != nil {
		t.Fatalf("Falha ao fechar arquivo temporário para teste %v\n", err)
	}
}

func cleanResources(tmpFile *os.File, tempDir string) {
	os.Remove(tmpFile.Name())
	os.Remove(tempDir)
}

func createTmpFile() (tmpFile *os.File, tempDir string, err error) {

	tempDir, err = ioutil.TempDir("", "example-equipment-rental-api")
	if err != nil {
		return
	}
	tmpFile, err = ioutil.TempFile(tempDir, "config-testing-*.json")
	return

}

func writeFile(tmpFile *os.File, content string) (err error) {

	text := []byte(content)
	if _, err = tmpFile.Write(text); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}
	return
}
