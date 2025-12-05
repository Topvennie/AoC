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
	invent := newInventory()

	lines := strings.Split(input, "\n")
	idx := 0

	for len(lines[idx]) > 0 {
		parts := strings.Split(lines[idx], "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		invent.add(ingredient{start: start, end: end})
		idx++
	}

	freshAmount := 0

	for _, line := range lines[idx+1:] {
		id, _ := strconv.Atoi(line)
		if invent.isFresh(id) {
			freshAmount++
		}
	}

	invent.normalize()

	fmt.Println(freshAmount)
}

func solve2(input string) {
	invent := newInventory()

	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		if len(line) == 0 {
			break
		}

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		invent.add(ingredient{start: start, end: end})
	}

	invent.normalize()

	total := invent.total()

	fmt.Println(total)
}
