package main

import (
	"github.com/Topvennie/AoC/pkg/coordinate"
)

var (
	costA = 3
	costB = 1
)

type machine struct {
	a     coordinate.Coord
	b     coordinate.Coord
	price coordinate.Coord
}

func (m *machine) cost() int {
	amountA, amountB, solved := m.solve()

	// Check answer
	if !solved {
		return 0
	}

	return amountA*costA + amountB*costB
}

func (m *machine) solve() (int, int, bool) {
	b := (m.a.X*m.price.Y - m.a.Y*m.price.X) / (m.a.X*m.b.Y - m.a.Y*m.b.X)
	a := (m.price.X - b*m.b.X) / m.a.X

	if a*m.a.X+b*m.b.X != m.price.X || a*m.a.Y+b*m.b.Y != m.price.Y {
		return -1, -1, false
	}

	return a, b, true
}

func sliceToFloat64(slice []int) []float64 {
	result := make([]float64, 0, len(slice))
	for _, number := range slice {
		result = append(result, float64(number))
	}

	return result
}
