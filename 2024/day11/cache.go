package main

// Key is a stone
// Value is a map with key the amount of blinks and value the amount of resulting stones
type cache map[stone]map[int]int

var c = make(cache)

func (c *cache) find(s stone, iter int) (int, bool) {
	entry, ok := (*c)[s]
	if !ok {
		return -1, false
	}
	iterEntry, ok := entry[iter]
	if !ok {
		return -1, false
	}

	return iterEntry, true
}

func (c *cache) add(s stone, iter, count int) {
	entry, ok := (*c)[s]
	if !ok {
		(*c)[s] = map[int]int{iter: count}
		return
	}

	_, ok = entry[iter]
	if !ok {
		entry[iter] = count
		(*c)[s] = entry
	}
}
