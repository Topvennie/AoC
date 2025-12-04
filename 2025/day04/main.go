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

	file, _ := os.ReadFile(input)
	data := strings.TrimSpace(string(file))

	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(data)
		return
	}

	solve2(data)
}

func solve1(input string) {
	lines := strings.Split(input, "\n")

	total := 0

	above := ""
	middle := ""
	below := ""

	for _, line := range lines {
		above = middle
		middle = below
		below = line

		total += len(accessibleIdx(above, middle, below))
	}

	above = middle
	middle = below
	below = ""

	total += len(accessibleIdx(above, middle, below))

	fmt.Println(total)
}

func solve2(input string) {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 {
		return
	}

	total := checkRows(0, 0, lines)

	fmt.Println(total)
}

func checkRows(startIdx, current int, rows []string) int {
	if startIdx < 0 {
		// Can happen if the first row has some changes
		return checkRows(0, current, rows)
	}
	if startIdx >= len(rows) {
		// Marks the end
		return current
	}

	above := ""
	if startIdx > 0 {
		above = rows[startIdx-1]
	}
	middle := rows[startIdx]
	below := ""
	if startIdx < len(rows)-1 {
		below = rows[startIdx+1]
	}

	toilets := accessibleIdx(above, middle, below)

	// No toilets found
	// Go to the next row
	if len(toilets) == 0 {
		return checkRows(startIdx+1, current, rows)
	}

	// Found some toilets
	// Mark them as moved and go back a row
	row := []byte(rows[startIdx])
	for _, toilet := range toilets {
		row[toilet] = '.'
	}
	rows[startIdx] = string(row)

	return checkRows(startIdx-1, current+len(toilets), rows)
}

func accessibleIdx(above, middle, below string) []int {
	toilets := []int{}

	for i := range middle {
		if middle[i] != '@' {
			continue
		}

		if toiletCount(i, above, middle, below) < 4 {
			toilets = append(toilets, i)
		}
	}

	return toilets
}

func toiletCount(idx int, above, middle, below string) int {
	toilets := 0

	// Check row above
	if isToilet(idx-1, above) {
		toilets++
	}
	if isToilet(idx, above) {
		toilets++
	}
	if isToilet(idx+1, above) {
		toilets++
	}

	// Check middle row
	if isToilet(idx-1, middle) {
		toilets++
	}
	if isToilet(idx+1, middle) {
		toilets++
	}

	// Check row below
	if isToilet(idx-1, below) {
		toilets++
	}
	if isToilet(idx, below) {
		toilets++
	}
	if isToilet(idx+1, below) {
		toilets++
	}

	return toilets
}

func isToilet(idx int, row string) bool {
	if idx < 0 {
		return false
	}

	if idx >= len(row) {
		return false
	}

	return row[idx] == '@'
}
