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
	t, coords := parse(input)

	amount := 0
	for _, c := range coords {
		found := make([]coordinate.Coord, 0)
		amount += t.findUnique(c, 1, &found)
	}

	fmt.Println(amount)
}

func solve2(input string) {
	t, coords := parse(input)

	amount := 0
	for _, c := range coords {
		amount += t.find(c, 1)
	}

	fmt.Println(amount)
}

func parse(inputFile string) (trailMap, []coordinate.Coord) {
	file, _ := os.ReadFile(inputFile)
	data := strings.TrimSpace(string(file))
	lines := strings.Split(data, "\n")

	t := new(len(lines[0]), len(lines))
	var coords = []coordinate.Coord{}
	for i, line := range lines {
		fmt.Println(line)
		for j, char := range line {
			number := int(char) - '0'
			t[i] = append(t[i], number)
			if number == 0 {
				coords = append(coords, coordinate.Coord{X: j, Y: i})
			}
		}
	}

	return t, coords
}
