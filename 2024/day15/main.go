package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

var reg = regexp.MustCompile("[0-9]+")

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
	w := parse(input, false)
	w.move()
	fmt.Println(w.score())
}

func solve2(input string) {
	w := parse(input, false)
	w.moveWide()
	fmt.Println(w.scoreWide())
}

func parse(inputFile string, wide bool) *warehouse {
	file, _ := os.ReadFile(inputFile)
	data := strings.TrimSpace(string(file))
	parts := strings.Split(data, "\n\n")

	warehouse := new()

	// Parse map
	if wide {
		parseMap(warehouse, parts[0])
	} else {
		parseMapWide(warehouse, parts[0])
	}

	// Parse movements
	for _, char := range parts[1] {
		switch char {
		case '>':
			warehouse.movements = append(warehouse.movements, coordinate.Right)
		case 'v':
			warehouse.movements = append(warehouse.movements, coordinate.Down)
		case '<':
			warehouse.movements = append(warehouse.movements, coordinate.Left)
		case '^':
			warehouse.movements = append(warehouse.movements, coordinate.Up)
		}
	}

	return warehouse
}

func parseMap(w *warehouse, m string) {
	lines := strings.Split(m, "\n")
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			switch lines[y][x] {
			case 'O':
				w.crates[*coordinate.New(x-1, y-1)] = true
			case '@':
				w.robot = *coordinate.New(x-1, y-1)
			case '#':
				w.walls[*coordinate.New(x-1, y-1)] = true
			}
		}
	}
	w.h = len(lines) - 2
	w.w = len(lines[0]) - 2
}

func parseMapWide(w *warehouse, m string) {
	lines := strings.Split(m, "\n")
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			switch lines[y][x] {
			case 'O':
				w.crates[*coordinate.New((x-1)*2, y-1)] = true
			case '@':
				w.robot = *coordinate.New((x-1)*2, y-1)
			case '#':
				w.walls[*coordinate.New((x-1)*2, y-1)] = true
				w.walls[*coordinate.New((x-1)*2+1, y-1)] = true
			}
		}
	}
	w.h = len(lines) - 2
	w.w = (len(lines[0]) - 2) * 2
}
