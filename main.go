package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Mazes!")
	grid := NewGrid(4, 4)
	bt := BinaryTree{
		grid,
	}

	bt.On()
	// fmt.Println(bt.g.String())
	fmt.Println(&bt.g)
}
