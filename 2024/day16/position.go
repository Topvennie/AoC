package main

import "github.com/Topvennie/AoC/pkg/coordinate"

type position struct {
	current  posDir
	previous []posDir

	cost int
	done bool // Reached finish line
}

type posDir struct {
	p coordinate.Coord
	d coordinate.Dir
}

func newPosition() *position {
	return &position{}
}

func (p *position) copy() *position {
	pos := &position{
		current:  p.current,
		previous: []posDir{},
		cost:     p.cost,
		done:     p.done,
	}

	pos.previous = append(pos.previous, p.previous...)

	return pos
}
