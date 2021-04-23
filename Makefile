APP?=alro

.PHONY: build
build:
	$(info build alro)
	go build -o build/${APP} cmd/main.go
	
