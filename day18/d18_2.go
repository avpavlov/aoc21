package main

import "aoc2021/utils"

func solve_d18_2(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var max int64
	for i1, line1 := range lines {
		for i2, line2 := range lines {
			if i1 != i2 {
				copy1 := line1
				num2 := parseSnailfishNum(&copy1).addSnailfishNum(parseSnailfishNum(&line2))
				for num2.reduceSnailfishNum() {
				}
				if m := num2.magnitude(); m > max {
					max = m
				}
			}
		}

	}

	return max
}
