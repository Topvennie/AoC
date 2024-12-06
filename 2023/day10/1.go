package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	vertical   = iota
	horizontal = iota
	north_east = iota
	north_west = iota
	south_west = iota
	south_east = iota
	ground     = iota
	start      = iota
)

type coordinate struct {
	y int
	x int
}

var up = coordinate{-1, 0}
var down = coordinate{1, 0}
var right = coordinate{0, 1}
var left = coordinate{0, -1}
var neutral = coordinate{0, 0}

var pipes [][]int

var movement = map[int][]coordinate{
	vertical:   {up, down},
	horizontal: {left, right},
	north_east: {up, right},
	north_west: {up, left},
	south_west: {down, left},
	south_east: {down, right},
	ground:     {neutral, neutral},
	start:      {up, right, down, left},
}

func (coord coordinate) sum(other coordinate) coordinate {
	return coordinate{
		coord.y + other.y,
		coord.x + other.x,
	}
}

func (coord coordinate) inBound() bool {
	return coord.y >= 0 && coord.x >= 0 &&
		coord.y < len(pipes) && coord.x < len(pipes[0])
}

func (coord coordinate) equal(other coordinate) bool {
	return coord.y == other.y && coord.x == other.x
}

func (coord coordinate) in(coords []coordinate) bool {
	for _, co := range coords {
		if coord.equal(co) {
			return true
		}
	}

	return false
}

func (coord coordinate) getNextLocation(previous coordinate) coordinate {
	for _, mov := range movement[pipes[coord.y][coord.x]] {
		newCoord := coord.sum(mov)

		if !newCoord.equal(previous) {
			return newCoord
		}
	}

	return coordinate{-1, -1}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	pipes = make([][]int, 0, 10)
	for i := range pipes {
		pipes[i] = make([]int, 0, 10)
	}
	lineIndex := 0
	wordIndex := 0
	var startLocation coordinate

	for scanner.Scan() {
		wordIndex = 0
		line := scanner.Text()
		lineElements := make([]int, 0, 10)
		for _, c := range line {
			elementType := 0
			switch c {
			case '|':
				elementType = vertical
			case '-':
				elementType = horizontal
			case 'L':
				elementType = north_east
			case 'J':
				elementType = north_west
			case '7':
				elementType = south_west
			case 'F':
				elementType = south_east
			case 'S':
				elementType = start
				startLocation = coordinate{lineIndex, wordIndex}
			default:
				elementType = ground
			}

			lineElements = append(lineElements, elementType)
			wordIndex++
		}

		pipes = append(pipes, lineElements)
		lineIndex++
	}

	nextLocations := make([]coordinate, 0, 2)

	startUp := startLocation.sum(up)
	startRight := startLocation.sum(right)
	startDown := startLocation.sum(down)
	startLeft := startLocation.sum(left)

	if startUp.inBound() && down.in(movement[pipes[startUp.y][startUp.x]]) {
		nextLocations = append(nextLocations, startUp)
	}
	if startRight.inBound() && left.in(movement[pipes[startRight.y][startRight.x]]) {
		nextLocations = append(nextLocations, startRight)
	}
	if startDown.inBound() && up.in(movement[pipes[startDown.y][startDown.x]]) {
		nextLocations = append(nextLocations, startDown)
	}
	if startLeft.inBound() && right.in(movement[pipes[startLeft.y][startLeft.x]]) {
		nextLocations = append(nextLocations, startLeft)
	}

	steps := 1

	previous1 := startLocation
	next1 := nextLocations[0]

	previous2 := startLocation
	next2 := nextLocations[1]

	for !next1.equal(next2) {
		tmp1 := next1
		tmp2 := next2

		next1 = next1.getNextLocation(previous1)
		next2 = next2.getNextLocation(previous2)

		previous1 = tmp1
		previous2 = tmp2

		steps++
	}

	fmt.Println("Result: ", steps)
}
