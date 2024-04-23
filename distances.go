package main

type Distances struct {
	Root  *Cell
	Cells map[*Cell]int
}

func NewDistances(root *Cell) Distances {
	cells := make(map[*Cell]int)
	cells[root] = 0
	return Distances{
		Root:  root,
		Cells: cells,
	}
}

func (d *Distances) Set(c *Cell, distance int) {
	d.Cells[c] = distance
}

func (d *Distances) Get(c *Cell) (int, bool) {
	value, exists := d.Cells[c]
	return value, exists
}

func (d *Distances) Keys() []*Cell {
	keys := make([]*Cell, 0, len(d.Cells))
	for k := range d.Cells {
		keys = append(keys, k)
	}
	return keys
}

// func (d *Distances) PathTo(goal *Cell) []Cell {
//   current := goal
//   path := []Cell{current}
//   for current != d.Root {
//     for _, neighbor := range current.Links() {
//       if d.Cells[neighbor] < d.Cells[current] {
//         path = append(path, neighbor)
//         current = neighbor
//         break
//       }
//     }
//   }
//   return path
// }
