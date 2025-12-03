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
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		b := newBank(line)
		total += b.largest(2)
	}

	fmt.Println(total)
}

func solve2(input string) {
	lines := strings.Split(input, "\n")

	total := 0

	for _, line := range lines {
		b := newBank(line)
		total += b.largest(12)
	}

	fmt.Println(total)
}
