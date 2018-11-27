case "${LLVM_VER}" in
    60)
        curl -XGET http://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz >> llvm.tar.xz
        tar -xf llvm.tar.xz
        mv $GOPATH/src/llvm.org/llvm/bindings/go/llvm-6.0.0.src ./llvm
        ;;
    70)
        curl -XGET http://releases.llvm.org/7.0.0/llvm-7.0.0.src.tar.xz >> llvm.tar.xz
        tar -xf llvm.tar.xz
        mv $GOPATH/src/llvm.org/llvm/bindings/go/llvm-7.0.0.src ./llvm
        ;;
    *)
        go get -d llvm.org/llvm/bindings/go/llvm
        mv $GOPATH/src/llvm.org/llvm/bindings/go/llvm .
        ;;
esac
