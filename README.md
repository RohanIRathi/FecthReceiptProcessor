# Solution for Fetch's take home exam of a Receipt Processor

This is the my implementation of the take home exam assigned by Fetch to create API endpoint for the Receipt Processor.

## Running application using Docker

I have already dockerized the application with Golang base image. To run the application, simply run `docker build . -t ${IMAGE_NAME}` followed by `docker container run --name ${CONTAINER_NAME} ${IMAGE_NAME}` and the application will start on port 8000.

## Running the Go app directly

In order to run the application using your device setup, You'll need to configure the Port on which the app will run, as well as the sqlite3 database name. You'll need to create a `.env` file and add the following:

1. `PORT=${YourPortNumber}`
1. `DB_URL=${SQLITE DB Name}`

The default for these values are:

1. Port = 8000
1. DB_URL = ./db.sqlite3

> Note: The Application requires Golang version >=1.23.3 and Postgres version >=15.2
