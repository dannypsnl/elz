all: build

GO_ENV=GO111MODULE=on

.PHONY: check generate test build go-build
check:
	@which protoc || echo "protobuffer compiler `protoc` not found"
	@which pb-rs || echo "use `cargo install pb-rs` to get it"
	@which protoc-gen-go || echo "use `go get -u github.com/golang/protobuf/protoc-gen-go` to get it"
generate:
	@$(GO_ENV) go generate ./...
test: generate go-build
	@cargo test
	@$(GO_ENV) go test ./... -count 1 -cover -failfast
build: generate go-build
	@cargo build
go-build: src/code_generate/code_generate.go
	@go build -o libcode_generate.so -buildmode=c-shared src/code_generate/code_generate.go
