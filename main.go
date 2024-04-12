package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Mazes!")
	rows := 20
	cols := 20
	grid := NewGrid(rows, cols)
	bt := BinaryTree{
		grid,
	}

	// bt.On()
	// fmt.Println(bt.g.String())
	// fmt.Println(&bt.g)
	_ = bt
	grid.Sidewinder()
	// fmt.Println(&grid)
	grid.ToPNG(10)
}
