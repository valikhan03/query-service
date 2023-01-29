package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Service struct {
	esconn *elasticsearch.Client
}

func NewService(conn *elasticsearch.Client) *Service {
	return &Service{
		esconn: conn,
	}
}

func (s *Service) GetAuction(ctx context.Context, id string)  (map[string]interface{}, error) {
	esreq := esapi.GetRequest{
		Index:      "auctions",
		DocumentID: id,
	}

	response, err := esreq.Do(ctx, s.esconn)
	if err != nil {
		log.Printf("Service.GetAuction: %x\n", err)
		return nil, err
	}

	if response.IsError() {
		log.Printf("Service.GetAuction: %s\n", response.Status())
		return nil, err
	}

	var res map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Printf("Service.GetAuction: %x\n", err)
		return nil, err
	}

	return res["_source"].(map[string]interface{}), nil
}


func (s *Service) GetProduct(ctx context.Context, id string) (map[string]interface{}, error) {
	esreq := esapi.GetRequest{
		Index:      "auction-products",
		DocumentID: id,
	}

	response, err := esreq.Do(ctx, s.esconn)
	if err != nil {
		log.Printf("Service.GetProduct: %x", err)
		return nil, err
	}

	if response.IsError() {
		log.Printf("Service.GetProduct: %s", response.Status())
		return nil, err
	}

	var res map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Printf("Service.GetProduct: %x", err)
		return nil, err
	}

	return res["_source"].(map[string]interface{}), nil
}

