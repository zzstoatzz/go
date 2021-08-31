package main

import (
	"math/rand"
	"time"
)

type Node struct {
	index    int
	capacity int
	occupied bool
}

type Edge struct {
	start int
	end   int
}

type Geometry struct {
	dimension int
	width     int
	N         int
}

type Graph struct {
	shape Geometry
	edges []Edge
	nodes []Node
}

type PossibilitySet struct {
	n Node
	e Edge
}

type Run struct {
	S int // max cluster size
}

func build(geom Geometry) Graph {
	channel := make(chan PossibilitySet)
	g := Graph{
		shape: geom,
	}
	go geom.possibilities(channel)

	for res := range channel {
		g.edges = append(g.edges, res.e)
		if len(g.edges)%2 == 0 {
			g.nodes = append(g.nodes, res.n)
		}

	}
	return g
}

func (geom Geometry) possibilities(c chan PossibilitySet) {
	defer close(c)
	for from := 0; from < geom.N; from++ {
		res := new(PossibilitySet)
		budget := exp(2, geom.dimension-1)
		res.n = Node{index: from, capacity: budget, occupied: false}
		for d := 0; d < geom.dimension; d++ {
			to := from + exp(geom.width, d)
			res.e = Edge{start: from, end: to % geom.N}
			c <- *res
		}
	}
}

func (g Graph) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(g.edges),
		func(i, j int) {
			g.edges[i], g.edges[j] = g.edges[j], g.edges[i]
		},
	)
}

// func (g Graph) evolve() {
// 	for i, edge := range g.edges {
// 		n1 := g.nodes[]
// 		n2 := edge.end
// 	}
// }
