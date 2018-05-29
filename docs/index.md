# Introduciton

Elz is a programming language for more productive.

# Quick Start

## Variable

We have two kinds of variable, global & local.

Global variable be defined at module level.

Here has a variable named `a`

```
a = 10
```

Then local? Local variable available in function, method, macro.

```
fn main() {
  let a = 10
}
```

## Function

Function in Elz is a procedure start with keyword `fn`.

```
fn add(l, r: i32) -> i32 { return l + r }
```

This example show that Elz will trying to complete type if you do not define it. At here, `l` be completed by type of `r`.


