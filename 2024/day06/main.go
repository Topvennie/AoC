package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

func main() {
	var part int
	flag.IntVar(&part, "p", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part: ", part)

	if part == 1 {
		solve1()
		return
	}

	solve2()
}

func solve1() {
	tiles, guard := parse()
	field := newField(tiles, guard)
	field.moveGuard()
	fmt.Println("Part 1: ", field.countVisited())
}

func solve2() {
	tiles, guard := parse()

	// Run guard once
	tilesCopy := deepCopyTiles(tiles)
	f := newField(tilesCopy, guard)
	f.moveGuard()
	path := f.getVisited()

	var amount int32
	var wg sync.WaitGroup

	// For each coordinate in the path, try to see if adding an obstacle helps
	for _, coord := range path {
		if tiles[coord.Y][coord.X] == guardTile {
			continue
		}

		wg.Add(1)

		go func(c coordinate.Coord) {
			defer wg.Done()

			tilesCopy := deepCopyTiles(tiles)
			f := newField(tilesCopy, guard)
			f.tiles[c.Y][c.X] = obstacle

			if f.checkLoop() {
				atomic.AddInt32(&amount, 1)
			}
		}(coord)
	}

	wg.Wait()

	fmt.Println("Part2: ", amount)
}

func parse() ([][]tile, guard) {
	file, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(file))
	lines := strings.Split(data, "\n")

	var tiles [][]tile
	guard := guard{}
	for i, line := range lines {
		tileRow := make([]tile, 0, len(line))
		for j, part := range line {
			switch part {
			case '.':
				tileRow = append(tileRow, empty)
			case '#':
				tileRow = append(tileRow, obstacle)
			case '>':
				tileRow = append(tileRow, guardTile)
				guard.dir = coordinate.Right
				guard.position = coordinate.Coord{X: j, Y: i}
			case '<':
				tileRow = append(tileRow, guardTile)
				guard.dir = coordinate.Left
				guard.position = coordinate.Coord{X: j, Y: i}
			case '^':
				tileRow = append(tileRow, guardTile)
				guard.dir = coordinate.Up
				guard.position = coordinate.Coord{X: j, Y: i}
			case 'v':
				tileRow = append(tileRow, guardTile)
				guard.dir = coordinate.Down
				guard.position = coordinate.Coord{X: j, Y: i}
			}
		}

		tiles = append(tiles, tileRow)
	}

	return tiles, guard
}

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
