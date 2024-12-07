package main

import (
	"fmt"
	"strconv"
)

type operator interface {
	exec() int
}

type add struct {
	x int
	y int
}

func newAdd(x, y int) operator {
	return &add{x: x, y: y}
}

func (a *add) exec() int {
	return a.x + a.y
}

type multiply struct {
	x int
	y int
}

func newMultiply(x, y int) operator {
	return &multiply{x: x, y: y}
}

func (m *multiply) exec() int {
	return m.x * m.y
}

type concat struct {
	x int
	y int
}

func newConcat(x, y int) operator {
	return &concat{x: x, y: y}
}

func (c *concat) exec() int {
	strNumber := fmt.Sprintf("%d%d", c.x, c.y)
	number, _ := strconv.Atoi(strNumber)

	return number
}
