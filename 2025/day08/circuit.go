package main

import (
	"math"
)

type coordinate struct {
	x int
	y int
	z int

	circuit *circuit
}

func (c *coordinate) distance(c2 *coordinate) float64 {
	return math.Sqrt(math.Pow((float64(c.x-c2.x)), 2) + math.Pow((float64(c.y-c2.y)), 2) + math.Pow((float64(c.z-c2.z)), 2))
}

func (c coordinate) equal(c2 coordinate) bool {
	return c.x == c2.x && c.y == c2.y && c.z == c2.z
}

type circuit []*coordinate
