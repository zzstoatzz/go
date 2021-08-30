package main

import (
	"fmt"
)

type Vertex struct {
	key      int
	adjacent []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func (g *Graph) addVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("existing... skipped")
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

func (g *Graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}

	test.addEdge(1, 2)
	test.addEdge(1, 2)
	test.addEdge(6, 2)
	test.addEdge(3, 2)
	test.Print()

}
