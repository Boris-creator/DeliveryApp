GOPATH:=$(shell $(shell which go) env GOPATH)

build:
	go build -o ./tmp/server cmd/main.go
swagger:
	$(GOPATH)/bin/swag init -generalInfo ./cmd/main.go -o ./api/docs --ot go,json