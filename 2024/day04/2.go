package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := parseInput()
	d := newData(lines)
	fmt.Println(d.getXMAS())
}

func parseInput() []string {
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	lines := strings.Split(input, "\n")
	return lines
}

type data struct {
	data [][]rune
}

func newData(lines []string) *data {
	d := make([][]rune, 0, len(lines))

	for _, line := range lines {
		d = append(d, []rune(line))
	}

	return &data{data: d}
}

func (d *data) getXMAS() int {
	// Get all mas

	amount := 0
	for i := range d.data {
		for j := range d.data[i] {
			coord := coordinate{x: i, y: j}
			// Check A
			if !d.checkLetter(coord, 'A') {
				continue
			}

			if d.checkSlope(coord, topLeft, bottomRight) && d.checkSlope(coord, topRight, bottomLeft) {
				amount++
			}
		}
	}

	return amount
}

func (d *data) checkSlope(c coordinate, dir1 direction, dir2 direction) bool {
	opposite := map[rune]rune{
		'M': 'S',
		'S': 'M',
	}

	pos1 := c.add(dirToCoord[dir1])
	if !d.inside(*pos1) {
		return false
	}

	pos2 := c.add(dirToCoord[dir2])
	if !d.inside(*pos2) {
		return false
	}

	rune1 := d.data[pos1.x][pos1.y]
	rune2, ok := opposite[rune1]
	if !ok {
		return false
	}

	return rune2 == d.data[pos2.x][pos2.y]
}

func (d *data) inside(c coordinate) bool {
	if c.x < 0 || c.x >= len(d.data) {
		return false
	}

	if c.y < 0 || c.y >= len(d.data[c.x]) {
		return false
	}

	return true
}

func (d *data) checkLetter(c coordinate, letter rune) bool {
	if c.x < 0 || c.x >= len(d.data) {
		return false
	}

	if c.y < 0 || c.y >= len(d.data[c.x]) {
		return false
	}

	return d.data[c.x][c.y] == letter
}

type coordir struct {
	coord coordinate
	dir   direction
}

type coordinate struct {
	x int
	y int
}

type direction int

const (
	topRight direction = iota + 1
	bottomRight
	bottomLeft
	topLeft
)

var dirToCoord = map[direction]coordinate{
	topRight:    {x: 1, y: -1},
	bottomRight: {x: 1, y: 1},
	bottomLeft:  {x: -1, y: 1},
	topLeft:     {x: -1, y: -1},
}

func (c *coordinate) add(c2 coordinate) *coordinate {
	return &coordinate{x: c.x + c2.x, y: c.y + c2.y}
}
