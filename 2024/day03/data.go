package main

import (
	"regexp"
	"strconv"
)

type data struct {
	lines   []string
	enabled bool
}

func newData(lines []string) *data {
	return &data{lines: lines, enabled: true}
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

func (d *data) multiply2() int {
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
