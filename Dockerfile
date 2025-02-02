FROM golang:1.23.3

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY /database_util/. ./database_util/
COPY /sql/. ./sql/
COPY /vendor/. ./vendor/

RUN mkdir build
RUN go build -o ./build/receiptprocessor

RUN echo "#!/bin/bash" > build.sh
RUN echo "goose --dir sql/schema postgres postgres://postgres:postgres@postgres:5432/receiptprocessor up" >> build.sh
RUN echo "./build/receiptprocessor" >> build.sh

RUN chmod 777 ./build.sh

CMD ["./build.sh"]