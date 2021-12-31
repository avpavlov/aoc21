package main

import . "aoc2021/utils"

func solve_d01_2(fileName string) int64 {
	var lines = ReadLines(fileName)
	var count = 0
	var sum int64 = 0
	for i, line := range lines {
		var m = ParseInt64(line)
		if i < 3 {
			sum += m
			continue
		}
		var newSum = sum + m - ParseInt64(lines[i-3])
		if newSum > sum {
			count++
		}
		sum = newSum
	}
	return int64(count)
}
