package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

var nodes map[string]Node

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findLCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	nodes = make(map[string]Node)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	starters := make([]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()

		sections := strings.Split(line, " = ")
		destinations := strings.Split(sections[1][1:len(sections[1])-1], ", ")

		nodes[sections[0]] = Node{
			destinations[0],
			destinations[1],
		}

		if string(sections[0][2]) == "A" {
			starters = append(starters, sections[0])
		}
	}

    all_steps := make([]int, 0, 10)

	for _, starter := range starters {
		steps := 0

		for found := false; !found; found = string(starter[2]) == "Z" {
			char := instructions[steps%len(instructions)]

			if char == 82 {
				starter = nodes[starter].right
			} else {
				starter = nodes[starter].left
			}

			steps++
		}

        all_steps = append(all_steps, steps)
	}

    fmt.Println(all_steps)

	fmt.Println("Result: ", findLCM(all_steps))
}
