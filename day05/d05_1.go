package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d05_1(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var result int64 = 0
	var counters = map[utils.Point]int{}
	for _, line := range lines {
		var points = strings.Split(line, " -> ")
		var p1 = utils.ToPoint(points[0])
		var p2 = utils.ToPoint(points[1])
		if p1.X == p2.X {
			min, max := utils.MinMax(p1.Y, p2.Y)
			for i := min; i <= max; i++ {
				p := utils.Point{p1.X, i}
				counters[p]++
				if counters[p] == 2 {
					result++
				}
			}
		} else if p1.Y == p2.Y {
			min, max := utils.MinMax(p1.X, p2.X)
			for i := min; i <= max; i++ {
				p := utils.Point{i, p1.Y}
				counters[p]++
				if counters[p] == 2 {
					result++
				}
			}

		}
	}
	return result
}
