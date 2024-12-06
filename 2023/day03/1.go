package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)


type Position struct {
	line int;
	index int;
}

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func nextToSign(lineIndex int, startIndex int, endIndex int, positions []Position) bool {
	for i := 0; i < len(positions); i++ {
		if positions[i].line < lineIndex - 1 {
			continue
		}

		if positions[i].line > lineIndex + 1 {
			return false
		}

		if positions[i].index >= startIndex - 1 && positions[i].index <= endIndex + 1 {
			return true
		}
	}
	
	return false
}

func addToSum(line string, startIndex int, endIndex int, sum *int) {
	result, err := strconv.Atoi(line[startIndex:endIndex])
	check(err)

	(*sum) += result
}

func main()  {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	signs := make([]Position, 0, 50)

	scanner := bufio.NewScanner(file)
	lineIndex := 0;

	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			if ! (c == '.' || unicode.IsDigit(c)) {
				signs = append(signs, Position{lineIndex, i})
			}
		}

		lineIndex++;
	}

	_, err = file.Seek(0, io.SeekStart)
	check(err)

	scanner = bufio.NewScanner(file)
	lineIndex = 0
	startIndex := -1
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			if (! unicode.IsDigit(c)) && startIndex != -1 {
				if nextToSign(lineIndex, startIndex, i - 1, signs) {
					addToSum(line, startIndex, i, &sum)
				}

				startIndex = -1
			} else if unicode.IsDigit(c) && startIndex == -1 {
				startIndex = i
			}
		}

		if startIndex != -1 {
			if nextToSign(lineIndex, startIndex, len(line) - 1, signs) {
				addToSum(line, startIndex, len(line), &sum)
			}
		}
		
		startIndex = -1
		lineIndex++
	}

	fmt.Println("Result: ", sum)
}
