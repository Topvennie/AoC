package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
    left string
    right string
}

var nodes map[string]Node

func main()  {
    file, _ := os.Open("input.txt")

    scanner := bufio.NewScanner(file)
    nodes = make(map[string]Node)

    scanner.Scan()
    instructions := scanner.Text()
    scanner.Scan()

    for scanner.Scan() {
        line := scanner.Text()

        sections := strings.Split(line, " = ")
        destinations := strings.Split(sections[1][1:len(sections[1]) - 1], ", ")

        nodes[sections[0]] = Node {
            destinations[0],
            destinations[1],
        }
    }

    steps := 0
    current := "AAA"

    for found := false; ! found; found = current == "ZZZ"  {
        char := instructions[steps % len(instructions)]

        if char == 82 {
            current = nodes[current].right
        } else {
            current = nodes[current].left
        }

        steps++
    }

    fmt.Println("Result: ", steps)
}