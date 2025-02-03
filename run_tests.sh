goose --dir=./sql/schema postgres postgres://postgres:postgres@postgres:5432/app_test?sslmode=disable up

go test -v --failfast -coverprofile=coverage.out

goose --dir=./sql/schema postgres postgres://postgres:postgres@postgres:5432/app_test?sslmode=disable reset

go tool cover -html=coverage.out