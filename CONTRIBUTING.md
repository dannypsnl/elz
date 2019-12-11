# Contributing

### Pull Request(PR)

Ensure all tests was pass, I use git hook for this, if you want to use it, run:

```bash
cp ./hooks/pre-push .git/hooks/pre-push
```

### Code style

I don't care about code style, but to ensure your auto formatter won't conflict with the current formatter, I list formatters are using in the project.
If you found there has languages using in the project but didn't have a formatter, feel free to use the formatter what you use and modify the list and send a PR for this.

Formatters:

- **Rust**: [rustfmt](https://github.com/rust-lang/rustfmt)
- **Markdown**: [prettier](https://prettier.io/)
