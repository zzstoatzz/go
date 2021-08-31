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

type Geometry struct {
	dimension int
	width     int
}

type Graph struct {
	shape Geometry
	edges []Edge
}

func exp(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}

func (geom Geometry) possibilities() []Edge {
	var edges []Edge
	N := exp(geom.width, geom.dimension)

	for i := 0; i < N; i++ {
		for d := 0; d < geom.dimension; d++ {
			to := i + exp(geom.width, d)
			edge := Edge{start: i, end: to % N}
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g Graph) randomize() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(g.edges),
		func(i, j int) {
			g.edges[i], g.edges[j] = g.edges[j], g.edges[i]
		},
	)
}

func main() {

	const R int = 2
	const L int = 4

	channel := make(chan []Edge)

	epoch := time.Now()

	geom := Geometry{
		dimension: R,
		width:     L,
	}
	go func() {
		channel <- geom.possibilities()
	}()
	edges := <-channel

	g := Graph{
		shape: geom,
		edges: edges,
	}
	g.randomize()

	elapsed := time.Since(epoch)
	fmt.Printf("took %s", elapsed)
}
