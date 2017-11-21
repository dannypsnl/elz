# elz
[![Build Status](https://travis-ci.org/elz-lang/elz.svg?branch=master)](https://travis-ci.org/elz-lang/elz)
<br>
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
Not yet!
## Usage
```bash
$ elz source.elz  # elz will compile it to llvm ir.
$ elz -c source.elz # -c can  help you compile it to object file. If it contain main function, then that will be a executable.
```
## Contributing
Watch [contributing.md](https://github.com/elz-lang/elz/blob/master/contributing.md)<br>
