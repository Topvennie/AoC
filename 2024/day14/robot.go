package main

import (
	"math"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

const (
	w = 101
	h = 103
)

type robot struct {
	p coordinate.Coord
	v coordinate.Coord
}

func (r *robot) simulate(amount int) {
	v := r.v
	v.X *= amount
	v.Y *= amount

	newPos := r.p.Add(v)
	newPos.X = ((newPos.X % w) + w) % w
	newPos.Y = ((newPos.Y % h) + h) % h

	r.p = *newPos
}

type quadrant int

const (
	no quadrant = iota
	topLeft
	topRight
	bottomRight
	bottomLeft
)

// Returns in which quadrant the robot is in
// Quadrant are numbered clockwise, starting in the top right
// -1 is returned when the robot isn't in any quadrant
func (r *robot) quadrant() quadrant {
	mX := int(math.Floor(w / 2))
	mY := int(math.Floor(h / 2))

	if r.p.X < mX {
		if r.p.Y < mY {
			return topLeft
		} else if r.p.Y > mY {
			return bottomLeft
		}
	}

	if r.p.X > mX {
		if r.p.Y < mY {
			return topRight
		} else if r.p.Y > mY {
			return bottomRight
		}
	}

	return no
}
