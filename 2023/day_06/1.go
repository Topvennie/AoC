package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumbers(str string, numberArray *[]float64) {
    sections := strings.Split(str, ":")
    numbers := strings.Fields(sections[1])

    for _, number := range numbers {
        num, _ := strconv.Atoi(number)
        *numberArray = append(*numberArray, float64(num))
    }
}

func main()  {
    file, _ := os.ReadFile("input.txt")

    input := string(file)

    times := make([]float64, 0, 10)
    distances := make([]float64, 0, 10)

    lines := strings.Split(input, "\n")

    getNumbers(lines[0], &times)
    getNumbers(lines[1], &distances)

    sum := 1.0

    for i, time := range times {
        discriminant := math.Pow(time, 2.0) - 4 * distances[i]

        upperBound := math.Ceil((time + math.Sqrt(discriminant)) / 2.0 - 1)
        lowerBound := math.Floor((time - math.Sqrt(discriminant)) / 2.0 + 1)

        fmt.Println(upperBound)
        fmt.Println(lowerBound)

        sum *= upperBound - lowerBound + 1
    }

    fmt.Println("Result: ", sum)
}