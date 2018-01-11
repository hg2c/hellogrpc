.PHONY: build
build:
	go build -o build/greeter_server greeter_server/main.go
	go build -o build/greeter_client greeter_client/main.go
