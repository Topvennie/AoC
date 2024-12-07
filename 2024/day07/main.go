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
	operators := []func(x, y int) operator{
		newAdd,
		newMultiply,
	}

	lines := parse(input)

	amount := 0
	for _, line := range lines {
		if line.solve(operators) {
			amount += line.result
		}
	}

	fmt.Println(amount)
}

func solve2(input string) {
	operators := []func(x, y int) operator{
		newAdd,
		newMultiply,
		newConcat,
	}

	lines := parse(input)

	amount := 0
	for _, line := range lines {
		if line.solve(operators) {
			amount += line.result
		}
	}

	fmt.Println(amount)
}

type line struct {
	result int
	input  []int
}

func parse(input string) []line {
	file, _ := os.ReadFile(input)
	data := strings.TrimSpace(string(file))
	rows := strings.Split(data, "\n")

	lines := make([]line, 0, len(rows))
	for _, row := range rows {
		idx := strings.Index(row, ":")
		result, _ := strconv.Atoi(row[:idx])
		var numbers []int

		parts := strings.Split(row[idx+2:], " ")
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			numbers = append(numbers, number)
		}

		lines = append(lines, line{result: result, input: numbers})
	}

	return lines
}

func (l *line) solve(operators []func(x, y int) operator) bool {
	if len(l.input) == 0 {
		return false
	}

	if len(l.input) == 1 {
		return l.input[0] == l.result
	}

	results := []int{l.input[0]}

	for i := 1; i < len(l.input); i++ {
		var tmp []int
		for _, result := range results {
			if result > l.result {
				continue
			}

			for _, op := range operators {
				r := op(result, l.input[i]).exec()
				tmp = append(tmp, r)
			}
		}
		results = tmp
	}

	for _, result := range results {
		if result == l.result {
			return true
		}
	}

	return false
}
