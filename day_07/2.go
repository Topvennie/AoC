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
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

// An abomination
func (handBid *HandBid) setStrength() {
	count := make(map[rune]int)

	for _, c := range handBid.hand {
		count[c]++
	}

	max := 0
	setMax := func(x int) {
		if x > max {
			max = x
		}
	}
	hasTwo := false
	hasThree := false

	maxJ := 0
	setMaxJ := func(x int) {
		if x > maxJ {
			maxJ = x
		}
	}
	hasTwoJ := false
	hasThreeJ := false

	jokers := count['J']

	for c, amount := range count {
		twoThreeNow := false

		if amount == 4 || amount == 5 {
			// Five / Four of a kind
			setMax(amount + 2)
		} else if amount == 3 {
			// Full house or Three of a kind
			if hasTwo || (hasTwoJ && c != 'J') {
				// Full house
				setMax(4)
			} else {
				// Three of a kind
				if !hasThree {
					twoThreeNow = true
				}
                if c == 'J' {
                    hasThreeJ = true
                } else {
                    hasThree = true
                }
				setMax(3)
			}
		} else if amount == 2 {
			// Full house or Two pair or One pair
			if hasThree || (hasThreeJ && c != 'J') {
				// Full house
				setMax(4)
			} else if hasTwo || (hasTwoJ && c != 'J') {
				// Two pair
				setMax(2)
			} else {
				// One pair
				if !hasTwo {
					twoThreeNow = true
				}
                if c == 'J' {
                    hasTwoJ = true
                } else {
                    hasTwo = true
                }
				setMax(1)
			}
		}

		if c != 'J' && jokers > 0 {
			amount += jokers

			if amount == 4 || amount == 5 {
				// Five / Four of a kind
				setMaxJ(amount + 2)
			} else if amount == 3 {
				// Full house or Three of a kind
				if hasTwo && !twoThreeNow {
					// Full house
					setMaxJ(4)
				} else {
					// Three of a kind
					hasThreeJ = true
					setMaxJ(3)
				}
			} else if amount == 2 {
				// Full house or Two pair or One pair
				if hasThree && ! twoThreeNow {
					// Full house
					setMaxJ(4)
				} else if hasTwo && !twoThreeNow {
					// Two pair
					setMaxJ(2)
				} else {
					// One pair
					hasTwoJ = true
					setMaxJ(1)
				}
			}
		}

	}

	if max > maxJ {
		handBid.handStrength = max
	} else {
		handBid.handStrength = maxJ
	}
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
