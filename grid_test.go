package main

import "testing"

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
}
