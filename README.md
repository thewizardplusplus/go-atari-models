# go-atari-models

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-atari-models?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-atari-models)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-atari-models)](https://goreportcard.com/report/github.com/thewizardplusplus/go-atari-models)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-atari-models.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-atari-models)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-atari-models/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-atari-models)

The library that implements checking and generating of [Atari Go](https://senseis.xmp.net/?AtariGo) moves.

_**Disclaimer:** this library was written directly on an Android smartphone with the AnGoIde IDE._

## Features

- representing the board as an associative array of stones with their positions as keys;
- immutable applicating moves to the board via copying the latter;
- checkings of moves:
  - taking into account self-capture;
- generating of moves via filtering from all possible ones:
  - pseudolegal moves;
  - legal moves (with additional checking for captures);
- encoding in [Smart Game Format](https://senseis.xmp.net/?SGF):
  - parsing:
    - of a stone color;
    - of a coordinate;
    - of a position;
    - of a board size;
    - of a move;
    - of a board;
  - serialization:
    - of a stone color;
    - of a coordinate;
    - of a position.

## Installation

```
$ go get github.com/thewizardplusplus/go-atari-models
```

## Examples

`atarimodels.Board.CheckMove()` with legal self-capture:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
)

func main() {
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	// | |B|W| | |
	// +-+-+-+-+-+
	// |B|W|X|W| |
	// +-+-+-+-+-+
	// | |B|W| | |
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	board := models.NewBoard(models.Size{Width: 5, Height: 5})
	for _, move := range []models.Move{
		{Color: models.Black, Point: models.Point{Column: 1, Row: 1}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 1}},
		{Color: models.Black, Point: models.Point{Column: 0, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 1, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 3, Row: 2}},
		{Color: models.Black, Point: models.Point{Column: 1, Row: 3}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 3}},
	} {
		board = board.ApplyMove(move)
	}

	move := models.Move{
		Color: models.Black,
		Point: models.Point{Column: 2, Row: 2},
	}
	fmt.Printf("%+v: %v\n", move, board.CheckMove(move))

	// Output: {Color:0 Point:{Column:2 Row:2}}: <nil>
}
```

`atarimodels.Board.CheckMove()` with illegal self-capture:

```go
package main

import (
	"fmt"

	models "github.com/thewizardplusplus/go-atari-models"
)

func main() {
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	// | | |W| | |
	// +-+-+-+-+-+
	// | |W|X|W| |
	// +-+-+-+-+-+
	// | | |W| | |
	// +-+-+-+-+-+
	// | | | | | |
	// +-+-+-+-+-+
	board := models.NewBoard(models.Size{Width: 5, Height: 5})
	for _, move := range []models.Move{
		{Color: models.White, Point: models.Point{Column: 2, Row: 1}},
		{Color: models.White, Point: models.Point{Column: 1, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 3, Row: 2}},
		{Color: models.White, Point: models.Point{Column: 2, Row: 3}},
	} {
		board = board.ApplyMove(move)
	}

	move := models.Move{
		Color: models.Black,
		Point: models.Point{Column: 2, Row: 2},
	}
	fmt.Printf("%+v: %v\n", move, board.CheckMove(move))

	// Output: {Color:0 Point:{Column:2 Row:2}}: self-capture
}
```

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
