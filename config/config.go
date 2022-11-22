package config

import "log"

var SvcConf *ServiceConfig

type PostgresDB struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DBType string `yaml:"dbtype"`
	DBName string `yaml:"dbname"`
}
type ServiceConfig struct {
	Postgres *PostgresDB `yaml:"postgers"`
	SvcPort  int         `yaml:"svcport"`
}

func NewServiceConfig() {
	SvcConf = &ServiceConfig{
		Postgres: &PostgresDB{
			Host:   "localhost",
			Port:   5432,
			DBType: "postgres",
			DBName: "postgres",
		},
		SvcPort: 8080,
	}

	log.Println("Created new config", SvcConf)
}
