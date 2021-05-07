APP?=alro

export GO111MODULE=on

.PHONY: download
download:
# download named module into module cache(local storing)
	go mod download

.PHONY: build
build:
	$(info build alro)
	go build -o build/${APP} cmd/main.go

.PHONY: proto
proto:
ifeq ($(shell which protoc),)
	$(info no protoc complier. install protoc)
else
	$(info gen file from .proto)
	$Q cd idl && \
	protoc -I. \
	--go_out=../model \
	--go_opt=paths=source_relative \
	--go-grpc_out=../model \
	--go-grpc_opt=require_unimplemented_servers=false,paths=source_relative *.proto
endif

