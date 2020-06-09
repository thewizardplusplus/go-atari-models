# go-atari-models

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-atari-models?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-atari-models)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-atari-models)](https://goreportcard.com/report/github.com/thewizardplusplus/go-atari-models)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-atari-models.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-atari-models)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-atari-models/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-atari-models)

The library that implements checking and generating of [Atari Go](https://senseis.xmp.net/?AtariGo) moves.

_**Disclaimer:** this library was written directly on an Android smartphone with the AnGoIde IDE._

## Installation

```
$ go get github.com/thewizardplusplus/go-atari-models
```

## Examples

`sgf.DecodePoint()`:

```go
package main

import (
	"fmt"

	"github.com/thewizardplusplus/go-atari-models/encoding/sgf"
)

func main() {
	point, _ := sgf.DecodePoint("eE")
	fmt.Printf("%+v\n", point)

	// Output: {Column:4 Row:30}
}
```

`sgf.EncodePoint()`:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
	"github.com/thewizardplusplus/go-atari-models/encoding/sgf"
)

func main() {
	point := sgf.EncodePoint(models.Point{Column: 4, Row: 30})
	fmt.Printf("%s\n", point)

	// Output: eE
}
```

`sgf.DecodeStoneStorage()`:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
	"github.com/thewizardplusplus/go-atari-models/encoding/sgf"
)

func main() {
	storage, _ := sgf.DecodeStoneStorage(
		"(;FF[4]SZ[7:9]GN[test];B[aa](;W[gi]N[test]))",
		models.NewBoard,
	)

	var moves []models.Move
	for _, point := range storage.Size().Points() {
		if color, ok := storage.Stone(point); ok {
			moves = append(moves, models.Move{Color: color, Point: point})
		}
	}

	fmt.Printf("%+v\n", storage.Size())
	fmt.Printf("%+v\n", moves)

	// Output:
	// {Width:7 Height:9}
	// [{Color:0 Point:{Column:0 Row:0}} {Color:1 Point:{Column:6 Row:8}}]
}
```

## License

The MIT License (MIT)

Copyright &copy; 2019-2020 thewizardplusplus
