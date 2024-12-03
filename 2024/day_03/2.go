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
	d := newData(lines)
	fmt.Println(d.multiply())
}

type data struct {
	lines   []string
	enabled bool
}

func parseInput() []string {
	data, _ := os.ReadFile("input.txt")
	input := string(data)

	lines := strings.Split(input, "\n")
	return lines
}

func newData(lines []string) *data {
	return &data{lines: lines, enabled: true}
}

func (d *data) multiply() int {
	re := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\))+|(do\(\))+|(don't\(\))+`)
	reNumbers := regexp.MustCompile("[0-9]+")

	total := 0

	for _, line := range d.lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			if match[0] == 'd' {
				// Possible enabled switch
				d.enabled = false
				if len(match) == 4 {
					d.enabled = true
				}

				continue
			}

			if !d.enabled {
				continue
			}

			numbers := reNumbers.FindAllString(match, -1)

			for i := 0; i < len(numbers); i += 2 {
				x, _ := strconv.Atoi(numbers[i])
				y, _ := strconv.Atoi(numbers[i+1])
				total += x * y
			}
		}
	}

	return total
}
