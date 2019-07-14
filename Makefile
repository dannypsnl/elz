all: build

.PHONY: check generate test build
check:
	@which protoc || echo "protobuffer compiler `protoc` not found"
	@which pb-rs || echo "use `cargo install pb-rs` to get it"
	@which protoc-gen-go || echo "use `go get -u github.com/golang/protobuf/protoc-gen-go` to get it"
generate:
	@go generate ./...
test: generate
	@cargo test
build: generate
	@cargo build
