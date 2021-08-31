package main

import (
	"fmt"
	"math"
	"time"
)

func exp(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}

// var dict = make(map[int]int)

// func findRoot(n int, ref []int) {
// 	if ref
// }

func (g Graph) summarize(elapsed time.Duration) {
	fmt.Printf("took %s", elapsed)
}
