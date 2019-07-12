all: build

.PHONY: generate test build
generate:
	@go generate ./...
test: generate
	@cargo test
build: generate
	@cargo build