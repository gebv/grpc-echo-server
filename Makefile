default: run

run: 
	GOPATH=${PWD} go run ./src/cmd/grpc-server/main.go --port=48655
.PHONY: run