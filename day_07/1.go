package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HandBid struct {
	hand         string
	handStrength int
	bid          int
}

var cardStrength map[rune]int = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func (handBid *HandBid) setStrength() {
	count := make(map[rune]int)

	for _, c := range handBid.hand {
		count[c]++
	}

	max := 0
	hasTwo := false
	hasThree := false

	for _, amount := range count {
		if amount == 4 || amount == 5 {
			// Five / Four of a kind
			max = amount + 2
			break
		} else if amount == 3 {
			// Full house or Three of a kind
			if hasTwo {
				// Full house
				max = 4
				break
			} else {
				// Three of a kind
				hasThree = true
				max = 3
			}
		} else if amount == 2 {
			// Full house or Two pair or One pair
			if hasThree {
				// Full house
				max = 4
				break
			} else if hasTwo {
				// Two pair
				max = 2
				break
			} else {
				// One pair
				hasTwo = true
                if max == 0 {
                    max = 1
                }
			}
		}
	}

	handBid.handStrength = max
}

func processLine(line string) (handBid HandBid) {
	words := strings.Fields(line)
	bid, _ := strconv.Atoi(words[1])

	handBid = HandBid{
		words[0],
		0,
		bid,
	}

	handBid.setStrength()

	return
}

func (handBid *HandBid) isWeaker(other HandBid) bool {
	if handBid.handStrength != other.handStrength {
		return handBid.handStrength < other.handStrength
	}

	for i, c := range handBid.hand {
		if c != rune(other.hand[i]) {
			return cardStrength[c] < cardStrength[rune(other.hand[i])]
		}
	}

	return false
}

func (handBid HandBid) insertToList(handBids []HandBid) (newList []HandBid) {
	i := 0

	for i < len(handBids) && handBid.isWeaker(handBids[i]) {
		i++
	}

	newList = make([]HandBid, i)
	copy(newList, handBids[:i])

	newList = append(newList, handBid)

	if len(handBids) > i {
		newList = append(newList, handBids[i:]...)
	}

	return
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	handBids := make([]HandBid, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()

		handBids = processLine(line).insertToList(handBids[0:])
    }

    sum := 0

    for i, handBid := range handBids {
        sum += (len(handBids) - i) * handBid.bid
    }

	fmt.Println("Result: ", sum)
}
