package main

import "aoc2021/utils"

func solve_d03_1(fileName string) int64 {
	var lines = utils.ReadLines(fileName)
	var counters = make([]int, len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				counters[i]++
			}
		}
	}
	var mid = len(lines) / 2
	var gamma = 0
	var epsilon = 0
	for _, counter := range counters {
		gamma *= 2
		epsilon *= 2
		if counter > mid {
			gamma++
		} else {
			epsilon++
		}
	}
	return int64(gamma * epsilon)
}
