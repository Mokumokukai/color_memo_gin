FROM golang:1.17.1 as builder

WORKDIR /go/src
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0

ARG HOST_NAME
ARG DB_NAME
ARG DB_PASS
ARG DB_USER
ARG DB_PORT

ENV DB_PATH mysql://${DB_USER}:${DB_PASS}@tcp(${HOST_NAME}:${DB_PORT})/${DB_NAME}
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
