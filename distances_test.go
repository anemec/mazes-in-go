package main

import "testing"

func TestDistanceExists(t *testing.T) {
	root := NewCell(0, 0)
	distances := NewDistances(root)
	cell := NewCell(1, 0)

	distances.Set(cell, 1)

	if _, ok := distances.Get(cell); !ok {
		t.Errorf("Expected to find root in distances")
	}
}

func TestDistanceValue(t *testing.T) {
	root := NewCell(0, 0)
	distances := NewDistances(root)
	cell := NewCell(1, 0)

	distances.Set(cell, 1)

	value, _ := distances.Get(cell)

	if value != 1 {
		t.Errorf("Expected to find root in distances")
	}
}

func TestDistancesDNE(t *testing.T) {
	root := NewCell(0, 0)
	distances := NewDistances(root)
	cell := NewCell(1, 0)

	if _, exists := distances.Get(cell); exists {
		t.Errorf("Expected to find root in distances")
	}
}
