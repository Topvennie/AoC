package main

type data struct {
	data [][]rune
}

func newData(lines []string) *data {
	d := make([][]rune, 0, len(lines))

	for _, line := range lines {
		d = append(d, []rune(line))
	}

	return &data{data: d}
}

func (d *data) getXMAS() int {
	// Loop over data
	amount := 0
	for i := range d.data {
		for j := range d.data[i] {
			coord := coordinate{x: i, y: j}
			// Check X
			if !d.checkLetter(coord, 'X') {
				continue
			}

			coordirs := make([]coordir, 0, len(directions))
			for _, dir := range directions {
				coordirs = append(coordirs, coordir{coord: coord, dir: dir})
			}

			// Check M
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'M')
			// Check A
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'A')
			// Check S
			coordirs = d.allNeighbours(coordirs)
			coordirs = d.checkLetters(coordirs, 'S')

			amount += len(coordirs)
		}
	}

	return amount
}

func (d *data) inside(c coordinate) bool {
	if c.x < 0 || c.x >= len(d.data) {
		return false
	}

	if c.y < 0 || c.y >= len(d.data[c.x]) {
		return false
	}

	return true
}

func (d *data) allNeighbours(c []coordir) []coordir {
	var coordirs []coordir
	for _, coordir := range c {
		newCoordir, valid := d.neighbour(coordir)
		if valid {
			coordirs = append(coordirs, newCoordir)
		}
	}

	return coordirs
}

func (d *data) neighbour(c coordir) (coordir, bool) {
	newCoord := c.coord.add(dirToCoord[c.dir])
	if d.inside(*newCoord) {
		return coordir{coord: *newCoord, dir: c.dir}, true
	}

	return c, false
}

func (d *data) checkLetters(coordirs []coordir, letter rune) []coordir {
	var coords []coordir
	for _, cd := range coordirs {
		if d.checkLetter(cd.coord, letter) {
			coords = append(coords, cd)
		}
	}

	return coords
}

func (d *data) checkLetter(c coordinate, letter rune) bool {
	if c.x < 0 || c.x >= len(d.data) {
		return false
	}

	if c.y < 0 || c.y >= len(d.data[c.x]) {
		return false
	}

	return d.data[c.x][c.y] == letter
}

func (d *data) getXMASSlope() int {
	// Get all mas

	amount := 0
	for i := range d.data {
		for j := range d.data[i] {
			coord := coordinate{x: i, y: j}
			// Check A
			if !d.checkLetter(coord, 'A') {
				continue
			}

			if d.checkSlope(coord, topLeft, bottomRight) && d.checkSlope(coord, topRight, bottomLeft) {
				amount++
			}
		}
	}

	return amount
}

func (d *data) checkSlope(c coordinate, dir1 direction, dir2 direction) bool {
	opposite := map[rune]rune{
		'M': 'S',
		'S': 'M',
	}

	pos1 := c.add(dirToCoord[dir1])
	if !d.inside(*pos1) {
		return false
	}

	pos2 := c.add(dirToCoord[dir2])
	if !d.inside(*pos2) {
		return false
	}

	rune1 := d.data[pos1.x][pos1.y]
	rune2, ok := opposite[rune1]
	if !ok {
		return false
	}

	return rune2 == d.data[pos2.x][pos2.y]
}
