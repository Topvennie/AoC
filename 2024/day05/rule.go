package main

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

func (r *ruleSet) format(seq sequence) sequence {
	for i, number := range seq {
		rules, ok := r.set[number]
		if !ok {
			continue
		}

		for _, rule := range rules {
			for j := i; j < len(seq); j++ {
				if seq[j] == rule.first {
					// Found! Swap them
					seq[i] = rule.first
					seq[j] = number
					return r.format(seq)
				}
			}
		}
	}

	return seq
}
