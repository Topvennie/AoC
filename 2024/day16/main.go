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
	maze := parse(input)
	minCost, _ := maze.solve()
	fmt.Println(minCost)
}

func solve2(input string) {
	maze := parse(input)
	_, solutions := maze.solve()
	fmt.Println(maze.spots(solutions))
}

func parse(inputFile string) *maze {
	file, _ := os.ReadFile(inputFile)
	data := strings.TrimSpace(string(file))
	lines := strings.Split(data, "\n")

	maze := newMaze()

	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			switch lines[y][x] {
			case '#':
				maze.walls[*coordinate.New(x-1, y-1)] = true
			case 'S':
				maze.start = *coordinate.New(x-1, y-1)
			case 'E':
				maze.end = *coordinate.New(x-1, y-1)
			}
		}
	}
	maze.h = len(lines) - 2
	maze.w = len(lines[0]) - 2

	return maze
}
