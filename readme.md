# elz

[![Build Status](https://travis-ci.org/elz-lang/elz.svg)](https://travis-ci.org/elz-lang/elz)

[![asciicast](https://asciinema.org/a/229973.svg)](https://asciinema.org/a/229973)

## Install

```bash
$ git clone https://github.com/elz-lang/elz.git && cd elz
$ git submodule init && go generate ./...
$ make install
```

Make sure your `$GOPATH/bin` in `$PATH`

## Usage

```bash
$ elz compile source.elz
```

## Example

> Warning, this is nota stable example, that might outdated anytime with developing

```
main = printf("hello, world\n")
```

## Contributing
