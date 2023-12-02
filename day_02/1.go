package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"unicode"
)


const Red = 12
const Green = 13
const Blue = 14


func check(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func isInt(s string) bool {
    for _, c := range s {
        if !unicode.IsDigit(c) {
            return false
        }
    }

    return true
}

func inBounds(values map[string] int) bool {
	total := map[string] int {
		"red": Red,
		"green": Green,
		"blue": Blue,
	}

	for k, v := range values {
		if v > total[k] {
			return false
		}
	}

	return true
}

func main()  {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		game := -1
		buff := -1
		highest := map[string] int {
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		for index, word := range words {
			if index == 1 {
				game, err = strconv.Atoi(word[:len(word) - 1])
				check(err)
			} else if index != 0 {
				if isInt(word) {
					buff, err = strconv.Atoi(word)
					check(err)
				} else {
					if strings.ContainsAny(string(word[len(word) -1]), ",;") {
						word = word[:len(word) - 1]
					}

					if buff > highest[word] {
						highest[word] = buff
					}
				}
			}
		}

		if inBounds(highest) {
			sum += game
		}
	}

	fmt.Println("Result: ", sum)
}