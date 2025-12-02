package main

import (
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func parseRange(line string) idRange {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		panic("invalid range")
	}

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	if start > end {
		tmp := start
		start = end
		end = tmp
	}

	return idRange{
		start: start,
		end:   end,
	}
}

func (r *idRange) check(simple bool) []int {
	correct := []int{}

	for i := r.start; i <= r.end; i++ {
		checked := false
		if simple {
			checked = checkNumberSimple(i)
		} else {
			checked = checkNumberAdvanced(i)
		}

		if checked {
			correct = append(correct, i)
		}
	}

	return correct
}

func checkNumberSimple(number int) bool {
	str := strconv.Itoa(number)
	if len(str)%2 != 0 {
		return false
	}

	middle := len(str) / 2

	return str[:middle] == str[middle:]
}

func checkNumberAdvanced(number int) bool {
	str := strconv.Itoa(number)

	for i := range len(str) / 2 {
		if checkStrLen(str, i+1) {
			return true
		}
	}

	return false
}

func checkStrLen(str string, length int) bool {
	if len(str)%length != 0 {
		return false
	}

	part := str[:length]
	for i := length; i+length < len(str)+1; i += length {
		if part != str[i:i+length] {
			return false
		}
	}

	return true
}
