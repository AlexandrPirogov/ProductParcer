FROM golang:alpine

WORKDIR /ProductParser

COPY . .

RUN go build -o main cmd/main.go && ./main