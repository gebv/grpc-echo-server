
protobufgen: 
	protoc --go_out=plugins=grpc:. ./src/apps/echo/*.proto
.PHONY: protobufgen


build:
	CGO_ENABLED=0 GOPATH=/go/src/app go build \
	    -o ./bin/grps-server \
		-v \
		-a ./src/cmd/grpc-server/main.go
.PHONY: build