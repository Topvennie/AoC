package main

type coordir struct {
	coord coordinate
	dir   direction
}

type coordinate struct {
	x int
	y int
}

type direction int

const (
	up direction = iota + 1
	right
	down
	left
	topRight
	bottomRight
	bottomLeft
	topLeft
)

var directions = []direction{up, right, down, left, topRight, bottomRight, bottomLeft, topLeft}

var dirToCoord = map[direction]coordinate{
	up:          {x: 0, y: -1},
	right:       {x: -1, y: 0},
	down:        {x: 0, y: 1},
	left:        {x: 1, y: 0},
	topRight:    {x: 1, y: -1},
	bottomRight: {x: 1, y: 1},
	bottomLeft:  {x: -1, y: 1},
	topLeft:     {x: -1, y: -1},
}

func (c *coordinate) add(c2 coordinate) *coordinate {
	return &coordinate{x: c.x + c2.x, y: c.y + c2.y}
}
