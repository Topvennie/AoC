package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

type field struct {
	tiles [][]tile
	guard guard
}

type guard struct {
	position coordinate.Coord
	dir      coordinate.Dir
}

var turn = map[coordinate.Dir]coordinate.Dir{
	coordinate.Up:    coordinate.Right,
	coordinate.Right: coordinate.Down,
	coordinate.Down:  coordinate.Left,
	coordinate.Left:  coordinate.Up,
}

type tile int

const (
	guardTile tile = iota + 1
	obstacle
	empty
	visited
)

func newField(tiles [][]tile, g guard) *field {
	return &field{tiles: tiles, guard: g}
}

func (f *field) moveGuard() {
	for f.guardLookingInside() {
		newPos := f.guard.position.Add(coordinate.DirToCoord[f.guard.dir])
		if f.tiles[newPos.Y][newPos.X] != obstacle {
			f.tiles[f.guard.position.Y][f.guard.position.X] = visited
			f.guard.position = *newPos
			f.tiles[f.guard.position.Y][f.guard.position.X] = guardTile
			continue
		}

		f.guard.dir = turn[f.guard.dir]
	}
}

func (f *field) checkLoop() bool {
	var positions []guard

	for f.guardLookingInside() {
		newPos := f.guard.position.Add(coordinate.DirToCoord[f.guard.dir])
		if f.tiles[newPos.Y][newPos.X] != obstacle {
			f.tiles[f.guard.position.Y][f.guard.position.X] = visited
			f.guard.position = *newPos
			f.tiles[f.guard.position.Y][f.guard.position.X] = guardTile
		} else {
			f.guard.dir = turn[f.guard.dir]
		}

		// Check if exact configuration is already present
		for _, g := range positions {
			if g.equal(f.guard) {
				return true
			}
		}

		positions = append(positions, f.guard)
	}

	return false
}

func (f *field) guardLookingInside() bool {
	// Check for X
	switch f.guard.position.X {
	case 0:
		if f.guard.dir == coordinate.Left {
			return false
		}
	case len(f.tiles[0]) - 1:
		if f.guard.dir == coordinate.Right {
			return false
		}
	}

	// Check for Y
	switch f.guard.position.Y {
	case 0:
		if f.guard.dir == coordinate.Up {
			return false
		}
	case len(f.tiles) - 1:
		if f.guard.dir == coordinate.Down {
			return false
		}
	}

	return true
}

func (f *field) countVisited() int {
	amount := 0

	for _, row := range f.tiles {
		for _, tile := range row {
			if tile == visited || tile == guardTile {
				amount++
			}
		}
	}

	return amount
}

func (f *field) getVisited() []coordinate.Coord {
	var coords []coordinate.Coord

	for y, row := range f.tiles {
		for x, tile := range row {
			if tile == visited || tile == guardTile {
				coords = append(coords, coordinate.Coord{X: x, Y: y})
			}
		}
	}

	return coords
}

func (f *field) print() {
	stringMap := map[tile]string{
		guardTile: "v",
		obstacle:  "#",
		empty:     ".",
		visited:   "X",
	}

	for _, row := range f.tiles {
		var result strings.Builder
		for _, tile := range row {
			result.WriteString(stringMap[tile])
		}
		fmt.Println(result.String())
	}
}

func (g guard) equal(g2 guard) bool {
	return g.dir == g2.dir && g.position.Equal(g2.position)
}

func deepCopyTiles(tiles [][]tile) [][]tile {
	copy := make([][]tile, len(tiles))
	for i := range tiles {
		copy[i] = slices.Clone(tiles[i])
	}
	return copy
}
