# Sensor API

This repository contains a Sensor API that collects data from sensors and performs various aggregations. The API is built using the Go language and the Gin framework.

## Features

- Generate the meta data for both sensors & sensor groups
- Collects data from sensors
- Performs aggregations on sensor data
- Built with Go and Gin framework
- PostgreSQL for data storage
- Docker Compose for local development and testing
- Includes Swagger documentation for API endpoints
- Redis for few aggregation endpoints

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Local Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/srik007/sensor-api.git
    cd sensor-api
    ```

2. **Rename `env.example` to  `.env` file in the project root and configure the data:**

    ```dotenv
    PP_ENV=local
    DB_HOST=postgres
    DB_DRIVER=postgres
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=postgres
    DB_PORT=5432
    PORT=8080
    REDIS_URL=redis:6379
    NUMBER_OF_SENSORS=3
    ```

## Running the API

1. Build docker image for current go application using

    ```bash
        docker build -t app:latest .
    ```
2. Use Docker Compose to set up the local environment:

    ```bash
    docker-compose up
    ```
The API will be available at http://localhost:8080

## Swagger Documentation

Swagger documentation is available at http://localhost:8080/swagger/index.html

## Testing

1. Use docker-compose.test.yml to spin up the containers

    ```bash
    docker-compose -f docker-compose.test.yml up
    ```
2. Run the tests

    ```bash
    go test
    ```
