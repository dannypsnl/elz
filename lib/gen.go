package llvm

//go:generate sh -c "bash ./get-llvm.sh"
//go:generate sh -c "go build ./tool/llvm-go.go && ./llvm-go print-config > ./llvm/llvm_config.go && rm llvm-go"
