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

	tach := newTachyon(len(lines[0]))

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case 'S':
				tach.addbeam(j, i)
			case '^':
				tach.addSplitter(j, i)
			}
		}
	}

	for range len(lines) - 1 {
		tach.extend()
	}

	fmt.Println(tach.splits)
}

func solve2(input string) {
	lines := strings.Split(input, "\n")

	tach := newTachyonQuantum(len(lines[0]))

	for i, line := range lines {
		for j, r := range line {
			switch r {
			case 'S':
				tach.addbeam(j, i)
			case '^':
				tach.addSplitter(j, i)
			}
		}
	}

	for range len(lines) - 1 {
		tach.extend()
	}

	total := 0
	for _, v := range tach.beams[tach.level] {
		total += v
	}

	fmt.Println(total)
}
