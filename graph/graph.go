package main

import (
	"math/rand"
	"time"
)

type Node struct {
	index    int
	capacity int
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

type Result struct {
	n Node
	e Edge
}

func pathFind(geom Geometry) Graph {
	channel := make(chan Result)
	g := Graph{
		shape: geom,
	}
	go geom.possibilities(channel)

	var odd bool = true
	for res := range channel {
		g.edges = append(g.edges, res.e)
		if odd {
			g.nodes = append(g.nodes, res.n)
		}
		odd = !odd

	}
	return g
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

func (geom Geometry) possibilities(c chan Result) {
	defer close(c)
	for from := 0; from < geom.N; from++ {
		res := new(Result)
		budget := exp(2, geom.dimension-1)
		res.n = Node{index: from, capacity: budget}
		for d := 0; d < geom.dimension; d++ {
			to := from + exp(geom.width, d)
			res.e = Edge{start: from, end: to % geom.N}
			c <- *res
		}
	}
}
