# Kentech-films

API  allow to registered users manage films. This API was created as solution for kentech technical assessment for golang backend developer in 2022.

## Architecture ##

## How do I get set up? ##

* Summary of set up

Install go.

Export environment variables.
```
export $(cat .env.example | grep -v ^# | xargs)
```

Go to cmd/kentech-films directory
```
cd cmd/kentech-films
```

Run application
```
go run application.go main.go server.go
```
The API is served in `http://localhost:8080`

* How to run tests

```
gotest='go test ./... -race -p 1'
```

## Docker ##

First build the project

```
export MY_WORKSPACE=$(pwd)
docker-compose build
```
How to run the API
```
docker-compose run --rm -p 8000:8000 api
```
The API is served in `http://localhost:8000`

How to run tests in docker compose

```
docker-compose run --rm test
```

## Swagger ##

Swagger documentation will be served in `http://localhost:8000/swagger/index.html#/` or `http://localhost:8080/swagger/index.html#/` depends in how you ran the API.
