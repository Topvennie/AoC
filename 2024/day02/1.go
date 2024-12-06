package main

import (
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

func main() {
	// Read
	data, _ := os.ReadFile("input.txt")
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
