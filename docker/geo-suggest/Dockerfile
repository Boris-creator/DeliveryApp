FROM golang:1.23.2-bullseye AS builder

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y unzip
RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" \
    && curl -LO $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip \
    && unzip protoc-25.1-linux-x86_64.zip -d ./downloads \
    && cp ./downloads/bin/protoc /usr/local/bin \
    && cp -R ./downloads/include/* /usr/local/include

RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app/cmd/geo-suggest
RUN make build

FROM ubuntu:22.04

RUN apt-get update && apt-get install make -y
WORKDIR /app
COPY --from=builder /app/tmp/geo-suggest ./tmp/geo-suggest
COPY --from=builder /app/cmd/geo-suggest/Makefile ./cmd/geo-suggest/Makefile
WORKDIR /app/cmd/geo-suggest

CMD make run