package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d08_1(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var result int64 = 0
	for _, line := range lines {
		var parts = strings.Split(line, " | ")
		//var seq = strings.Split(strings.TrimSpace(parts[0]), " ")
		var display = strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, d := range display {
			c := len(d)
			if c == 2 || c == 3 || c == 4 || c == 7 {
				result++
			}
		}
	}
	return result
}
