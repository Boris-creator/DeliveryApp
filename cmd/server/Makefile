GOPATH:=$(shell $(shell which go) env GOPATH)

build:
	go build -o ./tmp/server cmd/server/server.go
swagger:
	cd ../../ \
	&& $(GOPATH)/bin/swag init -generalInfo ./cmd/server/server.go -o ./api/docs --ot go,json