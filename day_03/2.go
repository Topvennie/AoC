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


type NumberPosition struct {
	line int;
	startIndex int;
	endIndex int;
	number int;
}

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}


func addToSum(lineIndex int, index int, numbers []NumberPosition, sum *int) {
	firstNumber := -1

	for i := 0; i < len(numbers); i++ {
		if numbers[i].line < lineIndex - 1 {
			continue
		}

		if numbers[i].line > lineIndex + 1 {
			return
		}

		if index >= numbers[i].startIndex - 1 && index <= numbers[i].endIndex + 1 {
			if firstNumber == -1 {
				firstNumber = numbers[i].number
			} else {
				(*sum) += firstNumber * numbers[i].number
				return
			}
		}
	}
}

func main()  {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	numbers := make([]NumberPosition, 0, 50)

	scanner := bufio.NewScanner(file)
	lineIndex := 0
	startIndex := -1

	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			if (! unicode.IsDigit(c)) && startIndex != -1 {
				number, err := strconv.Atoi(line[startIndex : i])
				check(err)
				numbers = append(numbers, NumberPosition{lineIndex, startIndex, i - 1, number})
				startIndex = -1
			} else if unicode.IsDigit(c) && startIndex == -1 {
				startIndex = i
			}
		}

		if startIndex != -1 {
			number, err := strconv.Atoi(line[startIndex : len(line)])
			check(err)
			numbers = append(numbers, NumberPosition{lineIndex, startIndex, len(line) - 1, number})
		}

		startIndex = -1
		lineIndex++;
	}

	_, err = file.Seek(0, io.SeekStart)
	check(err)

	scanner = bufio.NewScanner(file)
	lineIndex = 0
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i, c := range line {
			if c == '*' {
				addToSum(lineIndex, i, numbers, &sum)
			}
		}

		lineIndex++
	}

	fmt.Println("Result: ", sum)
}
