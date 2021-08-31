package main

import (
	"time"
)

const R int = 2
const L int = 4

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
	G.summarize(time.Since(epoch))
}
