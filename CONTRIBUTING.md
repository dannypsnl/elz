## Commit Rules

Setting git as:

```
$ git config commit.template $GOPATH/src/github.com/elz-lang/elz/.git_message.txt
```

## translating rule

I would write something at [contribute-doc](https://github.com/elz-lang/contribute-doc) as I can,
you can also contribute the document itself too!


## Develop Dependencies

- llvm & its tools(clang, lli, lld ...)

## Test in Development

- test
    ```bash
    $ cargo test
    ```
- compile
    ```bash
    $ cargo build
    ```
