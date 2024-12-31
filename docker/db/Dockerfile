FROM golang:1.23.2-bullseye AS builder
WORKDIR /app
RUN GOBIN=/app go install github.com/pressly/goose/v3/cmd/goose@latest

#
FROM postgres:17.2

WORKDIR /app

COPY --from=builder /app/goose .
COPY ./db/migrations ./migrations
COPY ./db/init.sh /docker-entrypoint-initdb.d/init.sh
EXPOSE 5432

WORKDIR /app/migrations