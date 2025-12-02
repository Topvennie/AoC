package main

type dial struct {
	previous int
	current  int
	zeros    int
}

func (d *dial) add(instr instruction) {
	amount := instr.amount
	if instr.dir == l {
		amount *= -1
	}

	d.previous = d.current
	d.current += amount

	d.normalize()
}

func (d *dial) normalize() {
	for d.current < 0 {
		d.current += 100
	}

	for d.current > 99 {
		d.current -= 100
	}

	if d.current == 0 {
		d.zeros++
	}
}
