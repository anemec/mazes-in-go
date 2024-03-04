package main

import "testing"

func TestCellFunctions(t *testing.T) {
	t.Run("it links a cell", func(t *testing.T) {
		c := NewCell(0, 0)

		want := NewCell(1, 1)
		c.Link(want, true)

		if len(c.LinkKeys()) != 1 {
			t.Errorf("got %d, want %d", len(c.LinkKeys()), 1)
		}
	})

	t.Run("it links the correct cell", func(t *testing.T) {
		c := NewCell(0, 0)
		want := NewCell(1, 1)

		c.Link(want, true)

		if c.Linked(want) != want.Linked(c) {
			t.Errorf("cells were not linked properly, got %v but wanted %v",
				c.Linked(want), want)
		}
	})

	t.Run("it unlinks the correct cell", func(t *testing.T) {
		c := NewCell(0, 0)
		want := NewCell(0, 1)

		c.Link(want, true)
		c.Unlink(want, true)

		if c.Linked(want) != false {
			t.Errorf("cells were not linked properly, got %v but wanted %v",
				c.Linked(want), want)
		}
	})

	t.Run("returns a list of neighbors", func(t *testing.T) {
		c := NewCell(1, 1)
		northCell := NewCell(0, 1)

		c.North = northCell
		got := c.Neighbors()

		if got[0] != northCell {
			t.Errorf("Neighbors does not return the correct cells, got %v want %v",
				got, northCell)
		}
	})

	t.Run("returns a list of neighbors", func(t *testing.T) {
		c := NewCell(1, 1)
		northCell := NewCell(0, 1)
		eastCell := NewCell(1, 2)
		southCell := NewCell(2, 1)
		westCell := NewCell(1, 0)

		c.North = northCell
		c.East = eastCell
		c.South = southCell
		c.West = westCell
		got := c.Neighbors()

		if len(got) != 4 {
			t.Errorf("Neighbors does not return the correct cells, got %v want %v",
				len(got), 4)
		}
	})
}
