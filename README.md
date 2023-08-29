# enum

[ [📄 docs](https://pkg.go.dev/github.com/orsinium-labs/enum) ] [ [🐙 github](https://github.com/orsinium-labs/enum) ] [ [❤️ sponsor](https://github.com/sponsors/orsinium) ]

Type safe enums for Go without code generation.

😎 Features:

* Type-safe, thanks to generics.
* No code generation.
* Well-documented, with working examples for every function.
* Flexible, supports both static and runtime definitions.
* Zero-dependency.

## 📦 Installation

```bash
go get github.com/orsinium-labs/enum
```

## 🛠️ Usage

Define:

```go
type Color enum.Member[string]

var (
  Red    = Color{"red"}
  Green  = Color{"green"}
  Blue   = Color{"blue"}
  Colors = enum.New(Red, Green, Blue)
)
```

Parse a raw value (`nil` is returned for invalid value):

```go
parsed := Colors.Parse("red")
```

Accept enum members as function arguments:

```go
func SetPixel(x, i int, c Color)
```

Loop over all enum members:

```go
for _, color := range Colors.Members() {
    // ...
}
```

## 🤔 QnA

1. **Does it use reflection?** A bit. Since `Color` isn't a real `enum.Member` but a different type that happened to have the same struct fields, we need to use reflection to extract wrapped values from enum members.
1. **What happens when enums are added in Go itself?** I'll keep it alive until someone uses it but I expect the project popularity to quickly die out when there is a native language support for enums. When you can mess with the compiler itself, you can do more. For example, this package can't provide exhaustiveness check for switch statements using enums (maybe only by implementing a linter) but proper language-level enums would most likely have it.
1. **Is it reliable?** Yes, pretty much. It has good tests but most importantly it's a small project with just a bit of the actual code that is hard to mess up.
1. **Is it maintained?** The project is pretty-much feature-complete, so there is nothing for me to commit and release daily. However, I accept contributions (see below).
1. **What if I found a bug?** Fork the project, fix the bug, write some tests, and open a Pull Request. I usually merge and release any contributions within a day.
