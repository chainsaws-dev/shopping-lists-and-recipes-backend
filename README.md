# Web server for recipes and shopping list
Web server made for storing data on shopping lists and recipes based on GO.
Made as easy to handle as possible.  
(in development)

Frontend source for this backend is available [here](https://github.com/chainsaws-dev/shopping-lists-and-recipes)

## Prerequisites
Several ffmpeg libraries reqired:
`sudo apt-get install libavcodec-dev`
`sudo apt-get install libavformat-dev`
`sudo apt-get install libavutil-dev`
`sudo apt-get install libswscale-dev`

PostgreSQL server installed and configured reqired:

1. Add repository
`sudo apt-get install wget ca-certificates`
`wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -`
`echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" |sudo tee  /etc/apt/sources.list.d/pgdg.list`

2. Install PostgreSQL
`sudo apt-get update`
`sudo apt -y install postgresql-13 postgresql-client-13`
`sudo service postgresql start` or `pg_ctlcluster 13 main start`

3. Change default password for postgres 
`sudo su - postgres`
`psql -c "alter user postgres with password 'new password'"`

## Getting source
Run this command to get source code:

`go get github.com/chainsaws-dev/shopping-lists-and-recipes-backend`

## Docs
Run this command to start documentation server on localhost:6060:

`godoc -http=:6060` 

Documentation on this server will be available [here](http://localhost:6060/pkg/shopping-lists-and-recipes/)

## Run
To run project without compilation run the following:

`go run *.go`

## Build
To install in GO bin directory run the following:

`go install *.go`

To build in project root directory run the following:

`go build *.go`

## Run unit tests (TODO)
`go test ./...`

## Run benchmarks (TODO)
`go test -bench .`

## Check test coverage 
`go test -cover`

## Export coverage stats
`go test -coverprofile c.out`
`go tool cover -html=c.out`


## Dockerfile

    FROM golang:1.16-alpine
    WORKDIR /go/src/shopping-lists-and-recipes
    ENV DATABASE_HOST db
    COPY . . 
    RUN go get -d -v ./...
    RUN apk update
    RUN apk add ffmpeg-dev build-base
    RUN go install 
    RUN rm -rf /go/src/*
    WORKDIR /go/bin
    COPY ./public ./public
    COPY ./logs ./logs
    COPY ./settings.json ./settings.json
    EXPOSE 10443
    EXPOSE 8080
    CMD ["./shopping-lists-and-recipes", "-clean", "-makedb", "-noresetroles", "-admincred:example@example.com@@password", "-url:http://localhost:8080"]

## Docker-compose.yml

version: "3.7"
services:
    db:
        image: postgres:13
        environment: 
            POSTGRES_USER: postgres
            POSTGRES_PASSWORD: password
            LC_COLLATE: 'C.UTF-8'
        restart: unless-stopped
        volumes:
        - ./postgres-data:/var/lib/postgresql/data
        ports:
        - '5432:5432'
    web:
        image: shopping-lists-and-recipes_web:latest
        build:
            context: .
            dockerfile: Dockerfile      
        depends_on:
        - db        
        restart: unless-stopped         
        volumes:
        - ./logs:/go/bin/logs
        - ./public:/go/bin/public
        - ./settings.json:/go/bin/settings.json
        ports:
        - 8080:8080
        - 10443:10443
