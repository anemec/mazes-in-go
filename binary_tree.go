package main

import "math/rand"

type BinaryTree struct {
	g Grid
}

// TODO: Add tests
func (bt BinaryTree) On() {
	for cell := range bt.g.EachCell() {
		neighbors := []*Cell{}

		if cell.North != nil {
			neighbors = append(neighbors, cell.North)
		}
		if cell.East != nil {
			neighbors = append(neighbors, cell.East)
		}

		if len(neighbors) > 0 {
			index := rand.Intn(len(neighbors))

			neighbor := neighbors[index]

			cell.Link(neighbor, true)
		}
	}
}
