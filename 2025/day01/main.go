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

	dial := dial{
		current: 50,
	}

	for _, line := range lines {
		instr := parseInstruction(line)

		dial.add(instr)
	}

	fmt.Printf("Zero amount: %d\n", dial.zeros)
}

func solve2(input string) {
	lines := strings.Split(input, "\n")

	dial := dial{
		current: 50,
	}

	for _, line := range lines {
		instr := parseInstruction(line)

		for range instr.amount {
			dial.add(instruction{
				amount: 1,
				dir:    instr.dir,
			})
		}
	}

	fmt.Printf("Zero amount: %d\n", dial.zeros)
}
