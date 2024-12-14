package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Topvennie/AoC/pkg/coordinate"
)

func main() {
	var part int
	var input string
	flag.IntVar(&part, "p", 1, "Part 1 or 2")
	flag.StringVar(&input, "i", "input.txt", "Location of the input file")
	flag.Parse()

	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(input)
		return
	}

	solve2(input)
}

func solve1(input string) {
	regions := parse(input)

	cost := 0
	for _, region := range regions {
		cost += region.area() * region.perimeter()
	}

	fmt.Println(cost)
}

func solve2(input string) {
	regions := parse(input)

	cost := 0
	for _, region := range regions {
		cost += region.area() * region.sides()
	}

	fmt.Println(cost)
}

func parse(inputFile string) []region {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		panic("File not found")
	}
	data := strings.TrimSpace(string(file))
	lines := strings.Split(data, "\n")

	regions := make(map[rune][]region)

	for y, line := range lines {
		for x, plant := range line {
			entry, ok := regions[plant]
			if ok {
				// Check if plant is already part of another region
				parsed := false
				for _, r := range entry {
					if r.contains(*coordinate.New(x, y)) {
						parsed = true
						break
					}
				}

				if parsed {
					// Already parsed, onto the next plant!
					continue
				}
			}

			// Create r
			r := region{plant: plant, locations: []coordinate.Coord{*coordinate.New(x, y)}}

			// Get region
			previousLength := 0
			for previousLength != len(r.locations) {
				previousLength = len(r.locations)

				for _, loc := range r.locations {
					for _, dir := range coordinate.Dirs {
						newLoc := loc.Add(coordinate.DirToCoord[dir])
						// Check if inside data
						if newLoc.Y >= 0 && newLoc.Y < len(lines) && newLoc.X >= 0 && newLoc.X < len(lines[0]) {
							// Check if not in region yet and the right rune
							if rune(lines[newLoc.Y][newLoc.X]) == r.plant && !r.contains(*newLoc) {
								r.locations = append(r.locations, *newLoc)
							}
						}
					}
				}
			}

			if !ok {
				regions[plant] = []region{r}
				continue
			}

			entry = append(entry, r)
			regions[plant] = entry
		}
	}

	var allRegions []region
	for _, v := range regions {
		allRegions = append(allRegions, v...)
	}

	return allRegions
}
