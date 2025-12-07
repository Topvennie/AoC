package main

import "slices"

// key == Y, value == X
type position map[int][]int

type tachyon struct {
	splitters position
	beams     position

	level  int
	splits int
	width  int
}

func newTachyon(width int) *tachyon {
	return &tachyon{
		splitters: make(position),
		beams:     make(position),
		level:     0,
		splits:    0,
		width:     width,
	}
}

func (t *tachyon) addbeam(x, y int) {
	beams, ok := t.beams[y]
	if !ok {
		beams = make([]int, 0)
	}

	beams = append(beams, x)

	t.beams[y] = beams
}

func (t *tachyon) addSplitter(x, y int) {
	splitters, ok := t.splitters[y]
	if !ok {
		splitters = make([]int, 0)
	}

	splitters = append(splitters, x)

	t.splitters[y] = splitters
}

func (t *tachyon) extend() {
	splitters := t.splitters[t.level+1]

	oldBeams := t.beams[t.level]
	newBeams := make([]int, 0, len(t.beams[t.level]))

	for _, beam := range oldBeams {
		if !slices.Contains(splitters, beam) {
			if t.valid(beam, newBeams) {
				newBeams = append(newBeams, beam)
			}
			continue
		}

		// We hit a splitter
		t.splits++
		if t.valid(beam-1, newBeams) {
			newBeams = append(newBeams, beam-1)
		}
		if t.valid(beam+1, newBeams) {
			newBeams = append(newBeams, beam+1)
		}
	}

	t.beams[t.level+1] = newBeams
	t.level++
}

func (t *tachyon) valid(x int, currentBeams []int) bool {
	if x < 0 {
		return false
	}
	if x >= t.width {
		return false
	}

	return !slices.Contains(currentBeams, x)
}

type positionCount map[int]int

type positionQuantum map[int]positionCount

type tachyonQuantum struct {
	splitters position
	beams     positionQuantum

	level int
	width int
}

func newTachyonQuantum(width int) *tachyonQuantum {
	return &tachyonQuantum{
		splitters: make(position),
		beams:     make(positionQuantum),
		level:     0,
		width:     width,
	}
}

func (t *tachyonQuantum) addbeam(x, y int) {
	beams, ok := t.beams[y]
	if !ok {
		beams = make(positionCount)
	}

	beams[x] = 1

	t.beams[y] = beams
}

func (t *tachyonQuantum) addSplitter(x, y int) {
	splitters, ok := t.splitters[y]
	if !ok {
		splitters = make([]int, 0)
	}

	splitters = append(splitters, x)

	t.splitters[y] = splitters
}

func (t *tachyonQuantum) extend() {
	splitters := t.splitters[t.level+1]

	oldBeams := t.beams[t.level]
	newBeams := make(positionCount)

	for x, count := range oldBeams {
		if !slices.Contains(splitters, x) {
			if t.valid(x) {

				newBeam, ok := newBeams[x]
				if !ok {
					newBeam = 0
				}
				newBeam += count
				newBeams[x] = newBeam
			}
			continue
		}

		// We hit a splitter
		if t.valid(x - 1) {
			newBeam, ok := newBeams[x-1]
			if !ok {
				newBeam = 0
			}
			newBeam += count
			newBeams[x-1] = newBeam
		}
		if t.valid(x + 1) {
			newBeam, ok := newBeams[x+1]
			if !ok {
				newBeam = 0
			}
			newBeam += count
			newBeams[x+1] = newBeam
		}
	}

	t.beams[t.level+1] = newBeams
	t.level++
}

func (t *tachyonQuantum) valid(x int) bool {
	if x < 0 {
		return false
	}
	if x >= t.width {
		return false
	}

	return true
}
