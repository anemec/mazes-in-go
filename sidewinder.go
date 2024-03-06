package main

import (
	"math/rand"
)

func (g *Grid) Sidewinder() {
	for row := range g.EachRow() {
		run := []*Cell{}

		for _, cell := range row {
			run = append(run, cell)

			atEastBoundary := cell.East == nil
			atNorthernBoundary := cell.North == nil

			shouldCloseOut := atEastBoundary || (!atNorthernBoundary && rand.Intn(2) == 0)

			if shouldCloseOut {
				index := rand.Intn(len(run))
				if run[index].North != nil {
					run[index].Link(run[index].North, true)
				}
				run = []*Cell{}
			} else {
				cell.Link(cell.East, true)
			}
		}
	}
}
