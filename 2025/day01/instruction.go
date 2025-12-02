package main

import "strconv"

type direction int

const (
	r direction = iota + 1
	l
)

type instruction struct {
	dir    direction
	amount int
}

func parseInstruction(line string) instruction {
	dir := r
	if line[0] == 'L' {
		dir = l
	}

	amount, _ := strconv.Atoi(line[1:])

	return instruction{
		dir:    dir,
		amount: amount,
	}
}
