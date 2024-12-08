package main

import (
	"github.com/Topvennie/AoC/pkg/coordinate"
)

type field struct {
	width     int
	height    int
	antennas  map[rune][]coordinate.Coord
	antinodes map[coordinate.Coord]bool
}

type tile int

const (
	empty tile = iota + 1
	antenna
	antinode
)

func newField(width, height int, antennas map[rune][]coordinate.Coord) *field {
	return &field{
		width:     width,
		height:    height,
		antennas:  antennas,
		antinodes: make(map[coordinate.Coord]bool),
	}
}

func (f *field) placeAntinodes() {
	for _, antennas := range f.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Get the movement needed to go from one antenna to another
				mov := antennas[i].Sub(antennas[j])

				// Check each location
				loc := antennas[i].Add(*mov)
				if f.inside(*loc) {
					f.antinodes[*loc] = true
				}

				loc = antennas[j].Sub(*mov)
				if f.inside(*loc) {
					f.antinodes[*loc] = true
				}
			}
		}
	}
}

func (f *field) placeManyAntinodes() {
	for _, antennas := range f.antennas {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Get the movement needed to go from one antenna to another
				mov := antennas[i].Sub(antennas[j])

				// Check each location
				loc := antennas[i].Add(*mov)
				for f.inside(*loc) {
					f.antinodes[*loc] = true
					loc = loc.Add(*mov)
				}

				loc = antennas[j].Sub(*mov)
				for f.inside(*loc) {
					f.antinodes[*loc] = true
					loc = loc.Sub(*mov)
				}

				// Add antenna locations
				f.antinodes[antennas[i]] = true
				f.antinodes[antennas[j]] = true
			}
		}
	}
}

func (f *field) inside(c coordinate.Coord) bool {
	if c.X < 0 || c.X >= f.width {
		return false
	}

	if c.Y < 0 || c.Y >= f.height {
		return false
	}

	return true
}
