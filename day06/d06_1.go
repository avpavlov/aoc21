package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d06_1(filename string) int64 {
	return solve_d06(filename, 80)
}
func solve_d06_2(filename string) int64 {
	return solve_d06(filename, 256)
}

func solve_d06(filename string, days int) int64 {
	var lines = utils.ReadLines(filename)

	var lanterns = [MAX_AGE + 1]int64{}
	var ages = utils.ParseInt64s(strings.Split(lines[0], ","))
	for _, age := range ages {
		lanterns[age]++
	}

	for d := 0; d < days; d++ {
		var newLanterns = [MAX_AGE + 1]int64{}

		for age, lantern := range lanterns {
			if age == 0 {
				newLanterns[6] += lantern
				newLanterns[8] += lantern
			} else {
				newLanterns[age-1] += lantern
			}
		}

		lanterns = newLanterns
	}

	var result int64 = 0
	for _, lantern := range lanterns {
		result += lantern
	}
	return result
}

const MAX_AGE = 8
