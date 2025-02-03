package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	database "github.com/RohanIRathi/ReceiptProcessor/database_util"
	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Println("Port not setup in the environment. Trying default PORT 8000")
	}
	portString = "8000"

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Println("DB_URL is not setup in the environment. Trying default connection string")
		dbUrl = "./db.sqlite3"
	}

	conn, err := sql.Open("sqlite3", dbUrl)

	if err != nil {
		log.Fatalf("Cannot connect to the database! Error: %v", err)
	}

	queries := database.New(conn)

	apiCfg := ApiConfig{
		DB: queries,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("POST /receipts/process", apiCfg.handleCreateReceipt)
	handler.HandleFunc("GET /receipts/{id}/points", apiCfg.handleGetPoints)

	router := &http.Server{
		Addr:    ":" + portString,
		Handler: handler,
	}
	log.Printf("Server running at %v", router.Addr)

	router.ListenAndServe()
}
