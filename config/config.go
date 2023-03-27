package config

import (
	"context"
	"log"

	"github.com/psbernardo/dockertest/config/database_maria"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	MariaDB *database_maria.Config `env:",prefix=MARIA_DB_"`
}

func Read() *Config {
	conf := &Config{}
	if err := envconfig.Process(context.Background(), conf); err != nil {

		log.Fatal(err)
	}
	return conf
}
