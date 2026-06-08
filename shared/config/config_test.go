package config_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dadencukillia/stakeholders/shared/config"
)

func TestSimpleConfigParsing(t *testing.T) {
	toml := `
	[database]
	user = "root"
	password = "mysupersecretpass"
	host = "postgres"
	port = 5173
	name = "stakeholders"
	`

	params, err := config.ParseConfig([]byte(toml))
	if err != nil {
		t.Fatal(err)
	}

	expected := config.ServiceConfig {
		Database: config.ServiceDatabaseConfig {
			User: "root",
			Password: "mysupersecretpass",
			Host: "postgres",
			Port: 5173,
			Name: "stakeholders",
		},
	}

	if !reflect.DeepEqual(params, expected) {
		t.Fatal(fmt.Errorf("Incorrect result"))
	}
}
