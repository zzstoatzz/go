package main

import (
	"fmt"

	"math"
)

func exp(base, power float64) int {
	return int(math.Pow(base, power))
}

func possibilities(r, l float64) [][2]int {
	var possible [][2]int // [[s1, s2_i], [s1, s2_{i+1}], ...]

	N := exp(l, r)

	for i := 0; i < N; i++ {
		for dim := 0; dim < int(r); dim++ {
			t := i + exp(l, float64(dim))
			edge := [2]int{i, t % N}
			possible = append(possible, edge)
		}

	}
	return possible
}

func main() {
	const R = 2
	const L = 4

	channel := make(chan [][2]int)

	go func() {
		channel <- possibilities(R, L)
	}()

	item := <-channel
	fmt.Println("received", item)
}
