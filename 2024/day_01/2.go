package main

import (
	"fmt"
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

	slices.SortFunc(right, func(i, j int) int {
		return i - j
	})

	// Compare lists
	difference := 0
	for _, i := range left {
		difference += i * search(right, i)
	}

	fmt.Println(difference)
}

func search(list []int, i int) int {
	amount := 0

	for _, l := range list {
		if l == i {
			amount++
		}

		if l > i {
			break
		}
	}

	return amount
}
