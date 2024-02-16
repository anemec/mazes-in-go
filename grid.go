package main

type Grid struct {
	Grid [][]*Cell
	Rows int
	Cols int
}

func NewGrid(rows, cols int) Grid {
	g := prepareGrid(rows, cols)
	grid := Grid{
		g,
		rows,
		cols,
	}
	return grid
}

func prepareGrid(rows, cols int) [][]*Cell {
	grid := make([][]*Cell, rows)
	for i := range grid {
		grid[i] = make([]*Cell, cols)
	}
	return grid
}
