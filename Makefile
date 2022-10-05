# Here you can reformat, check or build the binary.
GIT_TAG=$(shell git describe --tags --abbrev=0)

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test ./...

lint:
	@golangci-lint run ./...

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

version:
	@echo "version: ${GIT_TAG}"