package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Edge struct {
	start int
	end   int
}

func exp(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}

func possibilities(r, l int) []Edge {
	var possible []Edge // [[s1, s2_i], [s1, s2_{i+1}], ...]

	N := exp(l, r)

	for i := 0; i < N; i++ {
		for dim := 0; dim < r; dim++ {
			to := i + exp(l, dim)
			edge := Edge{start: i, end: to % N}
			possible = append(possible, edge)
		}

	}
	return possible
}

func main() {

	const R int = 2
	const L int = 4

	channel := make(chan []Edge)

	epoch := time.Now()

	go func() {
		channel <- possibilities(R, L)
	}()

	edges := <-channel

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(edges),
		func(i, j int) { edges[i], edges[j] = edges[j], edges[i] },
	)

	elapsed := time.Since(epoch)
	fmt.Printf("took %s", elapsed)
}
