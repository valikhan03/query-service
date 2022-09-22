package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"search-service/pb"

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

func (s *Service) GetAuction(ctx context.Context, req *pb.GetAuctionInfoRequest) (*pb.GetAuctionInfoResponse, error) {
	esreq := esapi.GetRequest{
		Index:      "auctions",
		DocumentID: req.AuctionId,
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

	data, err := json.Marshal(res["_source"].(map[string]interface{}))
	if err != nil {
		log.Printf("Service.GetAuction: %x\n", err)
		return nil, err
	}

	resp := pb.GetAuctionInfoResponse{
		Auction: data,
		Status:  200,
	}

	return &resp, nil
}

func (s *Service) SearchAuctions(ctx context.Context, req *pb.SearchAuctionsRequest) (*pb.SearchAuctionsResponse, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": map[string]interface{}{
					"query":                req.SearchRequest,
					"operator":             "or",
					"minimum_should_match": "50%",
				},
			},
		},
	}

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(&query)
	if err != nil {
		log.Printf("Service.SearchAuctions: %x", err)
		return nil, err
	}

	res, err := s.esconn.Search(
		s.esconn.Search.WithIndex("auctions"),
		s.esconn.Search.WithBody(&buffer),
		s.esconn.Search.WithSize(15),
		s.esconn.Search.WithFrom(15*(int(req.Page)-1)),
	)
	if err != nil {
		log.Printf("Service.SearchAuctions: %x", err)
		return nil, err
	}

	var resbody map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&resbody)
	if err != nil {
		log.Printf("Service.SearchAuctions: %x", err)
		return nil, err
	}

	response := pb.SearchAuctionsResponse{}

	for _, hit := range resbody["hits"].(map[string]interface{}) {
		data, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			log.Printf("Service.SearchAuctions: %x", err)
			return nil, err
		}
		response.Auctions = append(response.Auctions, data)
	}

	return &response, nil
}

func (s *Service) GetProduct(ctx context.Context, req *pb.GetProductInfoRequest) (*pb.GetProductInfoResponse, error) {
	esreq := esapi.GetRequest{
		Index:      "auction-products",
		DocumentID: req.ProductId,
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

	data, err := json.Marshal(res["_source"].(map[string]interface{}))
	if err != nil {
		log.Printf("Service.GetProduct: %x", err)
		return nil, err
	}

	resp := pb.GetProductInfoResponse{
		Product: data,
		Status:  200,
	}

	return &resp, nil
}

func (s *Service) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": map[string]interface{}{
					"query":                req.SearchRequest,
					"operator":             "or",
					"minimum_should_match": "50%",
				},
			},
		},
	}

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(&query)
	if err != nil {
		log.Printf("Service.SearchProduct: %x", err)
		return nil, err
	}

	res, err := s.esconn.Search(
		s.esconn.Search.WithIndex("auction-products"),
		s.esconn.Search.WithBody(&buffer),
		s.esconn.Search.WithSize(15),
		s.esconn.Search.WithFrom(15*(int(req.Page)-1)),
	)
	if err != nil {
		log.Printf("Service.SearchProduct: %x", err)
		return nil, err
	}

	var resbody map[string]interface{}

	err = json.NewDecoder(res.Body).Decode(&resbody)
	if err != nil {
		log.Printf("Service.SearchProduct: %x", err)
		return nil, err
	}

	response := pb.SearchProductsResponse{}

	for _, hit := range resbody["hits"].(map[string]interface{}) {
		data, err := json.Marshal(hit.(map[string]interface{})["_source"])
		if err != nil {
			log.Printf("Service.SearchProduct: %x", err)
			return nil, err
		}
		response.Products = append(response.Products, data)
	}

	return &response, nil
}