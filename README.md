# Solution for Fecth's take home exam of a Receipt Processor

This is the my implementation of the take home exam assigned by Fetch to create API endpoint for the Receipt Processor.

## Running application using Docker

I have already dockerized the application with Golang and Postgres services. To run the application, simply run `docker-compose up` and the application will start on port 8000.

## Running the Go app directly

In order to run the application using your device setup, You'll need to configure the Port on which the app will run, as well as the Postgres database settings. You'll need to create a `.env` file and add the following:

1. `PORT=${YourPortNumber}`
1. `DB_URL=${DatabaseConnectionString}`

The defaul for these values are:

1. Port = 8000
1. DB_URL = postgres://postgres:postgres@postgres:5432/receiptprocessor?sslmode=disable

> Note:
>
> 1. The DB_URL variable can be constructed as: `postgres://{Username}:{Password}@{Host}:{Port}/{DatabaseName}?sslmode=disable`
> 1. The Application requires Golang version >=1.23.3 and Postgres version >=15.2

```
Test cases can be run using the run_tests.bat or run_tests.sh scripts outside docker container shells.
```
