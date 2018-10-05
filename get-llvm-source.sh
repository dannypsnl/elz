case "${LLVM_VER}" in
    60)
        mkdir -p $GOPATH/src/llvm.org
        cd $GOPATH/src/llvm.org
        curl -XGET http://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz >> llvm.tar.xz
        tar -xf llvm.tar.xz
        mv llvm-6.0.0.src llvm
        go get github.com/dannypsnl/assert
        go get github.com/antlr/antlr4/runtime/Go/antlr
        ;;
    *)
        go get -d llvm.org/llvm/bindings/go/llvm
        # use `-u` to avoid update llvm
        go get -u -t -v ./...
        ;;
esac

go build $GOPATH/src/llvm.org/llvm/tools/llvm-go/llvm-go.go
./llvm-go print-config > $GOPATH/src/llvm.org/llvm/bindings/go/llvm/llvm_config.go
