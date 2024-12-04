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
	// Loop over data
	amount := 0
	for i := range d.data {
		for j := range d.data[i] {
			coord := coordinate{x: i, y: j}
			// Check X
			if !d.checkLetter(coord, 'X') {
				continue
			}

			coordirs := make([]coordir, 0, len(directions))
			for _, dir := range directions {
				coordirs = append(coordirs, coordir{coord: coord, dir: dir})
			}

			// Check M
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'M')
			// Check A
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'A')
			// Check S
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'S')

			amount += len(coordirs)
		}
	}

	return amount
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

func (d *data) allNeighbours(c []coordir) []coordir {
	var coordirs []coordir
	for _, coordir := range c {
		newCoordir, valid := d.neighbour(coordir)
		if valid {
			coordirs = append(coordirs, newCoordir)
		}
	}

	return coordirs
}

func (d *data) neighbour(c coordir) (coordir, bool) {
	newCoord := c.coord.add(dirToCoord[c.dir])
	if d.inside(*newCoord) {
		return coordir{coord: *newCoord, dir: c.dir}, true
	}

	return c, false
}

func (d *data) checkLetters(coordirs []coordir, letter rune) []coordir {
	var coords []coordir
	for _, cd := range coordirs {
		if d.checkLetter(cd.coord, letter) {
			coords = append(coords, cd)
		}
	}

	return coords
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
	up direction = iota + 1
	right
	down
	left
	topRight
	bottomRight
	bottomLeft
	topLeft
)

var directions = []direction{up, right, down, left, topRight, bottomRight, bottomLeft, topLeft}

var dirToCoord = map[direction]coordinate{
	up:          {x: 0, y: -1},
	right:       {x: -1, y: 0},
	down:        {x: 0, y: 1},
	left:        {x: 1, y: 0},
	topRight:    {x: 1, y: -1},
	bottomRight: {x: 1, y: 1},
	bottomLeft:  {x: -1, y: 1},
	topLeft:     {x: -1, y: -1},
}

func (c *coordinate) add(c2 coordinate) *coordinate {
	return &coordinate{x: c.x + c2.x, y: c.y + c2.y}
}
