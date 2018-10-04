CGO_CXXFLAGS=-std=c++11
CGO_LDFLAGS_ALLOW='Wl,(search_paths_first|headerpad_max_install_names)'
case "${LLVM_VER}" in
    60)
        export CC=clang-6.0
        export CGO_CPPFLAGS="`llvm-config-6.0 --cppflags`"
        export CGO_LDFLAGS="`llvm-config-6.0 --ldflags --libs --system-libs all`"
        mkdir -p $GOPATH/src/llvm.org
        cd $GOPATH/src/llvm.org
        curl -XGET http://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz >> llvm.tar.xz
        tar -xf llvm.tar.xz
        mv llvm-6.0.0.src llvm
        go build $GOPATH/src/llvm.org/llvm/tools/llvm-go/llvm-go.go
        ./llvm-go print-config > $GOPATH/src/llvm.org/llvm/bindings/go/llvm/llvm_config.go
        go get github.com/dannypsnl/assert
        go get github.com/antlr/antlr4/runtime/Go/antlr
        ;;
    *)
        export CC=clang-8
        export CGO_CPPFLAGS="`llvm-config-8 --cppflags`"
        export CGO_LDFLAGS="`llvm-config-8 --ldflags --libs --system-libs all`"
        go get -d llvm.org/llvm/bindings/go/llvm
        go build $GOPATH/src/llvm.org/llvm/tools/llvm-go/llvm-go.go
        ./llvm-go print-config > $GOPATH/src/llvm.org/llvm/bindings/go/llvm/llvm_config.go
        # use `-u` to avoid update llvm
        go get -u -t -v ./...
        ;;
esac

