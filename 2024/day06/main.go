package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

func main() {
	var part int
	var input string
	flag.IntVar(&part, "p", 1, "Part 1 or 2")
	flag.StringVar(&input, "i", "input.txt", "Location of the input file")
	flag.Parse()
	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(input)
		return
	}

	solve2(input)
}

func solve1(input string) {
	tiles, guard := parse(input)
	field := newField(tiles, guard)
	field.moveGuard()
	fmt.Println("Part 1: ", field.countVisited())
}

func solve2(input string) {
	tiles, guard := parse(input)

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

func parse(input string) ([][]tile, guard) {
	file, _ := os.ReadFile(input)
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
