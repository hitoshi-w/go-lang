# Build Stage
FROM golang:1.21.1-alpine3.16 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -buildvcs=false -o /app/build/main cmd/main.go

# Local
FROM golang:1.21.1 AS development
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest && go install github.com/rubenv/sql-migrate/sql-migrate@latest

EXPOSE 3000
