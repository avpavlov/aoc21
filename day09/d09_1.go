package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d09_1(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var levels = [][]int64{}
	for _, line := range lines {
		var parts = utils.ParseInt64s(strings.Split(line, ""))
		levels = append(levels, parts)
	}

	var result int64 = 0

	var lastR = len(levels) - 1
	for r, row := range levels {
		var lastC = len(row) - 1
		for c, value := range row {
			var up = r == 0 || levels[r-1][c] > value
			var down = r == lastR || levels[r+1][c] > value
			var left = c == 0 || row[c-1] > value
			var right = c == lastC || row[c+1] > value
			if up && down && left && right {
				result += value + 1
			}
		}
	}
	return result
}
