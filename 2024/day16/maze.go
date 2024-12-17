package main

import (
	"math"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

type maze struct {
	walls map[coordinate.Coord]bool
	start coordinate.Coord
	end   coordinate.Coord

	w int
	h int
}

func newMaze() *maze {
	return &maze{walls: make(map[coordinate.Coord]bool)}
}

func (m *maze) solve() (int, []position) {
	// Construct starting postion
	startPos := *newPosition()
	startPos.current = posDir{p: m.start, d: coordinate.Right}

	minCost := math.MaxInt
	posCost := map[posDir]int{} // Best cost for that position
	leaves := []position{startPos}

	solutions := []position{}

	for len(leaves) > 0 {
		newLeaves := []position{}
		for _, leaf := range leaves {
			// Prune
			if leaf.cost > minCost {
				continue
			}

			if entry, ok := posCost[leaf.current]; ok && entry < leaf.cost {
				continue
			}

			for _, pos := range m.possibilities(leaf) {
				// Prune
				if pos.cost > minCost {
					continue
				}

				entry, ok := posCost[pos.current]
				// TODO: inline
				if ok {
					if entry < pos.cost {
						continue
					}
				}

				if pos.done {
					if pos.cost < minCost {
						solutions = []position{pos}
						minCost = pos.cost
					} else if pos.cost == minCost {
						solutions = append(solutions, pos)
					}

					continue
				}

				posCost[pos.current] = pos.cost
				newLeaves = append(newLeaves, pos)
			}
		}

		leaves = newLeaves
	}

	return minCost, solutions
}

func (m *maze) spots(solutions []position) int {
	unique := map[coordinate.Coord]bool{}

	for _, sol := range solutions {
		for _, p := range sol.previous {
			if entry := unique[p.p]; !entry {
				unique[p.p] = true
			}
		}
	}

	return len(unique) + 1 // + 1 for the start position
}

func (m *maze) possibilities(p position) []position {
	positions := []position{}

	// 3 possiblities
	// 1: Move in dir direction
	// 2: Turn 90 left
	// 3: Turn 90 right

	left90 := map[coordinate.Dir]coordinate.Dir{
		coordinate.Up:    coordinate.Left,
		coordinate.Right: coordinate.Up,
		coordinate.Down:  coordinate.Right,
		coordinate.Left:  coordinate.Down,
	}
	right90 := map[coordinate.Dir]coordinate.Dir{
		coordinate.Up:    coordinate.Right,
		coordinate.Right: coordinate.Down,
		coordinate.Down:  coordinate.Left,
		coordinate.Left:  coordinate.Up,
	}

	// 1: Move
	newPos := *p.current.p.AddDir(p.current.d)
	if m.inside(newPos) && !m.walls[newPos] {
		// No wall and we haven't visited it before, let's move
		pCopy := *p.copy()

		pCopy.current.p = newPos
		pCopy.cost++
		pCopy.previous = append(pCopy.previous, p.current)
		pCopy.done = newPos.Equal(m.end)

		positions = append(positions, pCopy)
	}

	// 2: 90 Left
	newDir := left90[p.current.d]
	pCopy := *p.copy()

	pCopy.current.d = newDir
	pCopy.cost += 1000

	positions = append(positions, pCopy)

	// 3: 90 Right
	newDir = right90[p.current.d]
	pCopy = *p.copy()

	pCopy.current.d = newDir
	pCopy.cost += 1000

	positions = append(positions, pCopy)

	return positions
}

func (m *maze) inside(c coordinate.Coord) bool {
	if c.X < 0 || c.X >= m.w {
		return false
	}

	if c.Y < 0 || c.Y >= m.h {
		return false
	}

	return true
}
