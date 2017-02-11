default: run

run: 
	GOPATH=${PWD} go run ./src/cmd/grpc-server/main.go
.PHONY: run