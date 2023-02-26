package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/valikhan03/search-service/api/v1"
	"github.com/valikhan03/search-service/api/v1/handlers"
	"github.com/valikhan03/search-service/elastic"
	"github.com/valikhan03/search-service/models"
	"github.com/valikhan03/search-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitConfigs()
	service := services.NewService(elastic.NewElasticClient())
	handler := handlers.NewHandler(service)

	router := gin.Default()

	api.Routes(router, handler)
	
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", models.ConfigsGlobal.Server.Host, models.ConfigsGlobal.Server.Port),
		Handler: router,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	fmt.Printf("Starting server %s:%s...", models.ConfigsGlobal.Server.Host, models.ConfigsGlobal.Server.Port)

	server.ListenAndServe()
}


