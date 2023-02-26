package elastic

import (
	"log"
	"github.com/valikhan03/search-service/models"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() *elastic.Client {
	config := elastic.Config{
		Addresses: []string{models.ConfigsGlobal.Elastic.Addr},
		Username:  models.ConfigsGlobal.Elastic.Username,
		Password:  models.ConfigsGlobal.Elastic.Password,
	}
	client, err := elastic.NewClient(config)
	if err != nil {
		log.Fatalf("elasticsearch: %s\n", err.Error())
	}

	return client
}
