package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"github.com/valikhan03/search-service/models"
)

type Service struct {
	esconn *elasticsearch.Client
	dbconn *sqlx.DB
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


func (s *Service) GetOwnedAuctions(user_id int) ([]map[string]interface{}, error) {
	query := `select * from vw_owned_auctions where organizer_id=$1`
	var results []map[string]interface{}
	err := s.dbconn.Select(&results, query, user_id)
	if err != nil{
		return nil, err
	}
	if len(results)==0 {
		return nil, models.NO_DATA
	}
	return results, nil
}


func (s *Service) GetParticipatedAuctions(user_id int) ([]map[string]interface{}, error) {
	query := ``
	var results []map[string]interface{}
	err := s.dbconn.Select(&results, query, user_id)
	if err != nil{
		return nil, err
	}
	if len(results)==0 {
		return nil, models.NO_DATA
	}
	return results, nil
}


func (s *Service) GetJoinAuctionRequests(auction_id string) ([]map[string]interface{}, error) {
	query := `select * from admin.vw_attempt_requests where auction_id=$1`
	var results []map[string]interface{}
	err := s.dbconn.Select(&results, query, uuid.MustParse(auction_id))
	if err != nil{
		return nil, err
	}
	if len(results)==0 {
		return nil, models.NO_DATA
	}
	return results, nil
}


func (s *Service) GetAuctionParticipantsList(auction_id string) ([]map[string]interface{}, error) {
	
}