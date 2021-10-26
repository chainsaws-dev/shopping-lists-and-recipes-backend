FROM golang:1.17-alpine
WORKDIR /go/src/shopping-lists-and-recipes
ENV DATABASE_HOST db
COPY . . 
RUN go get -d -v ./...
RUN apk update
RUN apk add ffmpeg-dev build-base
WORKDIR /go/src/shopping-lists-and-recipes/cmd/app
RUN go build -o $GOPATH/bin/shopping-lists-and-recipes
RUN rm -rf /go/src/*
WORKDIR $GOPATH/bin
COPY ./cmd/app/public ./public
COPY ./cmd/app/logs ./logs
COPY ./cmd/app/settings.json ./settings.json
EXPOSE 10443
EXPOSE 8080
CMD ["./shopping-lists-and-recipes", "-clean", "-makedb", "-noresetroles"]
