package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generate(width, height int, file string) {
	var rows strings.Builder

	for range height {
		var row strings.Builder
		for range width {
			number := rand.Intn(10)
			row.WriteString(strconv.Itoa(number))
		}

		row.WriteString("\n")
		rows.WriteString(row.String())
	}

	os.WriteFile(file, []byte(rows.String()), os.ModePerm)
}
