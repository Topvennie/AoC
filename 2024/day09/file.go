package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type file []block

const free = -1

type block struct {
	id     int
	amount int
}

func (f *file) order() {
	for i := 0; i < len(*f); i++ {
		// Isn't free so go to the next
		if (*f)[i].id != free {
			continue
		}

		// It's a free space
		// Check if we have something to replace it with
		idx := -1
		for j := len(*f) - 1; j > i; j-- {
			if (*f)[j].id != free {
				idx = j
				break
			}
		}

		// Not found
		// We're done
		if idx == -1 {
			break
		}

		// Insert differently depending if the amount is equal, less or more
		freeAmount := (*f)[i].amount
		if freeAmount < (*f)[idx].amount {
			// Less, Move as much as possible
			freeAmount := (*f)[i].amount
			(*f)[i].id = (*f)[idx].id
			(*f)[idx].amount -= freeAmount
			(*f) = append(*f, block{id: free, amount: (*f)[idx].amount - freeAmount})
		} else if freeAmount > (*f)[idx].amount {
			// More, move everything
			idxBlock := (*f)[idx]
			(*f)[i].amount -= idxBlock.amount
			(*f) = slices.Delete((*f), idx, idx+1)
			(*f) = slices.Insert((*f), i, block{id: idxBlock.id, amount: idxBlock.amount})
		} else {
			// Equal size, swap
			swap := (*f)[idx]
			(*f)[idx] = (*f)[i]
			(*f)[i] = swap
		}
	}

	// Combine any potential free spaces at the back
	for (*f)[len(*f)-1].id == free && (*f)[len(*f)-2].id == free {
		(*f)[len(*f)-2].amount += (*f)[len(*f)-1].amount
		(*f) = (*f)[:len(*f)-1]
	}
}

func (f *file) orderWhole() {
	for i := len(*f) - 1; i >= 0; i-- {
		// It's free , try next
		if (*f)[i].id == free {
			continue
		}

		// Check for free space where it fits
		idx := -1
		for j := 0; j < i; j++ {
			if (*f)[j].id == free && (*f)[j].amount >= (*f)[i].amount {
				idx = j
				break
			}
		}

		if idx == -1 {
			// Not found, onto the next
			continue
		}

		if (*f)[i].amount < (*f)[idx].amount {
			// Free space is bigger
			// Move everything
			(*f)[idx].amount -= (*f)[i].amount
			oldID := (*f)[i].id
			(*f)[i].id = free
			(*f) = slices.Insert(*f, idx, block{id: oldID, amount: (*f)[i].amount})
		} else {
			// Same size, swap
			swap := (*f)[idx]
			(*f)[idx] = (*f)[i]
			(*f)[i] = swap
		}
	}

	// Combine any potential free spaces at the back
	for (*f)[len(*f)-1].id == free && (*f)[len(*f)-2].id == free {
		(*f)[len(*f)-2].amount += (*f)[len(*f)-1].amount
		(*f) = (*f)[:len(*f)-1]
	}
}

func (f *file) checksum() int {
	amount := 0
	i := 0
	for _, b := range *f {
		if b.id == free {
			i += b.amount
			continue
		}

		for range b.amount {
			amount += i * b.id
			i++
		}
	}

	return amount
}

func (f *file) print() {
	for _, block := range *f {
		char := "."
		if block.id != free {
			char = strconv.Itoa(block.id)
		}
		fmt.Printf(strings.Repeat(char, block.amount))
	}

	fmt.Printf("\n")
}
