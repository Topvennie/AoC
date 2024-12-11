package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	stones := parse(input)

	amount := 0
	blinks := 25

	for _, stone := range stones {
		amount += stone.evolveCount(1, blinks)
	}

	fmt.Println(amount)
}

func solve2(input string) {
	stones := parse(input)

	amount := 0
	blinks := 75

	for _, stone := range stones {
		amount += stone.evolveCount(1, blinks)
	}

	fmt.Println(amount)
}

func parse(inputFile string) []stone {
	file, _ := os.ReadFile(inputFile)
	data := strings.TrimSpace(string(file))

	numbers := strings.Split(data, " ")
	stones := make([]stone, 0, len(numbers))
	for _, number := range numbers {
		n, _ := strconv.Atoi(number)
		stones = append(stones, stone(n))
	}

	return stones
}
