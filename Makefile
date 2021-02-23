.PHONY: test compile

build: test
	go build -o build/provisioner

test:
	go test ./...
