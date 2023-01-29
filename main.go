package main

import (
	"fmt"

	"search-service/elastic"
	"search-service/models"
	"search-service/services"
	"search-service/api/v1"
	"search-service/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitConfigs()
	service := services.NewService(elastic.NewElasticClient())
	handler := handlers.NewHandler(service)

	router := gin.Default()

	api.Routes(router, handler)
	

	fmt.Printf("Starting server %s:%s...", models.ConfigsGlobal.Server.Host, models.ConfigsGlobal.Server.Port)
	router.Run(fmt.Sprintf("%s:%s", models.ConfigsGlobal.Server.Host, models.ConfigsGlobal.Server.Port))	
}


