package main

type Cell struct {
	North *Cell
	East  *Cell
	South *Cell
	West  *Cell
	Links map[*Cell]bool
	Row   int
	Col   int
}

func NewCell(row, col int) Cell {
	return Cell{nil, nil, nil, nil, make(map[*Cell]bool), row, col}
}

func (c *Cell) Link(cell *Cell, bidi bool) *Cell {
	c.Links[cell] = true
	if bidi {
		cell.Link(c, false)
	}
	return c
}

func (c *Cell) Unlink(cell *Cell, bidi bool) *Cell {
	delete(c.Links, cell)
	if bidi {
		cell.Link(c, false)
	}
	return c
}

func (c *Cell) LinkKeys() []*Cell {
	keys := make([]*Cell, len(c.Links))
	i := 0
	for key := range c.Links {
		keys[i] = key
		i++
	}
	return keys
}

func (c *Cell) Linked(cell *Cell) bool {
	return c.Links[cell]
}
