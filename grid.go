package main

import (
	"math/rand"
	"strings"
)

var randGen = rand.New(rand.NewSource(42))

type Grid struct {
	Grid [][]*Cell
	Rows int
	Cols int
}

func NewGrid(rows, cols int) Grid {
	g := prepareGrid(rows, cols)
	g = configureCells(g)
	grid := Grid{
		Grid: g,
		Rows: rows,
		Cols: cols,
	}
	return grid
}

// TODO: Implement Stringer interface

func prepareGrid(rows, cols int) [][]*Cell {
	grid := make([][]*Cell, rows)
	for i := range grid {
		grid[i] = make([]*Cell, cols)
		for j := range grid[i] {
			c := NewCell(i, j)
			grid[i][j] = c
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

func (g *Grid) RandomCell() *Cell {
	row := randGen.Intn(g.Rows)
	col := randGen.Intn(g.Cols)

	return g.Get(row, col)
}

func (g *Grid) Size() int {
	return g.Rows * g.Cols
}

func (g *Grid) EachRow() <-chan []*Cell {
	ch := make(chan []*Cell)
	go func() {
		defer close(ch)
		for _, row := range g.Grid {
			ch <- row
		}
	}()
	return ch
}

func (g *Grid) EachCell() <-chan *Cell {
	ch := make(chan *Cell)
	go func() {
		defer close(ch)
		for row := range g.EachRow() {
			for _, cell := range row {
				ch <- cell
			}
		}
	}()
	return ch
}

func (g *Grid) String() string {
	output := "+" + strings.Repeat("---+", g.Cols) + "\n"

	for row := range g.EachRow() {
		top := "|"
		bottom := "+"

		for _, cell := range row {
			if cell == nil {
				cell = NewCell(-1, -1)
			}

			body := "   "
			eastBoundary := ""
			if cell.Linked(cell.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top += body
			top += eastBoundary

			southBoundary := ""
			if cell.Linked(cell.South) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}

			corner := "+"
			bottom += southBoundary
			bottom += corner

		}

		output += top
		output += "\n"
		output += bottom
		output += "\n"
	}
	return output
}
