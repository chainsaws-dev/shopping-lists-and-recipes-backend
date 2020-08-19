# Web server for recipes and shopping list
Web server made for storing data on shopping lists and recipes based on GO.
Made as easy to handle as possible.  
(in development)

Frontend source for this backend is available [here](https://github.com/chainsaws-dev/shopping-lists-and-recipes)

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