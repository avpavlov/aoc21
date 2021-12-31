package main

import (
	"aoc2021/utils"
	"log"
	"strings"
)

var OPEN = "([{<"
var CLOSE = ")]}>"
var COMPLETE_COST = map[uint8]int64{'(': 1, '[': 2, '{': 3, '<': 4}

func solve_d10_2(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var incomplete = []int64{}

outerLoop:
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
					continue outerLoop
				}
			} else {
				panic("Cannot be here")
			}
		}
		if len(stack) == 0 {
			continue
		}
		var fixCost int64 = 0
		for j := len(stack) - 1; j >= 0; j-- {
			c := COMPLETE_COST[stack[j]]
			if c < 1 || c > 4 {
				log.Panicf("Cost %d", c)
			}
			fixCost = fixCost*5 + c
		}
		incomplete = append(incomplete, fixCost)

	}

	utils.SortInt64s(incomplete)
	return incomplete[(len(incomplete)-1)/2]
}
