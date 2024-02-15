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

	t.Run("it links the correct cell", func(t *testing.T) {
		c := NewCell(0, 0)
		want := NewCell(1, 1)

		c.Link(&want, true)

		if c.Linked(&want) != want.Linked(&c) {
			t.Errorf("cells were not linked properly, got %v but wanted %v",
				c.Linked(&want), want)
		}
	})

	t.Run("it unlinks the correct cell", func(t *testing.T) {
		c := NewCell(0, 0)
		want := NewCell(0, 1)

		c.Link(&want, true)
		c.Unlink(&want, true)

		if c.Linked(&want) != false {
			t.Errorf("cells were not linked properly, got %v but wanted %v",
				c.Linked(&want), want)
		}
	})
}
