package main

import (
	"slices"
)

type ingredient struct {
	start int
	end   int
}

type inventory struct {
	fresh []ingredient
}

func newInventory() *inventory {
	return &inventory{
		fresh: []ingredient{},
	}
}

func (i *inventory) add(ingr ingredient) {
	idx := 0
	for _, f := range i.fresh {
		if f.start > ingr.start {
			break
		}

		idx++
	}

	i.fresh = slices.Insert(i.fresh, idx, ingr)
}

func (i *inventory) normalize() {
	for idx := range i.fresh {
		if idx == len(i.fresh)-1 {
			return
		}

		if i.fresh[idx].end >= i.fresh[idx+1].start {
			if i.fresh[idx].end <= i.fresh[idx+1].end {
				i.fresh[idx].end = i.fresh[idx+1].end
			}

			if idx <= len(i.fresh)-2 {
				i.fresh = append(i.fresh[:idx+1], i.fresh[idx+2:]...)
			}

			i.normalize()

			return
		}
	}
}

func (i *inventory) isFresh(id int) bool {
	for _, f := range i.fresh {
		if f.start > id {
			break
		}

		if f.end >= id {
			return true
		}
	}

	return false
}

func (i *inventory) total() int {
	total := 0

	for _, f := range i.fresh {
		total += f.end - f.start + 1
	}

	return total
}
