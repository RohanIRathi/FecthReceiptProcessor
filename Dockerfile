FROM golang:1.23.3

WORKDIR /app

# Go installations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY /database_util/. ./database_util/
COPY /sql/. ./sql/
COPY /vendor/. ./vendor/
COPY *.go ./
RUN go env -w CGO_ENABLED=1

# Running test cases
RUN goose --dir sql/schema sqlite3 ./db-test.sqlite3 up
RUN go test
RUN rm -rf ./db-test.sqlite3

# Creating the build
RUN mkdir build
RUN go build -o ./build/receiptprocessor

RUN goose --dir sql/schema sqlite3 ./db.sqlite3 up

CMD ["./build/receiptprocessor"]