package main

import (
	"fmt"
	"math"
	"time"
)

type Edge struct {
	start int
	end   int
}

func exp(base, power float64) int {
	return int(math.Pow(base, power))
}

func possibilities(r, l float64) []Edge {
	var possible []Edge // [[s1, s2_i], [s1, s2_{i+1}], ...]

	N := exp(l, r)

	for i := 0; i < N; i++ {
		for dim := 0; dim < int(r); dim++ {
			to := i + exp(l, float64(dim))

			edge := Edge{start: i, end: to % N}
			possible = append(possible, edge)
		}

	}
	return possible
}

// func randomEdges(edges []) {

// }

func main() {

	const R = 2
	const L = 4

	channel := make(chan []Edge)

	epoch := time.Now()

	go func() {
		channel <- possibilities(R, L)
	}()

	allEdges := <-channel

	elapsed := time.Since(epoch)

	fmt.Println(allEdges)
	fmt.Printf("took %s", elapsed)
}
