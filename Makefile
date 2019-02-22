all: test build

GO_MOD=GO111MODULE=on

.PHONY: deps test build generate coverage
deps:
	@$(GO_MOD) go get ./...
test: deps generate
	@$(GO_MOD) go test -v ./... -count 1 -cover
build: deps generate
	@$(GO_MOD) go build
generate:
	@$(GO_MOD) go generate ./...
coverage: deps
	@$(GO_MOD) go test -coverprofile=coverage.txt ./...
