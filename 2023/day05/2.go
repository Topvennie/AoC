package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TranslationRange struct {
	source      int
	destination int
	length      int
}

type Range struct {
	start  int
	length int
}

func getNumbers(lines []string) (numbers []int) {
	numbers = make([]int, 0, 10)

	for _, line := range lines {
		values := strings.Fields(line)

		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	return
}

func processSeeds(numbers []int) (seeds []Range) {
	seeds = make([]Range, 0)

	for i := 0; i < len(numbers); i += 2 {
		seeds = append(seeds, Range{
			numbers[i],
			numbers[i+1] - 1,
		})
	}

	return
}

func processMap(translations *[]TranslationRange, numbers []int) {
	*translations = make([]TranslationRange, 0, 10)

	for i := 0; i < len(numbers); i += 3 {
		*translations = append(*translations, TranslationRange{
			numbers[i+1],
			numbers[i],
			numbers[i+2] - 1,
		})
	}
}

func getTranslationDestination(current int, tr TranslationRange) int {
	diff := current - tr.source

	return tr.destination + diff
}

func getTranslation(current int, translations []TranslationRange) int {
	if len(translations) == 0 {
		return current
	}

	for _, tr := range translations {
		if current >= tr.source && current <= tr.source+tr.length {
			return getTranslationDestination(current, tr)
		}
	}

	return current
}

func main() {
	file, _ := os.ReadFile("input.txt")

	input := string(file)

	sections := strings.Split(input, "\n\n")
	var seeds []Range
	maps := make([][]TranslationRange, 7)

	for i, section := range sections {
		lines := strings.Split(section, "\n")
		numbers := getNumbers(lines)

		if i == 0 {
			seeds = processSeeds(numbers)
		} else {
			processMap(&maps[i-1], numbers)
		}
	}

	min := int(^uint(0) >> 1)

	for _, seedRange := range seeds {
		var current int
		for i := 0; i <= seedRange.length; i++ {
			current = seedRange.start + i

			for _, translation := range maps {
				current = getTranslation(current, translation)
			}

			if current < min {
				min = current
			}
		}
	}

	fmt.Println("Result: ", min)
}
