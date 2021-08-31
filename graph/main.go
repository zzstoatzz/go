package main

import (
	"fmt"
	"time"
)

const R int = 3
const L int = 16

var geometry = Geometry{
	dimension: R,
	width:     L,
	N:         exp(L, R),
}

func main() {

	epoch := time.Now()

	// initialize graph with in/deterministic edge/node
	var G Graph = build(geometry)

	// randomize potential edges per swap func
	G.shuffle()

	// summarize
	elapsed := time.Since(epoch)
	fmt.Printf("took %s", elapsed)
}
