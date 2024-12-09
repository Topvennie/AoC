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
	f := parse(input)
	f.order()
	fmt.Println(f.checksum())
}

func solve2(input string) {
	f := parse(input)
	f.orderWhole()
	fmt.Println(f.checksum())
}

func parse(inputFile string) file {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic("File not found")
	}

	line := strings.TrimSpace(string(data))

	id := 0
	freeSpace := false
	file := []block{}

	for _, char := range line {
		amount := int(char - '0')

		if freeSpace {
			file = append(file, block{id: free, amount: amount})
		} else {
			file = append(file, block{id: id, amount: amount})
			id++
		}

		freeSpace = !freeSpace
	}

	return file
}
