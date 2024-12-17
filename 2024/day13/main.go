package main

import (
	"bufio"
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
	machines := parse(input)

	totalCost := 0
	for _, machine := range machines {
		totalCost += machine.cost()
	}

	fmt.Println(totalCost)
}

func solve2(input string) {
	machines := parse(input)

	totalCost := 0
	for _, machine := range machines {
		machine.price.X += 10000000000000
		machine.price.Y += 10000000000000
		totalCost += machine.cost()
	}

	fmt.Println(totalCost)
}

func parse(inputFile string) []machine {
	file, err := os.Open(inputFile)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := make([]machine, 0)
	m := machine{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			m = machine{}
		}

		var x, y int
		switch {
		case strings.HasPrefix(line, "Button A"):
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			m.a = *coordinate.New(x, y)

		case strings.HasPrefix(line, "Button B"):
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			m.b = *coordinate.New(x, y)

		case strings.HasPrefix(line, "Prize"):
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			m.price = *coordinate.New(x, y)
			machines = append(machines, m)
		}
	}

	return machines
}
