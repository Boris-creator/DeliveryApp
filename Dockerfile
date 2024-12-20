FROM golang:1.23.2-bullseye

WORKDIR /app

COPY . .

RUN go mod download
RUN go install github.com/air-verse/air@latest

EXPOSE 3000

CMD air
RUN cd cmd/geo-suggest && make run