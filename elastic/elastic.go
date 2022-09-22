package elastic

import (
	"log"
	"search-service/models"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() *elastic.Client {
	config := elastic.Config{
		Addresses: models.ConfigsGlobal.Elastic.Addrs,
		Username: models.ConfigsGlobal.Elastic.Username,
		Password: models.ConfigsGlobal.Elastic.Password,
	}
	client, err := elastic.NewClient(config)
	if err != nil{
		log.Printf("elasticsearch: %x\n", err)
	}

	return client
}