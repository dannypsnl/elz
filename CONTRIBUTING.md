# Contributing

### Pull Request(PR)

Ensure all tests was pass, here was how I did it, create `.git/hooks/pre-push` with the following content:

```shell
#!/bin/sh

# An example hook script to verify what is about to be pushed.  Called by "git
# push" after it has checked the remote status, but before anything has been
# pushed.  If this script exits with a non-zero status nothing will be pushed.
#
# This hook is called with the following parameters:
#
# $1 -- Name of the remote to which the push is being done
# $2 -- URL to which the push is being done
#
# If pushing without using a named remote those arguments will be equal.
#
# Information about the commits which are being pushed is supplied as lines to
# the standard input in the form:
#
#   <local ref> <local sha1> <remote ref> <remote sha1>
#
# This sample shows how to prevent push of commits where the log message starts
# with "WIP" (work in progress).

set -e

echo "testing before push"
cargo test

exit 0
```

### Code style

I don't care about code style, but to ensure your auto formatter won't conflict with the current formatter, I list formatters are using in the project.
If you found there has languages using in the project but didn't have a formatter, feel free to use the formatter what you use and modify the list and send a PR for this.

Formatters:

- **Rust**: [rustfmt](https://github.com/rust-lang/rustfmt)
- **Markdown**: [prettier](https://prettier.io/)
