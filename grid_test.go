package main

import (
	"math/rand"
	"testing"
)

func TestGridSetup(t *testing.T) {
	t.Run("Tests the initilization of a grid",
		func(t *testing.T) {
			grid := NewGrid(2, 2)

			if grid.Rows != 2 {
				t.Errorf("got %d rows, wanted %d rows", grid.Rows, 2)
			}
		})

	t.Run("Tests grid rows are defined properly", func(t *testing.T) {
		grid := NewGrid(2, 2)

		if len(grid.Grid) != 2 {
			t.Errorf("got %d but wanted %d", len(grid.Grid), 2)
		}
	})

	t.Run("Tests grid cols are defined properly", func(t *testing.T) {
		grid := NewGrid(2, 2)

		if len(grid.Grid[0]) != 2 {
			t.Errorf("got %d but wanted %d", len(grid.Grid[0]), 2)
		}
	})

	t.Run("configures North cell", func(t *testing.T) {
		grid := NewGrid(4, 4)

		if grid.Grid[1][1].North != grid.Grid[0][1] {
			t.Errorf("wanted North cell %v but got %v", grid.Grid[1][1], grid.Grid[0][1])
		}
	})

	t.Run("configures East cells", func(t *testing.T) {
		grid := NewGrid(4, 4)

		if grid.Grid[1][1].East != grid.Grid[1][2] {
			t.Errorf("wanted North cell %v but got %v", grid.Grid[1][1], grid.Grid[1][2])
		}
	})

	t.Run("configures South cells", func(t *testing.T) {
		grid := NewGrid(4, 4)

		if grid.Grid[1][1].South != grid.Grid[2][1] {
			t.Errorf("wanted North cell %v but got %v", grid.Grid[1][1], grid.Grid[2][1])
		}
	})

	t.Run("configures West cells", func(t *testing.T) {
		grid := NewGrid(4, 4)

		if grid.Grid[1][1].West != grid.Grid[1][0] {
			t.Errorf("wanted North cell %v but got %v", grid.Grid[1][1], grid.Grid[1][0])
		}
	})

	t.Run("returns a cell if not out of bounds", func(t *testing.T) {
		grid := NewGrid(2, 2)

		want := grid.Grid[0][0]
		got := grid.Get(0, 0)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("returns nil if out of bounds", func(t *testing.T) {
		grid := NewGrid(2, 2)

		got := grid.Get(2, 2)

		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}
	})

	t.Run("returns a random cell", func(t *testing.T) {
		grid := NewGrid(4, 4)
		randGen = rand.New(rand.NewSource(42))

		got := grid.RandomCell()
		want := grid.Get(1, 3)

		if got != want {
			t.Errorf("got %v but wanted %v", got, want)
		}
	})

	t.Run("loops over each row in a grid", func(t *testing.T) {
		grid := NewGrid(2, 2)

		got := 0
		for range grid.EachRow() {
			got += 1
		}
		want := 2

		if got != want {
			t.Errorf("got %v row, want %v rows", got, want)
		}
	})

	t.Run("loops over each cell in a grid", func(t *testing.T) {
		grid := NewGrid(2, 2)

		got := 0
		for range grid.EachCell() {
			got += 1
		}
		want := 4

		if got != want {
			t.Errorf("got %v row, want %v rows", got, want)
		}
	})
}
