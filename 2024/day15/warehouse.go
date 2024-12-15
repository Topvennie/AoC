package main

import (
	"strings"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

type warehouse struct {
	robot  coordinate.Coord
	crates map[coordinate.Coord]bool
	walls  map[coordinate.Coord]bool

	movements []coordinate.Dir

	w int
	h int
}

func new() *warehouse {
	return &warehouse{
		crates:    make(map[coordinate.Coord]bool),
		walls:     make(map[coordinate.Coord]bool),
		movements: make([]coordinate.Dir, 0),
	}
}

func (w *warehouse) move() {
	for _, mov := range w.movements {
		if !w.valid(mov) {
			continue
		}

		dest := *w.robot.Add(coordinate.DirToCoord[mov])

		// Adjust robot position
		w.robot = dest

		// Check if there was a box
		box := false
		if entry := w.crates[dest]; entry {
			w.crates[dest] = false
			box = true
		}

		for box {
			dest = *dest.Add(coordinate.DirToCoord[mov])

			// Check for a following box
			box = false
			if entry := w.crates[dest]; entry {
				box = true
			}

			w.crates[dest] = true
		}
	}
}

func (w *warehouse) moveWide() {
	for _, mov := range w.movements {
		if !w.validWide(mov) {
			continue
		}

		dest := *w.robot.Add(coordinate.DirToCoord[mov])

		// Adjust robot position
		w.robot = dest

		// Check if there was a box
		box := false
		if entry := w.crates[dest]; entry {
			w.crates[dest] = false
			box = true
			w.moveCrate(dest, mov)
		}

		if !box {
			left := *dest.Add(coordinate.DirToCoord[coordinate.Left])
			if entry := w.crates[left]; entry {
				w.crates[left] = false
				box = true
				w.moveCrate(left, mov)
			}
		}
	}
}

func (w *warehouse) moveCrate(pos coordinate.Coord, dir coordinate.Dir) {
	dest := *pos.Add(coordinate.DirToCoord[dir])
	if entry := w.crates[dest]; entry {
		w.moveCrate(dest, dir)
	}
	left := *dest.Add(coordinate.DirToCoord[coordinate.Left])
	if entry := w.crates[left]; entry {
		w.crates[left] = false
		w.moveCrate(left, dir)
	}
	right := *dest.Add(coordinate.DirToCoord[coordinate.Right])
	if entry := w.crates[right]; entry {
		w.crates[right] = false
		w.moveCrate(right, dir)
	}

	w.crates[dest] = true
}

func (w *warehouse) score() int {
	total := 0
	for k, v := range w.crates {
		if !v {
			continue
		}

		total += k.X + 1 + (k.Y+1)*100
	}

	return total
}

func (w *warehouse) scoreWide() int {
	total := 0
	for k, v := range w.crates {
		if !v {
			continue
		}

		total += k.X + 2 + (k.Y+1)*100
	}

	return total
}

// valid checks if a move is valid
func (w *warehouse) valid(mov coordinate.Dir) bool {
	dest := *w.robot.Add(coordinate.DirToCoord[mov])
	for w.inside(dest) {
		// Is there a wall?
		if entry := w.walls[dest]; entry {
			return false
		}

		// Is there a crate?
		if entry := w.crates[dest]; entry {
			// There is a crate, go again!
			dest = *dest.Add(coordinate.DirToCoord[mov])
			continue
		}

		// It's clear
		return true
	}

	return false
}

func (w *warehouse) validWide(mov coordinate.Dir) bool {
	crateCheck := map[coordinate.Dir]func(coordinate.Coord, coordinate.Dir) bool{
		coordinate.Up:    w.validWideCrateVertical,
		coordinate.Down:  w.validWideCrateVertical,
		coordinate.Left:  w.validWideCrateHorizontalLeft,
		coordinate.Right: w.validWideCrateHorizontalRight,
	}

	dest := *w.robot.Add(coordinate.DirToCoord[mov])
	if !w.inside(dest) {
		return false
	}

	// Is there a wall?
	if entry := w.walls[dest]; entry {
		return false
	}

	// Is there a crate?
	left := *dest.Add(coordinate.DirToCoord[coordinate.Left])
	if entry := w.crates[left]; entry {
		return crateCheck[mov](left, mov)
	}

	if entry := w.crates[dest]; entry {
		return crateCheck[mov](dest, mov)
	}

	// It's clear
	return true
}

func (w *warehouse) validWideCrateHorizontalRight(pos coordinate.Coord, mov coordinate.Dir) bool {
	dest := *pos.Add(coordinate.DirToCoord[mov])
	destRight := *dest.Add(coordinate.DirToCoord[mov])

	if !w.inside(dest) {
		return false
	}

	if !w.inside(destRight) {
		return false
	}

	// Check for walls
	if entry := w.walls[dest]; entry {
		return false
	}

	if entry := w.walls[destRight]; entry {
		return false
	}

	if entry := w.crates[destRight]; entry {
		return w.validWideCrateHorizontalRight(destRight, mov)
	}

	return true
}

func (w *warehouse) validWideCrateHorizontalLeft(pos coordinate.Coord, mov coordinate.Dir) bool {
	dest := *pos.Add(coordinate.DirToCoord[mov])
	destLeft := *dest.Add(coordinate.DirToCoord[mov])

	if !w.inside(dest) {
		return false
	}

	// Check for walls
	if entry := w.walls[dest]; entry {
		return false
	}

	if entry := w.crates[destLeft]; entry {
		return w.validWideCrateHorizontalLeft(destLeft, mov)
	}

	return true
}

func (w *warehouse) validWideCrateVertical(pos coordinate.Coord, mov coordinate.Dir) bool {
	dest := *pos.Add(coordinate.DirToCoord[mov])
	destRight := *dest.Add(coordinate.DirToCoord[coordinate.Right])
	destLeft := *dest.Add(coordinate.DirToCoord[coordinate.Left])

	if !w.inside(dest) || !w.inside(destRight) {
		return false
	}

	// Check for walls
	if entry := w.walls[dest]; entry {
		return false
	}

	if entry := w.walls[destRight]; entry {
		return false
	}

	valid := true

	// Check for crates
	if entry := w.crates[destLeft]; entry {
		valid = valid && w.validWideCrateVertical(destLeft, mov)
	}

	if entry := w.crates[dest]; entry {
		valid = valid && w.validWideCrateVertical(dest, mov)
	}

	if entry := w.crates[destRight]; entry {
		valid = valid && w.validWideCrateVertical(destRight, mov)
	}

	return valid
}

func (w *warehouse) inside(c coordinate.Coord) bool {
	if c.X < 0 || c.X >= w.w {
		return false
	}

	if c.Y < 0 || c.Y >= w.h {
		return false
	}

	return true
}

func (w *warehouse) print() string {
	lines := make([][]rune, 0, w.h)
	for range w.h {
		line := make([]rune, 0, w.w)
		for range w.w {
			line = append(line, '.')
		}
		lines = append(lines, line)
	}

	// Place robot
	lines[w.robot.Y][w.robot.X] = '@'

	// Place all boxes
	for k, v := range w.crates {
		if !v {
			continue
		}

		lines[k.Y][k.X] = 'O'
	}

	// Place all walls
	for k, v := range w.walls {
		if !v {
			continue
		}

		lines[k.Y][k.X] = '#'
	}

	var overview strings.Builder
	for _, line := range lines {
		overview.WriteString(string(line))
		overview.WriteString("\n")
	}

	return overview.String()
}

func (w *warehouse) printWide() string {
	lines := make([][]rune, 0, w.h)
	for range w.h {
		line := make([]rune, 0, w.w)
		for range w.w {
			line = append(line, '.')
		}
		lines = append(lines, line)
	}

	// Place robot
	lines[w.robot.Y][w.robot.X] = '@'

	// Place all boxes
	for k, v := range w.crates {
		if !v {
			continue
		}

		lines[k.Y][k.X] = '['
		lines[k.Y][k.X+1] = ']'
	}

	// Place all walls
	for k, v := range w.walls {
		if !v {
			continue
		}

		lines[k.Y][k.X] = '#'
	}

	var overview strings.Builder
	for _, line := range lines {
		overview.WriteString(string(line))
		overview.WriteString("\n")
	}

	return overview.String()
}
