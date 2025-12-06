package main

type problem struct {
	numbers []int
}

func newProblem() *problem {
	return &problem{
		numbers: []int{},
	}
}

func (p *problem) add(input int) {
	p.numbers = append(p.numbers, input)
}

type operator string

const (
	add      operator = "+"
	multiply operator = "*"
)

func (p *problem) solve(op operator) int {
	switch op {
	case add:
		return reduce(p.numbers, func(curr int, n int) int { return curr + n }, 0)
	case multiply:
		return reduce(p.numbers, func(curr int, n int) int { return curr * n }, 1)
	}
	return 0
}

func reduce[T any, U any](input []T, fn func(U, T) U, acc U) U {
	for i := range input {
		acc = fn(acc, input[i])
	}

	return acc
}
