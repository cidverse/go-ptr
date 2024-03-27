# go-ptr

> A collection of utility functions to work with pointers, commonly used for optional values.

[![Go Reference](https://pkg.go.dev/badge/github.com/cidverse/go-ptr.svg)](https://pkg.go.dev/github.com/cidverse/go-ptr)

## Installation

```shell
go get -u github.com/cidverse/go-ptr
```

## Usage

_generics_

```go
fmt.Printf("%#v\n", ptr.Ptr("hello world"))
```

_Primitives_

```go
fmt.Printf("%#v\n", ptr.String("hello world"))
fmt.Printf("%#v\n", ptr.Int("hello world"))
```

## License

Released under the [MIT license](./LICENSE).
