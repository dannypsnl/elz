all: generate deps

GO_MOD=GO111MODULE=on

.PHONY: install deps test generate coverage
install: generate
	@$(GO_MOD) go install ./src/cmd/elz
	@$(MAKE) -C core
deps:
	@$(GO_MOD) go get ./...
test: deps generate
	@$(GO_MOD) go test ./... -count 1 -cover -failfast
generate:
	@$(GO_MOD) go generate ./...
coverage: deps
	@$(GO_MOD) go test -coverprofile=coverage.txt ./...
