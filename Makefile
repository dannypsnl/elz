all: test

.PHONY: test build coverage
test:
	@go test -v ./... -count 1 -cover
build:
	@go build
coverage:
	@go test -coverprofile=coverage.txt ./...
