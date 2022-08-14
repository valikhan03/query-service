package elastic

import(
	elastic "github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() *elastic.Client {
	config := elastic.Config{}
	client, err := elastic.NewClient(config)
	if err != nil{

	}

	return client
}