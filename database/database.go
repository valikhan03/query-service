package database

import(
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"

	"github.com/valikhan03/search-service/models"
)

func InitDatabase() *sqlx.DB {
	configs := models.ConfigsGlobal.DB
	connStr := fmt.Sprintf("host=%s, port=%d, user=%s, name=%s, password=%s, sslmode=%s", 
					configs.Host, configs.Port, configs.User, configs.DBName, configs.Password, configs.SSLMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}