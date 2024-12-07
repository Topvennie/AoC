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
	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(input)
		return
	}

	solve2(input)
}

func solve1(input string) {
	lines := parse(input)
	d := newData(lines)
	fmt.Println(d.getXMAS())
}

func solve2(input string) {
	lines := parse(input)
	d := newData(lines)
	fmt.Println(d.getXMASSlope())
}

func parse(inputFile string) []string {
	data, _ := os.ReadFile(inputFile)
	input := string(data)

	lines := strings.Split(input, "\n")
	return lines
}
