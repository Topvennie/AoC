package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var part int
	var input string
	flag.IntVar(&part, "p", 1, "Part 1 or 2")
	flag.StringVar(&input, "i", "input.txt", "Location of the input file")
	flag.Parse()

	file, _ := os.ReadFile(input)
	data := strings.TrimSpace(string(file))

	fmt.Printf("Running part: %d with input file: %s\n", part, input)

	if part == 1 {
		solve1(data)
		return
	}

	solve2(data)
}

func solve1(input string) {
	coords := make([]*coordinate, 0)              // All coordinates
	distances := make(map[float64][2]*coordinate) // Key -> distance, Value -> both coordinates

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		coord := &coordinate{x: x, y: y, z: z}

		for _, c := range coords {
			distances[c.distance(coord)] = [2]*coordinate{c, coord}
		}

		coords = append(coords, coord)
	}

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	circuits := make([]*circuit, 0)

	for i := range 1000 {
		coords := distances[keys[i]]

		// Both coordinates are not part of a circuit yet
		if coords[0].circuit == nil && coords[1].circuit == nil {
			newCoords := coords[:]
			circuits = append(circuits, (*circuit)(&newCoords))
			coords[0].circuit = circuits[len(circuits)-1]
			coords[1].circuit = circuits[len(circuits)-1]
			continue
		}

		// Both coordinates are part of a circuit
		if coords[0].circuit != nil && coords[1].circuit != nil {
			// They are already part of the same circuit
			if coords[0].circuit == coords[1].circuit {
				continue
			}

			// Transfer all items from the shortest circuit to the longest
			source := coords[0].circuit
			dest := coords[1].circuit
			if len(*coords[1].circuit) < len(*coords[0].circuit) {
				source = coords[1].circuit
				dest = coords[0].circuit
			}

			for _, coord := range *source {
				coord.circuit = dest
				*dest = append(*dest, coord)
			}

			*source = nil
			continue
		}

		// One of the two is part of a circuit
		if coords[0].circuit != nil {
			*coords[0].circuit = append(*coords[0].circuit, coords[1])
			coords[1].circuit = coords[0].circuit
			continue
		}

		*coords[1].circuit = append(*coords[1].circuit, coords[0])
		coords[0].circuit = coords[1].circuit
	}

	slices.SortFunc(circuits, func(a, b *circuit) int { return len(*a) - len(*b) })

	result := len(*circuits[len(circuits)-1]) * len(*circuits[len(circuits)-2]) * len(*circuits[len(circuits)-3])

	fmt.Println(result)
}

func solve2(input string) {
	coords := make([]*coordinate, 0)              // All coordinates
	distances := make(map[float64][2]*coordinate) // Key -> distance, Value -> both coordinates

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		coord := &coordinate{x: x, y: y, z: z}

		for _, c := range coords {
			distances[c.distance(coord)] = [2]*coordinate{c, coord}
		}

		coords = append(coords, coord)
	}

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	circuits := make([]*circuit, 0)
	keyIdx := 0

	for idx, k := range keys {
		// Check if all junctions are connected
		if len(circuits) == 1 {
			// It's possible that we'done if all coords have circuit pointers
			done := true
			for _, c := range coords {
				if c.circuit == nil {
					done = false
					break
				}
			}

			if done {
				keyIdx = idx - 1
				break
			}
		}

		coords := distances[k]

		// Both coordinates are not part of a circuit yet
		if coords[0].circuit == nil && coords[1].circuit == nil {
			newCoords := coords[:]
			circuits = append(circuits, (*circuit)(&newCoords))
			coords[0].circuit = circuits[len(circuits)-1]
			coords[1].circuit = circuits[len(circuits)-1]
			continue
		}

		// Both coordinates are part of a circuit
		if coords[0].circuit != nil && coords[1].circuit != nil {
			// They are already part of the same circuit
			if coords[0].circuit == coords[1].circuit {
				continue
			}

			// Transfer all items from the shortest circuit to the longest
			source := coords[0].circuit
			dest := coords[1].circuit
			if len(*coords[1].circuit) < len(*coords[0].circuit) {
				source = coords[1].circuit
				dest = coords[0].circuit
			}

			for _, coord := range *source {
				coord.circuit = dest
				*dest = append(*dest, coord)
			}

			idx := slices.IndexFunc(circuits, func(c *circuit) bool { return c == source })
			circuits = append(circuits[:idx], circuits[idx+1:]...)
			continue
		}

		// One of the two is part of a circuit
		if coords[0].circuit != nil {
			*coords[0].circuit = append(*coords[0].circuit, coords[1])
			coords[1].circuit = coords[0].circuit
			continue
		}

		*coords[1].circuit = append(*coords[1].circuit, coords[0])
		coords[0].circuit = coords[1].circuit
	}

	last := distances[keys[keyIdx]]
	fmt.Println(last[0].x * last[1].x)
}
