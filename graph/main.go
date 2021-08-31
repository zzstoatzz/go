package main

import (
	"fmt"
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

	var G Graph = pathFind(geometry)

	G.shuffle()

	elapsed := time.Since(epoch)
	fmt.Printf("took %s", elapsed)
}
