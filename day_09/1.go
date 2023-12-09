package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func all[T any](elements []T, it func (T) bool) bool {
    for _, el := range elements {
        if ! it(el) {
            return false
        }
    }

    return true
}

func getNextNumber(numbers []int) int {
    differences := make([]int, 0, 5)
    for i := 0; i < len(numbers) - 1; i++ {
        differences = append(differences, numbers[i + 1] - numbers[i])
    }

    zeros := all(differences, func (el int) bool { return el == 0 })

    nextNumber := 0

    if ! zeros {
        nextNumber = getNextNumber(differences)
    }

    return numbers[len(numbers) - 1] + nextNumber
}

func parseLine(line string) []int {
    numbers := make([]int, 0, 5)

    sections := strings.Fields(line)

    for _, section := range sections {
        numb, _ := strconv.Atoi(section)
        numbers = append(numbers, numb)
    }

    return numbers
}

func main()  {
    file, _ := os.Open("input.txt")

    scanner := bufio.NewScanner(file)

    sum := 0

    for scanner.Scan() {
        line := scanner.Text()

        numbers := parseLine(line)
        sum += getNextNumber(numbers)
    }

    fmt.Println("Result: ", sum)
}
