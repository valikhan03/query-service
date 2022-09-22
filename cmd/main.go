package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"search-service/elastic"
	"search-service/models"
	"search-service/pb"
	"search-service/services"
)

func main() {
	models.InitConfigs()
	service := services.NewService(elastic.NewElasticClient())
	server := grpc.NewServer(grpc.WithInsecure())

	pb.RegisterSearchServiceServer(server, service)

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%s", models.ConfigsGlobal.Server.Host, models.ConfigsGlobal.Server.Port))
	if err != nil {
		panic(err)
	}

	server.Serve(conn)
}
