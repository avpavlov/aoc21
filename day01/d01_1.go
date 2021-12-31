package main

import . "aoc2021/utils"

func solve_d01_1(fileName string) int64 {
	var lines = ReadLines(fileName)
	var prev int64 = -1
	var count = 0
	for _, line := range lines {
		var m = ParseInt64(line)
		if prev != -1 && prev < m {
			count++
		}
		prev = m
	}
	return int64(count)
}
