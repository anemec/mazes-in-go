package main

type Grid struct {
	Grid [][]*Cell
	Rows int
	Cols int
}

func NewGrid(rows, cols int) Grid {
	g := prepareGrid(rows, cols)
	g = configureCells(g)
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
		for j := range grid[i] {
			c := NewCell(i, j)
			grid[i][j] = &c
		}
	}
	return grid
}

func configureCells(g [][]*Cell) [][]*Cell {
	for i := range g {
		for j := range g[i] {
			if i > 0 {
				g[i][j].North = g[i-1][j]
			}
			if j < len(g[i])-1 {
				g[i][j].East = g[i][j+1]
			}
			if i < len(g)-1 {
				g[i][j].South = g[i+1][j]
			}
			if j > 0 {
				g[i][j].West = g[i][j-1]
			}

		}
	}
	return g
}

func (g *Grid) Get(row, col int) *Cell {
	if (row >= 0 && row < g.Rows) && (col >= 0 && col < g.Cols) {
		return g.Grid[row][col]
	}
	return nil
}
