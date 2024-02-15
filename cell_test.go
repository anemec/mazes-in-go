package main

import "testing"

func TestCellFunctions(t *testing.T) {
	t.Run("it links a cell", func(t *testing.T) {
		c := NewCell(0, 0)

		want := NewCell(1, 1)
		c.Link(&want, true)

		if len(c.LinkKeys()) != 1 {
			t.Errorf("got %d, want %d", len(c.LinkKeys()), 1)
		}
	})
}
