# elz

[![version badges](https://img.shields.io/badge/version-0.1.0-blue.svg)](https://github.com/elz-lang/elz/releases)
[![Build Status](https://travis-ci.org/elz-lang/elz.svg?branch=master)](https://travis-ci.org/elz-lang/elz)
[![codecov](https://codecov.io/gh/elz-lang/elz/branch/master/graph/badge.svg)](https://codecov.io/gh/elz-lang/elz)
[![Changelog](https://img.shields.io/badge/changelog-changelog-orange.svg)](https://github.com/elz-lang/elz/blob/master/CHANGELOG.md)
[![Contributing](https://img.shields.io/badge/contributing-contributing-blue.svg)](https://github.com/elz-lang/elz/blob/master/CONTRIBUTING.md)
[![GitHub license](https://img.shields.io/github/license/elz-lang/elz.svg)](https://github.com/elz-lang/elz/blob/master/LICENSE)

Elz is a modern programming language focus on production. I hope it can change daily work of us.

## Use

[![asciicast](https://asciinema.org/a/PzTDQQbkMDOGxiq0656dHON3y.png)](https://asciinema.org/a/PzTDQQbkMDOGxiq0656dHON3y)

## Install

```bash
$ git clone https://github.com/elz-lang/elz.git && cd elz
$ cargo install
```

Make sure your `$GOPATH/bin` is one of `$PATH`

### Dependencies

- antlr-runtime for go<br>
`go get github.com/antlr/antlr4/runtime/Go/antlr`
- go-llvm(with LLVM >= 6.0)<br>
`go get -d llvm.org/llvm/bindings/go/llvm`<br>
Then compile it. You can follow [Go bindings](http://llvm.org/svn/llvm-project/llvm/trunk/bindings/go/README.txt)<br>
And see my [suggest](http://routedan.blogspot.com/2017/12/go-binding-llvm.html)
- llvm tools(clang, lli, lld ...)

## Usage

Only for develop just now.

```bash
$ elz source.elz  # elz will compile it to llvm ir then print it.
```
