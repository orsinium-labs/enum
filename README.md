# enum

Type safe enums for Go without code generation.

Features:

* Type-safe, thanks to generics.
* No code generation.
* Well-documented, with working examples for every function.
* Flexible, supports both static and runtime definitions.

## Installation

```bash
go get github.com/orsinium-labs/enum
```

## Usage

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
