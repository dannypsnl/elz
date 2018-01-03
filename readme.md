# elz
elz is a modern programming language.<br>
It trying to solve the most complex issue in production programming.<br>
Why concurrency is so hard? Why design a useful and simple data type is so hard?<br>
Why the access level is so complex?(Who want to know public, external, internal, private, file private?)<br>
Why don't have meta programming? How to testing?<br>
Don't worry anymore. elz is design for this all.<br>
## prototype version
Here is my little try.<br>
[nim-elz](https://github.com/elz-lang/nim-elz)<br>
It based on translate elz to nim, and use it compiler to work.<br>
But that is not enough. nim can not implement some features of elz.<br>
## Install
```bash
$ go get github.com/elz-lang/elz
$ go install
```
Make sure your `$GOPATH/bin` is one of `$PATH`
### Dependencies
- antlr-runtime for go
`go get github.com/antlr/antlr4/runtime/Go/antlr`
- go-llvm
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
Watch [contributing.md](https://github.com/elz-lang/elz/blob/master/contributing.md)<br>
