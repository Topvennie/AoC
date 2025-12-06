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

	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(string(file))
		return
	}

	solve2(string(file))
}

func solve1(input string) {
	problems := make([]*problem, 0)

	lines := strings.Split(input, "\n")

	// Remove leading or tailing empty line
	for len(lines[0]) == 0 {
		lines = lines[1:]
	}
	for len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	lineIdx := 0

	for _, line := range lines[:len(lines)-1] {
		parts := strings.SplitSeq(line, " ")
		partIdx := 0

		for part := range parts {
			if part == "" {
				continue
			}

			if partIdx >= len(problems) {
				problems = append(problems, newProblem())
			}

			number, _ := strconv.Atoi(part)
			problems[partIdx].add(number)

			partIdx++
		}

		lineIdx++
	}

	parts := strings.SplitSeq(lines[len(lines)-1], " ")
	partIdx := 0
	total := 0

	for part := range parts {
		if part == "" {
			continue
		}

		var op operator
		if part == "+" {
			op = add
		} else {
			op = multiply
		}

		total += problems[partIdx].solve(op)

		partIdx++
	}

	fmt.Println(total)
}

func solve2(input string) {
	lines := strings.Split(input, "\n")

	// Remove leading or tailing empty lines
	for len(lines[0]) == 0 {
		lines = lines[1:]
	}
	for len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	problems := make([]*problem, 0)

	for i := range len(lines[0]) {
		if lines[len(lines)-1][i] == '*' || lines[len(lines)-1][i] == '+' {
			problems = append(problems, newProblem())
		}

		number := make([]byte, 0)
		for _, line := range lines[:len(lines)-1] {
			if line[i] != ' ' {
				number = append(number, line[i])
			}
		}
		if len(number) > 0 {
			n, _ := strconv.Atoi(string(number))
			problems[len(problems)-1].add(n)
		}
	}

	parts := strings.SplitSeq(lines[len(lines)-1], " ")
	partIdx := 0
	total := 0

	for part := range parts {
		if part == "" {
			continue
		}

		var op operator
		if part == "+" {
			op = add
		} else {
			op = multiply
		}

		total += problems[partIdx].solve(op)
		partIdx++
	}

	fmt.Println(total)
}
