# Decode

[![Go Report Card](https://goreportcard.com/badge/github.com/snwfdhmp/decode)](https://goreportcard.com/report/github.com/snwfdhmp/decode) [![Documentation](https://godoc.org/github.com/snwfdhmp/decode?status.svg)](http://godoc.org/github.com/snwfdhmp/decode) [![cover.run](https://cover.run/go/github.com/snwfdhmp/decode.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fsnwfdhmp%2Fdecode)

## Install

Install with `go get github.com/snwfdhmp/decode`.

Import with

```golang
import "github.com/snwfdhmp/decode"
```

## Usage

Decode YAML

```golang
var data interface{}
err := decode.YAML("config.yml", &data)
```

or decode JSON

```golang
var data interface{}
err := decode.JSON("config.json", &data)
```

## Documentation

Complete documentation can be found on [godoc](http://godoc.org/github.com/snwfdhmp/decode)

## Contributors

- [snwfdhmp](http://github.com/snwfdhmp) - author