package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var part int
	var input string
	flag.IntVar(&part, "p", 1, "Part 1 or 2")
	flag.StringVar(&input, "i", "input.txt", "Location of the input file")
	flag.Parse()

	file, _ := os.ReadFile(input)
	data := strings.TrimSpace(string(file))

	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(data)
		return
	}

	solve2(data)
}

func solve1(input string) {
	parts := strings.Split(input, ",")

	ranges := make([]idRange, 0, len(parts))
	for _, part := range parts {
		ranges = append(ranges, parseRange(part))
	}

	result := 0

	for _, r := range ranges {
		results := r.check(true)

		for _, r := range results {
			result += r
		}
	}

	fmt.Printf("Total: %d\n", result)
}

func solve2(input string) {
	parts := strings.Split(input, ",")

	ranges := make([]idRange, 0, len(parts))
	for _, part := range parts {
		ranges = append(ranges, parseRange(part))
	}

	result := 0

	for _, r := range ranges {
		results := r.check(false)

		for _, r := range results {
			result += r
		}
	}

	fmt.Printf("Total: %d\n", result)
}
