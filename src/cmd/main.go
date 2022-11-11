package main

import (
	"agrokan-backend/src/business/domain"
	"agrokan-backend/src/business/usecase"
	"agrokan-backend/src/handler/rest"
	sql "agrokan-backend/src/lib/postgresql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env from local file")
	}

	dbConfig := sql.Config{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
	db := sql.Init(dbConfig)

	d := domain.Init(db)

	uc := usecase.Init(d)

	r := rest.Init(uc)

	r.Run()
}
