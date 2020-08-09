# Go Arithmetic Expression Parser

**aparser** provides a library to build arithmetic expression parsers. It also provides a pre-build arithmetic expression parser for common arithmetic expressions.

The pre-build arithmetic expression parser supports the following operations:

- unary: -
- binary: +, -, *, /, ^ (power), % (modulo)
- ternary: cond ? a : b, if (cond) a else b
- logic: &&, ||, !, ^ (exclusive or)
- relational: ==, !=, <, <=, >, >=
- functions: abs, acos, asin, atan, cos, cosh, exp, log, log10, sign, sin, sinh, sqrt, tan, tanh, round 

# Design Goals

The focus of this library is to demonstrate how to write a simple arithmetic expression parser in Go. It is neither complete, nor optimized for performance.

# Usage

To use the pre-build arithmetic expression parser, simply call the Eval function:

```go
r, err := Eval("3 + 4")
if err != nil {
    panic(err)
}
fmt.Println(r)
```

# Thanks

Thanks to [Rüdiger Klaehn](https://github.com/rklaehn) who wrote the original C# arithmetic expression parser, from which I learned how to write parsers.
It also inspired and guided me to write this code.

# License

BSD 3-Clause License