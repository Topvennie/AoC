package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	increasing = iota + 1
	decreasing
)

const maxCorrections = 1

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

func solve1(inputFile string) {
	// Read
	data, _ := os.ReadFile(inputFile)
	input := string(data)

	// Parse
	amount := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)

		// Avoid empty lines
		if len(parts) < 2 {
			continue
		}

		one, _ := strconv.Atoi(parts[0])
		two, _ := strconv.Atoi(parts[1])

		if math.Abs(float64(one)-float64(two)) > 3 || one == two {
			continue
		}

		direction := increasing
		if one > two {
			direction = decreasing
		}

		previous := two

		valid := true
		for _, part := range parts[2:] {
			current, _ := strconv.Atoi(part)
			if current < previous && direction == decreasing {
				// Decreasing
				if previous-current < 4 {
					previous = current
					continue
				}

			} else if current > previous && direction == increasing {
				// Increasing
				if current-previous < 4 {
					previous = current
					continue
				}
			}

			valid = false
			break
		}

		if valid {
			amount++
		}
	}

	fmt.Println(amount)
}

func solve2(inputFile string) {
	// Read
	data, _ := os.ReadFile(inputFile)
	input := string(data)

	// Parse
	amount := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)

		// Avoid empty lines
		if len(parts) < 2 {
			continue
		}

		// Try forwards
		result := runForwards(parts)

		if !result {
			// Didn't work, lets go backwards
			result = runBackwards(parts)
		}

		if result {
			amount++
		}
	}

	fmt.Println(amount)
}

func runForwards(parts []string) bool {
	first, _ := strconv.Atoi(parts[0])
	second, _ := strconv.Atoi(parts[1])

	if math.Abs(float64(first)-float64(second)) > 3 || first == second {
		return false
	}

	direction := increasing
	if first > second {
		direction = decreasing
	}

	previous := second

	valid := true
	corrections := 0

	for _, part := range parts[2:] {
		current, _ := strconv.Atoi(part)
		if current < previous && direction == decreasing {
			// Decreasing
			if previous-current < 4 {
				previous = current
				continue
			}

		} else if current > previous && direction == increasing {
			// Increasing
			if current-previous < 4 {
				previous = current
				continue
			}
		}

		// Use a correction if any left
		if corrections < maxCorrections {
			corrections++
			continue
		}

		valid = false
		break
	}

	return valid
}

func runBackwards(parts []string) bool {
	last, _ := strconv.Atoi(parts[len(parts)-1])
	secondLast, _ := strconv.Atoi(parts[len(parts)-2])

	if math.Abs(float64(last)-float64(secondLast)) > 3 || last == secondLast {
		return false
	}

	direction := increasing
	if secondLast > last {
		direction = decreasing
	}

	previous := secondLast

	valid := true
	corrections := 0

	for i := len(parts) - 3; i >= 0; i-- {
		current, _ := strconv.Atoi(parts[i])
		if current < previous && direction == increasing {
			// Increasing
			if previous-current < 4 {
				previous = current
				continue
			}
		} else if current > previous && direction == decreasing {
			// Decreasing
			if current-previous < 4 {
				previous = current
				continue
			}
		}

		// Use a correction if any left
		if corrections < maxCorrections {
			corrections++
			continue
		}

		valid = false
		break
	}

	return valid
}
