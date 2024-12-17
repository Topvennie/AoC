package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"

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
	robots := parse(input)
	quadrants := make([]quadrant, 5)

	for _, r := range robots {
		r.simulate(100)
		quadrants[r.quadrant()]++
	}

	total := 1
	for i := 1; i < len(quadrants); i++ {
		total *= int(quadrants[i])
	}

	fmt.Println(total)
}

func solve2(input string) {
	robots := parse(input)

	amount := 10000           // Arbitrary number
	iter := 0                 // The iteration with lowest danger score
	danger := math.MaxInt - 1 // Lowest danger score
	quadrants := []int{0, 0, 0, 0, 0}

	for i := range amount {
		quadrants = []int{0, 0, 0, 0, 0}
		// Simulate all robots by one second
		for i := range robots {
			robots[i].simulate(1)
			quadrants[robots[i].quadrant()]++
		}

		total := 1
		for j := 1; j < len(quadrants); j++ {
			total *= int(quadrants[j])
		}

		if total < danger {
			iter = i
			danger = total
		}
	}

	fmt.Printf("iteration: %d with danger score: %d\n", iter+1, danger)
}

func parse(inputFile string) []robot {
	file, err := os.Open(inputFile)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	robots := make([]robot, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		r := robot{}
		var cX, cY, pX, pY int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &cX, &cY, &pX, &pY)
		r.p = *coordinate.New(cX, cY)
		r.v = *coordinate.New(pX, pY)

		robots = append(robots, r)
	}

	return robots
}
