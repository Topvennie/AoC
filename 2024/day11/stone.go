package main

import (
	"math"
)

type stone int

func (s stone) evolveCount(iter, maxIter int) int {
	amount, ok := c.find(s, maxIter-iter)
	if ok {
		return amount
	}

	amount = 0

	evolved := s.evolve()
	if iter == maxIter {
		amount += len(evolved)
	} else {
		for _, evolution := range evolved {
			amount += evolution.evolveCount(iter+1, maxIter)
		}
	}

	c.add(s, maxIter-iter, amount)

	return amount
}

func (s stone) evolve() []stone {
	if s == 0 {
		return []stone{1}
	}

	length := s.length()
	if length%2 == 0 {
		return s.split(length)
	}

	return []stone{stone(s * 2024)}
}

func (s stone) length() int {
	if s == 0 {
		return 1
	}

	count := 0
	for s != 0 {
		s /= 10
		count++
	}

	return count
}

func (s stone) split(length int) []stone {
	divider := int(math.Pow(float64(10), float64(length/2)))

	first := int(s) / divider
	second := (int(s) % divider)

	return []stone{stone(first), stone(second)}
}
