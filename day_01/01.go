package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Fatalf("Error: %v", e)
	}
}

func main() {
	file, err := os.Open("input_01.txt")
	check(err)
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		first := -1
		last := -1
		for _, char := range scanner.Text() {
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
