package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumber(str string)  float64 {
    sections := strings.Split(str, ":")
    number := strings.ReplaceAll(sections[1], " ", "")

    num, _ := strconv.Atoi(number)

    return float64(num)
}

func main()  {
    file, _ := os.ReadFile("input.txt")

    input := string(file)

    lines := strings.Split(input, "\n")

    time := getNumber(lines[0])
    distance := getNumber(lines[1])

    discriminant := math.Pow(time, 2.0) - 4 * distance

    upperBound := math.Ceil((time + math.Sqrt(discriminant)) / 2.0 - 1)
    lowerBound := math.Floor((time - math.Sqrt(discriminant)) / 2.0 + 1)

    fmt.Println("Result: ", int(upperBound - lowerBound + 1))
}