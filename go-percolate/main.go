package main

import (
	"fmt"
	"math/rand"
)

var summary string

// N ~ num total nodes
// d ~ num spatial dimensions

const L = 10
const d = 1
const N = L * d
const V = 3 * N

var shape = make(map[int]int, L)

func matrix(rows, cols int) [][]int {
    m := make([][]int, rows)
    for row := range m {
        m[row] = make([]int, cols)
    }
    return m
}

type Tree struct {
	connections []bond
}

type bond struct {
	head int
	tail int
}

func potentialBonds() Tree {
	var t *Tree
	var connections [V]bond
	for r, width := range shape {
		var plane [width]int
		

	return bonds
}

func connectPair(i int) int {

}

func findRoot() {

}

func cluster(bonds [N]bond) [N]int {
	var P [N]int

	for i := 0; i < N; i++ {
		P[i] = connectPair(i)
	}

}

func main() {
	events := potentialBonds()
	bonds := events.connections

	rand.Shuffle(
		len(bonds), func(i, j int) { bonds[i], bonds[j] = bonds[j], bonds[i] },
	)

	fmt.Println(summary)
}
