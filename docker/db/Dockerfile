#compile goose
FROM golang:1.23.2-bullseye AS builder
WORKDIR /app
RUN GOBIN=/app go install github.com/pressly/goose/v3/cmd/goose@latest

#
FROM postgres:17.2

WORKDIR /app

# Install h3 extension
# TODO: debug errors
RUN apt-get update && apt-get install -y postgresql-server-dev-14 pgxnclient make cmake
RUN pgxn install h3 \
    && psql -U ${DB_USER} -c "CREATE EXTENSION IF NOT EXISTS h3;"


COPY --from=builder /app/goose .
COPY ./db/migrations ./migrations
COPY ./db/init.sh /docker-entrypoint-initdb.d/init.sh
EXPOSE 5432

WORKDIR /app/migrations