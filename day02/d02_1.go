package main

import . "aoc2021/utils"

func solve_d02_1(fileName string) int64 {
	var lines = ReadLines(fileName)
	var distance = 0
	var depth = 0
	for _, line := range lines {
		d, qs := SplitPair(line, " ")
		var x = int(ParseInt64(qs))
		if d == "forward" {
			distance += x
		} else if d == "down" {
			depth += x
		} else if d == "up" {
			depth -= x
		} else {
			panic(d)
		}
	}
	return int64(distance * depth)
}
