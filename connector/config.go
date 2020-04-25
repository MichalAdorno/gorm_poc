package connector

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type DbConfig struct {
	DbName   string `envconfig:"POSTGRES_DB"`
	Host     string `envconfig:"DOCKER_POSTGRES_IP"`
	Port     string `envconfig:"POSTGRES_PORT"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
}

func ReadInDbConfig() *DbConfig {
	var dbConfig DbConfig
	err := envconfig.Process("", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &dbConfig
}
