package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sequence []int

type rule struct {
	first  int
	second int
}

type ruleSet struct {
	set map[int][]rule // The key is the second number in the rule
}

func newRuleSet(rules []rule) *ruleSet {
	set := make(map[int][]rule)

	for _, r := range rules {
		entry, ok := set[r.second]
		if ok {
			entry = append(entry, r)
			set[r.second] = entry
			continue
		}

		set[r.second] = []rule{r}
	}

	return &ruleSet{set: set}
}

func (r *ruleSet) check(seq sequence) bool {
	for i, number := range seq {
		rules, ok := r.set[number]
		if !ok {
			// No rules found, continue
			continue
		}

		// The second number is present, check for first number
		for _, rule := range rules {
			found := false
			for j := i; j < len(seq); j++ {
				if seq[j] == rule.first {
					// First is found, illegal!
					found = true
					break
				}
			}

			if found {
				return false
			}
		}
	}

	return true
}

func parse() ([]rule, []sequence) {
	data, _ := os.ReadFile("input.txt")
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

func main() {
	rules, sequences := parse()
	ruleSet := newRuleSet(rules)

	amount := 0
	for _, seq := range sequences {
		if ruleSet.check(seq) {
			amount += seq[len(seq)/2]
		}
	}

	fmt.Println(amount)
}
