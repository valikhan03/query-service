package models

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigsGlobal Configs

type Configs struct {
	DB      DBConfigs
	Elastic ElasticConfigs
	Server  ServerConfigs
}

type DBConfigs struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"-"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type ElasticConfigs struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"-"`
}

type ServerConfigs struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfigs() {
	var dbConfigs DBConfigs
	var elasticConfigs ElasticConfigs
	var serverConfigs ServerConfigs

	dbConfFile, err := ioutil.ReadFile("configs/db.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(dbConfFile, &dbConfigs)
	if err != nil{
		log.Fatal(err)
	}

	dbConfigs.Password = os.Getenv("DB_PASSWORD")

	elasticConfFile, err := ioutil.ReadFile("configs/elastic.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(elasticConfFile, &elasticConfigs)
	if err != nil {
		log.Fatal(err)
	}

	elasticConfigs.Password = os.Getenv("ELASTIC_PASSWORD")

	serverConfFile, err := ioutil.ReadFile("configs/server.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(serverConfFile, &serverConfigs)
	if err != nil {
		log.Fatal(err)
	}

	ConfigsGlobal.DB = dbConfigs
	ConfigsGlobal.Elastic = elasticConfigs
	ConfigsGlobal.Server = serverConfigs
}
