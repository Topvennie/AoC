package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Read the file
	data, _ := os.ReadFile("text.txt")
	input := string(data)

	// Parse input
	var left, right []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Split line by spaces
		parts := strings.Fields(line)

		// Convert to integers and append to respective arrays
		if len(parts) == 2 {
			l, _ := strconv.Atoi(parts[0])
			r, _ := strconv.Atoi(parts[1])
			left = append(left, l)
			right = append(right, r)
		}
	}

	// Sort list
	slices.SortFunc(left, func(i, j int) int {
		return i - j
	})
	slices.SortFunc(right, func(i, j int) int {
		return i - j
	})

	// Compare lists
	var difference float64
	for i := range left {
		difference += math.Abs(float64(left[i] - right[i]))
	}

	fmt.Println(int(difference))
}
