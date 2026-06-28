.PHONY: build build-server build-client run-server run-client proto clean

build: build-server build-client

build-server:
	go build -o bin/server ./server

build-client:
	go build -o bin/client ./client

run-server:
	go run ./server

run-client:
	go run ./client

proto:
	protoc --proto_path=protobuf --go_out=. --go-grpc_out=. protobuf/climber.proto

clean:
	rm -rf bin/
