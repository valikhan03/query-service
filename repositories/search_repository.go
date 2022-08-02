package repositories

import(
	"github.com/elastic/go-elasticsearch/v8"
)

type SearchRepository struct{
	elastic *elasticsearch.Client
}