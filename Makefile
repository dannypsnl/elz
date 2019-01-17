all: test build

.PHONY: deps test build coverage
deps:
	@go get -t ./...
test: deps
	@go test -v ./... -count 1 -cover
build: deps
	@go build
coverage: deps
	@go test -coverprofile=coverage.txt ./...
