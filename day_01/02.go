package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

// Not correct yet :'(
func main() {
	file, err := os.Open("input_01.txt")
	check(err)
	defer file.Close()

	replace_map := map[string]string {
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		index := -1
		original := ""
		replacement := ""
		for index != len(line) + 1 {
			index = len(line) + 1

			for k, v := range replace_map {
				char_index := strings.Index(line, k)
	
				if char_index != -1 && char_index < index {
					index = char_index
					original = k
					replacement = v
				}
			}

			if index != len(line) + 1 {
				line = strings.Replace(line, original, replacement, 1)
			}
		}

		first := -1
		last := -1

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == -1 {
					first = int(char) - '0'
				}
				last = int(char) - '0'
			}			
		}

		sum += 10 * first + last
	}

	fmt.Println("Result: ", sum)
}
