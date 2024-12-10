package main

import (
	"slices"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

type trailMap [][]int

func new(width, height int) trailMap {
	t := make([][]int, 0, height)
	for range height {
		t = append(t, make([]int, 0, width))

	}
	return trailMap(t)
}

func (t trailMap) findUnique(start coordinate.Coord, expected int, found *[]coordinate.Coord) int {
	amount := 0

	for _, dir := range coordinate.Dirs {
		next := start.Add(coordinate.DirToCoord[dir])
		if t.inside(*next) && t[next.Y][next.X] == expected {
			if expected == 9 {
				if !slices.ContainsFunc(*found, func(pos coordinate.Coord) bool { return pos.Equal(*next) }) {
					*found = append(*found, *next)
					amount++
				}

				continue
			}

			amount += t.findUnique(*next, expected+1, found)
		}
	}

	return amount
}

func (t trailMap) find(start coordinate.Coord, expected int) int {
	amount := 0

	for _, dir := range coordinate.Dirs {
		next := start.Add(coordinate.DirToCoord[dir])
		if t.inside(*next) && t[next.Y][next.X] == expected {
			if expected == 9 {
				amount++
				continue
			}

			amount += t.find(*next, expected+1)
		}
	}

	return amount
}

func (t trailMap) inside(c coordinate.Coord) bool {
	return c.Y >= 0 && c.Y < len(t) && c.X >= 0 && c.X < len(t[0])
}
