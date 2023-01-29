package models

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigsGlobal Configs

type Configs struct {
	Elastic ElasticConfigs
	Server  ServerConfigs
}

type ElasticConfigs struct {
	Addrs    []string `yaml:"addrs"`
	Username string   `yaml:"username"`
	Password string   `yaml:"-"`
}

type ServerConfigs struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfigs() {
	var elasticConfigs ElasticConfigs
	var serverConfigs ServerConfigs

	elasticConfFile, err := ioutil.ReadFile("configs/elastic.yaml")
	if err != nil{
		log.Fatal(err)
	}

	err = yaml.Unmarshal(elasticConfFile, &elasticConfigs)
	if err != nil{
		log.Fatal(err)
	}

	elasticConfigs.Password = os.Getenv("ELASTIC_PASSWORD")

	serverConfFile, err := ioutil.ReadFile("configs/server.yaml")
	if err != nil{
		log.Fatal(err)
	}

	err = yaml.Unmarshal(serverConfFile, &serverConfigs)
	if err != nil{
		log.Fatal(err)
	}

	ConfigsGlobal.Elastic = elasticConfigs
	ConfigsGlobal.Server = serverConfigs
}