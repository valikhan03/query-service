package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/valikhan03/tool"
)

func (s *Service) Search(ctx context.Context, req string, page int) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	auctions, err := s.search(ctx, req, page, tool.AuctionsIDX)
	if err != nil {
		return nil, err
	}

	lots, err := s.search(ctx, req, page, tool.LotsIDX)
	if err != nil {
		return nil, err
	}

	res[tool.AuctionsIDX] = auctions
	res[tool.LotsIDX] = lots

	return res, nil
}

func (s *Service) search(ctx context.Context, req string, page int, index string) ([]interface{}, error) {
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"title": map[string]interface{}{
	// 				"query":                req,
	// 				"operator":             "or",
	// 				"minimum_should_match": "50%",
	// 			},
	// 		},
	// 	},
	// }

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query": req,
				"fields": []string{"title", "description"},
				"type":"best_fields",
				"operator":"and",
			},
		},
	}

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(query)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	res, err := s.esconn.Search(
		s.esconn.Search.WithIndex(index),
		s.esconn.Search.WithBody(&buffer),
		//s.esconn.Search.WithSize(15),
		//s.esconn.Search.WithFrom(15*page-1),
		s.esconn.Search.WithTrackTotalHits(true),
		s.esconn.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	resbody := make(map[string]interface{})

	err = json.NewDecoder(res.Body).Decode(&resbody)
	if err != nil {
		log.Fatalf("Service.Search: %s", err.Error())
	}

	if res.IsError() {
		log.Printf("[%s] %s: %s",
			res.Status(),
			resbody["error"].(map[string]interface{})["type"],
			resbody["error"].(map[string]interface{})["reason"],
		)
		return nil, errors.New(fmt.Sprintf("[%s] %s: %s", res.Status(), resbody["error"].(map[string]interface{})["type"], resbody["error"].(map[string]interface{})["reason"]))
	}

	fmt.Println(resbody)
	var response []interface{}

	for _, hit := range resbody["hits"].(map[string]interface{})["hits"].([]interface{}) {
		response = append(response, hit.(map[string]interface{})["_source"])
	}

	return response, nil
}
