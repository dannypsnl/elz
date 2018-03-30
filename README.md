# elz

[![Build Status](https://travis-ci.org/elz-lang/elz.svg?branch=master)](https://travis-ci.org/elz-lang/elz)

Elz is a modern programming language focus on production. I hope it can change daily work of us.

## Install

```bash
$ go get -d github.com/elz-lang/elz # -d help you don't check compile
$ go get -d github.com/antlr/antlr4/runtime/Go/antlr
$ go get -d llvm.org/llvm/bindings/go/llvm # Goto see Dependencies
$ cd $GOPATH/src/github.com/elz-lang/elz/ && go install
```

Make sure your `$GOPATH/bin` is one of `$PATH`

#### cgo flags problem could happen at Go 1.9.4

If you see a compilation error while compiling your code with Go 1.9.4 or later as follows,

`go build llvm.org/llvm/bindings/go/llvm: invalid flag in #cgo LDFLAGS: -Wl,-headerpad_max_install_names`
you need to setup $CGO_LDFLAGS_ALLOW to allow a compiler to specify some linker options:

`$ export CGO_LDFLAGS_ALLOW='-Wl,(-search_paths_first|-headerpad_max_install_names)'`

### Dependencies

- antlr-runtime for go<br>
`go get github.com/antlr/antlr4/runtime/Go/antlr`
- go-llvm<br>
`go get -d llvm.org/llvm/bindings/go/llvm`<br>
Then compile it. You can follow [Go bindings](http://llvm.org/svn/llvm-project/llvm/trunk/bindings/go/README.txt)<br>
And see my [suggest](http://routedan.blogspot.com/2017/12/go-binding-llvm.html)
- llvm tools(clang, lli, lld ...)

## Usage

Only for develop just now.

```bash
$ elz source.elz  # elz will compile it to llvm ir then print it.
```

## Contributing

Watch [contributing.md](https://github.com/elz-lang/elz/blob/master/CONTRIBUTING.md)
