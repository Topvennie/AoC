package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
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

func solve1(inputFile string) {
	// Read the file
	data, _ := os.ReadFile(inputFile)
	input := string(data)

	// Parse input
	var left, right []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Split line by spaces
		parts := strings.Fields(line)

		// Convert to integers and append to respective arrays
		if len(parts) == 2 {
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			left = append(left, l)
			right = append(right, r)
		}
	}

	// Sort list
	slices.SortFunc(left, func(i, j int) int {
		return i - j
	})
	slices.SortFunc(right, func(i, j int) int {
		return i - j
	})

	// Compare lists
	var difference float64
	for i := range left {
		difference += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Println(int(difference))
}

func solve2(inputFile string) {
	// Read the file
	data, _ := os.ReadFile(inputFile)
	input := string(data)

	// Parse input
	var left, right []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Split line by spaces
		parts := strings.Fields(line)

		// Convert to integers and append to respective arrays
		if len(parts) == 2 {
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			left = append(left, l)
			right = append(right, r)
		}
	}

	slices.SortFunc(right, func(i, j int) int {
		return i - j
	})

	// Compare lists
	difference := 0
	for _, i := range left {
		difference += i * search(right, i)
	}

	fmt.Println(difference)
}

func search(list []int, i int) int {
	amount := 0

	for _, l := range list {
		if l == i {
			amount++
		}

		if l > i {
			break
		}
	}

	return amount
}
