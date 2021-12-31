package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d11_2(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var levels = [][]int64{}
	for _, line := range lines {
		var parts = utils.ParseInt64s(strings.Split(line, ""))
		levels = append(levels, parts)
	}

	var total = len(levels) * len(levels[0])
	var result int64 = 0
	for step := int64(0); true; step++ {
		toFlash := []utils.Point{}
		for r, row := range levels {
			for c, value := range row {
				value++
				row[c] = value
				if value == 10 {
					toFlash = append(toFlash, utils.Point{int64(r), int64(c)})
				}
			}
		}

		for f := 0; f < len(toFlash); f++ {
			result++
			p := toFlash[f]
			for _, ff := range flash(levels, p.X, p.Y) {
				toFlash = append(toFlash, ff)
			}
		}

		if total == len(toFlash) {
			return step + 1
		}
		for _, p := range toFlash {
			levels[p.X][p.Y] = 0
		}
		toFlash = []utils.Point{}
	}
	return result
}
