package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := parseInput()
	d := data{lines: lines}
	fmt.Println(d.multiply())
}

type data struct {
	lines []string
}

func parseInput() []string {
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	lines := strings.Split(input, "\n")
	return lines
}

func (d *data) multiply() int {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	reNumbers := regexp.MustCompile("[0-9]+")

	total := 0

	for _, line := range d.lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			numbers := reNumbers.FindAllString(match, -1)
			if len(numbers) != 2 {
				continue
			}

			x, _ := strconv.Atoi(numbers[0])
			y, _ := strconv.Atoi(numbers[1])
			total += x * y
		}
	}

	return total
}
