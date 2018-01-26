## Commit Rules
[add]: create a new file or add a new function.
[feat]: create a new feature. !!! A known feature, if not, create an issue then implement it
[fix]: Fixed known bug, !!! If this bug is an unknown bug, please create an issue then fixed it.
[clean]: comment, reformat(gofmt only), refactor. Include improving tests
[test]: new test

## translating rule
I will write a gitbook to describe more infomation.
[elz-contributing](https://www.gitbook.com/book/dannypsnl/elz-contributing/welcome)

## Develop Dependencies
- antlr-runtime for go
`go get github.com/antlr/antlr4/runtime/Go/antlr`
- go-llvm
`go get -d llvm.org/llvm/bindings/go/llvm`<br>
Then compile it. You can follow [Go bindings](http://llvm.org/svn/llvm-project/llvm/trunk/bindings/go/README.txt)<br>
And see my [suggest](http://routedan.blogspot.com/2017/12/go-binding-llvm.html)
- llvm tools(clang, lli, lld ...)
- antlr4
```bash
$ cd /usr/local/lib
$ sudo curl -O http://www.antlr.org/download/antlr-4.7-complete.jar
$ export CLASSPATH=".:/usr/local/lib/antlr-4.7-complete.jar:$CLASSPATH"
$ alias antlr4='java -jar /usr/local/lib/antlr-4.7-complete.jar'
$ alias grun='java org.antlr.v4.gui.TestRig'
```
