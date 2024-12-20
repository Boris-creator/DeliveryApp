BINARY = ../../tmp/geo-suggest
PROTO_PATH = ../../proto

compile:
	cd $(PROTO_PATH) && \
	make compile
build: compile
	go build -o $(BINARY) .
run: build 
	$(BINARY)