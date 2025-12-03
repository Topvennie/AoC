package main

import (
	"strconv"
	"strings"
)

type bank struct {
	batteries []rune
}

func newBank(batteriesStr string) *bank {
	batteries := make([]rune, 0, len(batteriesStr))

	for _, char := range batteriesStr {
		batteries = append(batteries, char)
	}

	b := &bank{
		batteries: batteries,
	}

	return b
}

func (b *bank) largest(amount int) int {
	runes := make([]string, 0, amount)

	currentIdx := 0
	for range amount {
		idx := findLargestIndex(b.batteries[currentIdx : len(b.batteries)-(amount-len(runes)-1)])
		runes = append(runes, string(b.batteries[currentIdx+idx]))

		currentIdx += idx + 1
	}

	resultStr := strings.Join(runes, "")
	result, _ := strconv.Atoi(resultStr)

	return result
}

func findLargestIndex(numbers []rune) int {
	idx := 0
	var largest rune = '/'

	for i, number := range numbers {
		if number > largest {
			largest = number
			idx = i
		}
	}

	return idx
}
