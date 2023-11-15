FROM golang:alpine

WORKDIR /ProductParser

COPY . .

RUN go build -o /usr/bin/gopars cmd/main.go 