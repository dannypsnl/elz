all: generate deps

GO_MOD=GO111MODULE=on

.PHONY: install deps test generate coverage
install: generate
	@echo "installing elz compiler..."
	@$(GO_MOD) go install ./src/cmd/elz
	@echo "elz compiler installed"
	@$(MAKE) -C core
deps:
	@$(GO_MOD) go get ./...
test: deps generate
	@$(GO_MOD) go test ./... -count 1 -cover -failfast
	@$(MAKE) -C $(shell pwd)/core/list test
	@$(MAKE) -C $(shell pwd)/core/memory test
	@$(MAKE) -C $(shell pwd)/core/string test
generate:
	@$(GO_MOD) go generate ./...
coverage: deps
	@$(GO_MOD) go test -coverprofile=coverage.txt ./...
