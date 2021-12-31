package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d05_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var result int64 = 0
	var counters = map[utils.Point]int{}
	for _, line := range lines {
		var points = strings.Split(line, " -> ")
		var p1 = utils.ToPoint(points[0])
		var p2 = utils.ToPoint(points[1])
		minY, maxY := utils.MinMax(p1.Y, p2.Y)
		minX, maxX := utils.MinMax(p1.X, p2.X)
		minDelta, maxDelta := utils.MinMax(maxX-minX, maxY-minY)
		if minDelta == maxDelta || minDelta == 0 {
			var n = maxDelta + 1
			var xDir, yDir = utils.Sign(p2.X - p1.X), utils.Sign(p2.Y - p1.Y)
			for i := int64(0); i < n; i++ {
				p := utils.Point{p1.X + xDir*i, p1.Y + yDir*i}
				counters[p]++
				if counters[p] == 2 {
					result++
				}
			}

		}
	}
	return result
}
