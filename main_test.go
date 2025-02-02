package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RohanIRathi/ReceiptProcessor/database_util"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var apiCfg ApiConfig

func TestMain(m *testing.M) {
	godotenv.Load()
	test_db_url := os.Getenv("TEST_DB_URL")
	if test_db_url == "" {
		test_db_url = "postgres://postgres:postgres@localhost:5432/app_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", test_db_url)
	if err != nil {
		log.Fatalf("Database connection failed! Abort tests!")
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database connection failed! Abort tests!")
		os.Exit(1)
	}

	apiCfg = ApiConfig{
		DB: database_util.New(db),
	}

	code := m.Run()

	os.Exit(code)
}

func TestRespondWithJSON(t *testing.T) {
	w := httptest.NewRecorder()

	respondWithJSON(w, 200, `{"status": "Success"}`)

	var response string
	res := w.Result()
	err := json.NewDecoder(res.Body).Decode(&response)

	if err != nil || res.StatusCode != http.StatusOK || (response) != `{"status": "Success"}` {
		t.Fail()
	}

	type InvalidData struct {
		Name    string
		Channel chan int
	}

	res.Body.Close()

	w = httptest.NewRecorder()
	respondWithJSON(w, 200, InvalidData{Name: "", Channel: make(chan int)})

	res = w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}
