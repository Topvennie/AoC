package main

import (
	"slices"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

type region struct {
	plant     rune
	locations []coordinate.Coord
}

func (r region) area() int {
	return len(r.locations)
}

func (r region) perimeter() int {
	length := 0

	for _, p := range r.locations {
		length += 4 - r.neighbours(p)
	}

	return length
}

func (r region) sides() int {
	// Scan top to bottom
	slices.SortFunc(
		r.locations,
		func(a coordinate.Coord, b coordinate.Coord) int {
			if a.Y == b.Y {
				return a.X - b.X
			}

			return a.Y - b.Y
		},
	)

	topCount := 0
	previous := make([]coordinate.Coord, 0)
	for i := 0; i < len(r.locations); i++ {
		x := r.locations[i].X
		y := r.locations[i].Y
		// Check if anyone is directly above this location
		above := false
		for j := 0; j < i; j++ {
			if r.locations[j].Y == y-1 && r.locations[j].X == x {
				above = true
				break
			}
		}
		if above {
			continue
		}

		// No one above. Check if someone to the left
		previous = append(previous, r.locations[i])
		if len(previous) > 1 && r.locations[i-1].X == x-1 && r.locations[i-1].Y == y {
			// We have a neighbour, check if they already form a side
			if previous[len(previous)-2].X == x-1 && previous[len(previous)-2].Y == y {
				continue
			}
		}

		topCount++
	}

	// Scan bottom to top
	bottomCount := 0
	previous = nil
	for i := len(r.locations) - 1; i >= 0; i-- {
		x := r.locations[i].X
		y := r.locations[i].Y
		// Check if anyone is directly under this location
		under := false
		for j := len(r.locations) - 1; j > i; j-- {
			if r.locations[j].Y == y+1 && r.locations[j].X == x {
				under = true
				break
			}
		}
		if under {
			continue
		}

		// No one under. Check if someone to the right
		previous = append(previous, r.locations[i])
		if len(previous) > 1 && r.locations[i+1].X == x+1 && r.locations[i+1].Y == y {
			// We have a neighbour, check if they already form a side
			if previous[len(previous)-2].X == x+1 && previous[len(previous)-2].Y == y {
				continue
			}
		}

		bottomCount++
	}

	// Scan left to right
	slices.SortFunc(
		r.locations,
		func(a coordinate.Coord, b coordinate.Coord) int {
			if a.X == b.X {
				return a.Y - b.Y
			}

			return a.X - b.X
		},
	)

	leftCount := 0
	previous = nil
	for i := 0; i < len(r.locations); i++ {
		x := r.locations[i].X
		y := r.locations[i].Y
		// Check if anyone is directly to the left
		left := false
		for j := 0; j < i; j++ {
			if r.locations[j].X == x-1 && r.locations[j].Y == y {
				left = true
				break
			}
		}
		if left {
			continue
		}

		// No one to the left, check if someone is above
		previous = append(previous, r.locations[i])
		if len(previous) > 1 && r.locations[i-1].X == x && r.locations[i-1].Y == y-1 {
			// We have a neighbour, check if they already form a side
			if previous[len(previous)-2].X == x && previous[len(previous)-2].Y == y-1 {
				continue
			}
		}

		leftCount++
	}

	// Scan right to left
	rightCount := 0
	previous = nil
	for i := len(r.locations) - 1; i >= 0; i-- {
		x := r.locations[i].X
		y := r.locations[i].Y
		// Check if anyone is directly to the right
		right := false
		for j := len(r.locations) - 1; j > i; j-- {
			if r.locations[j].X == x+1 && r.locations[j].Y == y {
				right = true
				break
			}
		}
		if right {
			continue
		}

		// No one to the right, check if someone is under
		previous = append(previous, r.locations[i])
		if len(previous) > 1 && r.locations[i+1].X == x && r.locations[i+1].Y == y+1 {
			// We have a neighbour, check if they already form a side
			if previous[len(previous)-2].X == x && previous[len(previous)-2].Y == y+1 {
				continue
			}
		}

		rightCount++
	}

	return topCount + bottomCount + leftCount + rightCount
}

func (r region) neighbours(coord coordinate.Coord) int {
	neighbours := 0
	for _, dir := range coordinate.Dirs {
		neighbour := coord.Add(coordinate.DirToCoord[dir])
		if r.contains(*neighbour) {
			neighbours++
		}
	}

	return neighbours
}

func (r region) contains(coord coordinate.Coord) bool {
	return slices.ContainsFunc(r.locations, func(c coordinate.Coord) bool { return c.Equal(coord) })
}
