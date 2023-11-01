FROM golang:1.21-alpine
EXPOSE 10443
EXPOSE 8080
WORKDIR /go/src/shopping-lists-and-recipes
ENV DATABASE_HOST db
RUN apk add --no-cache build-base git
COPY . . 
RUN go get -d -v ./...
WORKDIR /go/src/shopping-lists-and-recipes/cmd/app
RUN go build -o $GOPATH/bin/shopping-lists-and-recipes
RUN rm -rf /go/src/*
WORKDIR $GOPATH/bin
COPY ./cmd/app/public ./public
COPY ./cmd/app/logs ./logs
COPY ./cmd/app/settings.json ./settings.json
CMD ["./shopping-lists-and-recipes", "-clean", "-makedb"]
