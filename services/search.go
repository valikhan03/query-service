package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
)


const(
	auctionsIndex = "auctions"
	productsIndex = "products"
)

func (s *Service) Search(ctx context.Context, req string, page int) ([]interface{}, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": map[string]interface{}{
					"query":                req,
					"operator":             "or",
					"minimum_should_match": "50%",
				},
			},
		},
	}

	var buffer bytes.Buffer
	var elasticIndex string

	switch{
		
	}

	err := json.NewEncoder(&buffer).Encode(&query)
	if err != nil {
		log.Printf("Service.Search: %x", err)
		return nil, err
	}

	res, err := s.esconn.Search(
		s.esconn.Search.WithIndex(elasticIndex),
		s.esconn.Search.WithBody(&buffer),
		s.esconn.Search.WithSize(15),
		s.esconn.Search.WithFrom(15*page-1),
	)
	if err != nil {
		log.Printf("Service.Search: %x", err)
		return nil, err
	}

	var resbody map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&resbody)
	if err != nil {
		log.Printf("Service.Search: %x", err)
		return nil, err
	}

	var response []interface{}

	for _, hit := range resbody["hits"].(map[string]interface{}) {
		response = append(response, hit.(map[string]interface{})["_source"])
	}

	return response, nil
}
