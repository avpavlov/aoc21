package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d17_2(filename string) int64 {
	var lines = utils.ReadLines(filename)
	_, coords := utils.SplitPair(lines[0], ": ")
	targetXRangeStr, targetYRangeStr := utils.SplitPair(coords, ", ")
	targetXRange := utils.ParseInt64s(strings.Split(targetXRangeStr[2:], ".."))
	targetXMin := targetXRange[0]
	targetXMax := targetXRange[1]
	targetYRange := utils.ParseInt64s(strings.Split(targetYRangeStr[2:], ".."))
	targetYMin := targetYRange[0]
	targetYMax := targetYRange[1]

	result := map[utils.Point]bool{}

	for xi := int64(1); xi <= targetXMax; xi++ {
	loop_yi:
		for yi := targetYMin; yi <= 1_000; yi++ {
			var x = xi
			var y = yi
			var px = int64(0)
			var py = int64(0)
			for px <= targetXMax && py >= targetYMin {
				px += x
				py += y
				if px >= targetXMin && px <= targetXMax && py >= targetYMin && py <= targetYMax {
					result[utils.Point{xi, yi}] = true
					continue loop_yi
				}
				if x > 0 {
					x--
				}
				y--
			}
		}
	}

	return int64(len(result))
}
