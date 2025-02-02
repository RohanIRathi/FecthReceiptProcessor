ECHO "Creating test db..."

ECHO CREATE DATABASE app_test; | psql -v ON_ERROR_STOP=1 "postgres://postgres:postgres@localhost:5432/receiptprocessor"

goose --dir=./sql/schema postgres postgres://postgres:postgres@localhost:5432/app_test?sslmode=disable up

go test -v --failfast -coverprofile=coverage.out

goose --dir=./sql/schema postgres postgres://postgres:postgres@localhost:5432/app_test?sslmode=disable reset

ECHO DROP DATABASE app_test; | psql -v ON_ERROR_STOP=1 "postgres://postgres:postgres@localhost:5432/receiptprocessor"

go tool cover -html=coverage.out