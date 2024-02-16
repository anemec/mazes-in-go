package main

type Grid struct {
	Rows int
	Cols int
}

func NewGrid(rows, cols int) Grid {
	grid := Grid{
		rows,
		cols,
	}
	return grid
}
