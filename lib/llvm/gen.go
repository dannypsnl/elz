package llvm

//go:generate sh -c "go build ./tool/llvm-go.go && ./llvm-go print-config > ./llvm_config.go && rm llvm-go"
