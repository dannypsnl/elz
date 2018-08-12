## Commit Rules

Setting git as:

```
$ git config commit.template $GOPATH/src/github.com/elz-lang/elz/.git_message.txt
```

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
- antlr4 ~4.7.1
    ```bash
	$ git submodule init
	$ git submodule update
	$ go generate ./... # use this command to generate new antlr4 parser!
    ```

## Test in Development

- go test
    ```bash
    # This command have to execute at root of project
    # make sure count=1 to avoid cache
    $ go test ./... -count=1
    ```
- compile example/test.elz
    ```bash
    $ ./elz.sh example/test.elz
    ```
