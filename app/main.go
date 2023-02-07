package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ygutara/xsis-test/app/config"
	"github.com/ygutara/xsis-test/cinema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("error connect to DB")
	}

	cinema_ := cinema.Cinema{DB: db}
	router := mux.NewRouter()
	cinema_.MovieRoute(router)

	// CORS
	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "POST", "DELETE", "PATCH"},
		AllowCredentials: true,
	})

	// Handler
	fmt.Println("Listening to port 8080...")
	handler := corsWrapper.Handler(router)
	http.ListenAndServe(":8080", handler)
}
