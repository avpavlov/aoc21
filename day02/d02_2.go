package main

import "aoc2021/utils"

func solve_d02_2(fileName string) int64 {
	var lines = utils.ReadLines(fileName)
	var distance = 0
	var depth = 0
	var aim = 0
	for _, line := range lines {
		d, qs := utils.SplitPair(line, " ")
		var x = int(utils.ParseInt64(qs))
		if d == "forward" {
			distance += x
			depth += aim * x
		} else if d == "down" {
			aim += x
		} else if d == "up" {
			aim -= x
		} else {
			panic(d)
		}
	}
	return int64(distance * depth)
}
