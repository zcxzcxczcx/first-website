FROM golang:alpine
WORKDIR /first-website 
COPY . /first-website 
entrypoint go run main.go