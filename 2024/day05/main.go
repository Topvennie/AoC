package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	rules, sequences := parse(input)
	ruleSet := newRuleSet(rules)

	amount := 0
	for _, seq := range sequences {
		if ruleSet.check(seq) {
			amount += seq[len(seq)/2]
		}
	}

	fmt.Println(amount)
}

func solve2(input string) {
	rules, sequences := parse(input)
	ruleSet := newRuleSet(rules)

	amount := 0
	for _, seq := range sequences {
		if !ruleSet.check(seq) {
			s := ruleSet.format(seq)
			amount += s[len(s)/2]
		}
	}

	fmt.Println(amount)
}

func parse(input string) ([]rule, []sequence) {
	data, _ := os.ReadFile(input)
	parts := strings.Split(string(data), "\n\n")

	var rules []rule
	rulePart := strings.Split(parts[0], "\n")
	for _, r := range rulePart {
		parts := strings.Split(r, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		rules = append(rules, rule{first: first, second: second})
	}

	var sequences []sequence
	sequencePart := strings.Split(parts[1], "\n")
	for _, s := range sequencePart {
		var seq sequence
		parts := strings.Split(s, ",")
		for _, part := range parts {
			number, _ := strconv.Atoi(part)
			seq = append(seq, number)
		}
		sequences = append(sequences, seq)
	}

	return rules, sequences
}
