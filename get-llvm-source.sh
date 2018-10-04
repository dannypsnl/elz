case "${LLVM_VER}" in
    60)
        mkdir -p $GOPATH/src/llvm.org
        cd $GOPATH/src/llvm.org
        curl -XGET http://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz >> llvm.tar.xz
        tar -xf llvm.tar.xz
        mv llvm-6.0.0.src llvm
        cd llvm && git init && git remote add origin https://github.com/llvm-mirror/llvm.git
        ;;
    *)
        go get -d llvm.org/llvm/bindings/go/llvm
        ;;
esac
