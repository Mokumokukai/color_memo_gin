FROM golang:1.17.1-alpine3.13 as builder

WORKDIR /go/src
COPY . .
RUN apk update && apk add git
RUN go mod download
RUN go get github.com/cosmtrek/air@v1.29.0

