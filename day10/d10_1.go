package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d10_1(filename string) int64 {
	var OPEN = "([{<"
	var CLOSE = ")]}>"
	var COST = map[string]int64{")": 3, "]": 57, "}": 1197, ">": 25137}

	var lines = utils.ReadLines(filename)
	var result int64 = 0
	for _, line := range lines {
		var tokens = strings.Split(line, "")
		var stack = []uint8{}
		for _, token := range tokens {
			if strings.Contains(OPEN, token) {
				stack = append(stack, token[0])
				continue
			}
			var i = strings.Index(CLOSE, token)
			if i >= 0 {
				last := len(stack) - 1
				if last >= 0 && stack[last] == OPEN[i] {
					stack = stack[:last]
				} else {
					result += COST[token]
					break
				}
			} else {
				panic("Cannot be here")
			}

		}
	}
	return result
}
