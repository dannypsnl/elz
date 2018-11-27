package llvm

//go:generate sh -c "go build ./tool/llvm-go.go && ./llvm-go print-config > ./llvm_config.go && rm llvm-go"

// #cgo LDFLAGS: -L/usr/local/Cellar/llvm/7.0.0_1/lib -Wl,-search_paths_first -Wl,-headerpad_max_install_names -lffi
import "C"
