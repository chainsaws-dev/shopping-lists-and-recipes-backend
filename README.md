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
`sudo service postgresql start` or `pg_ctlcluster 12 main start`

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