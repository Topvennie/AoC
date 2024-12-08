package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	width, height, antennas := parse(input)
	field := newField(width, height, antennas)
	field.placeAntinodes()
	fmt.Println(len(field.antinodes))
}

func solve2(input string) {
	width, height, antennas := parse(input)
	field := newField(width, height, antennas)
	field.placeManyAntinodes()
	fmt.Println(len(field.antinodes))
}

func parse(file string) (int, int, map[rune][]coordinate.Coord) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic("Input file not found")
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	antennas := make(map[rune][]coordinate.Coord)
	for i, line := range lines {
		for j, char := range line {
			if char != '.' && char != '#' {
				coord := coordinate.Coord{X: j, Y: i}
				entry, ok := antennas[char]
				if !ok {
					entry = make([]coordinate.Coord, 0)
				}

				entry = append(entry, coord)
				antennas[char] = entry
			}
		}

	}

	return len(lines[0]), len(lines), antennas
}
