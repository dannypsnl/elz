all: test build

GO_MOD=GO111MODULE=on

.PHONY: test build coverage
test:
	@$(GO_MOD) go test -v ./... -count 1 -cover
build:
	@$(GO_MOD) go build
coverage:
	@$(GO_MOD) go test -coverprofile=coverage.txt ./...
