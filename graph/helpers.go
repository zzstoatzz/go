package main

import (
	"math"
)

func exp(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}

// func (g Graph) print() string {
// 	for n := range g.nodes {
// 		fmt.Println(n)
// 	}
// }
