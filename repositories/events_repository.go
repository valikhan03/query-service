package repositories

import (
	"bytes"
	"context"
	"encoding/json"

	elastic "github.com/elastic/go-elasticsearch/v8"
	esapi "github.com/elastic/go-elasticsearch/v8/esapi"
)

type EventsRepository struct{
	client *elastic.Client
}

func (er *EventsRepository) CreateAuction(obj map[string]interface{}) {
	data, err := json.Marshal(obj)
	if err != nil{

	}

	req := esapi.CreateRequest{
		Index: "",
		DocumentID: obj["id"].(string),
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError() {

	}

	
}

func (er *EventsRepository) UpdateAuction(obj map[string]interface{}){
	data, err := json.Marshal(obj)
	if err != nil{

	}
	req := esapi.UpdateRequest{
		Index: "",
		DocumentID: obj["id"].(string),
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (er *EventsRepository) DeleteAuction(obj map[string]interface{}) {
	req := esapi.DeleteRequest{
		Index: "",
		DocumentID: obj["id"].(string),
	}

	req.Do(context.Background(), er.client)
}

func (er *EventsRepository) UpdateAuctionStatus(obj map[string]interface{}) {
	req := esapi.UpdateRequest{
		Index: "",
		DocumentID: "",
		//...
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (er *EventsRepository) AddProduct(obj map[string]interface{}) {
	data, err := json.Marshal(obj)
	if err != nil{

	}
	req := esapi.CreateRequest{
		Index: "",
		DocumentID: "",
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (er *EventsRepository) UpdateProduct(obj map[string]interface{}){
	data, err := json.Marshal(obj)
	if err != nil{

	}
	req := esapi.UpdateRequest{
		Index: "",
		DocumentID: "",
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (er *EventsRepository) DeleteProduct(obj map[string]interface{}) {
	req := esapi.DeleteRequest{
		Index: "",
		DocumentID: obj["id"].(string),
	}

	res, err := req.Do(context.Background(), er.client)
	if err != nil{

	}

	if res.IsError(){
		
	}
}
 
func (er *EventsRepository) Addparticipant() {}

func (er *EventsRepository) RemoveParticipant() {}