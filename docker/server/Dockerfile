FROM golang:1.23.2-bullseye

WORKDIR /app

COPY . /app

WORKDIR /app
RUN go mod download
RUN go install github.com/air-verse/air@latest

EXPOSE 3000

CMD air