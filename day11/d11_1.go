package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d11_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var levels = [][]int64{}
	for _, line := range lines {
		var parts = utils.ParseInt64s(strings.Split(line, ""))
		levels = append(levels, parts)
	}

	var result int64 = 0
	for step := 0; step < 100; step++ {
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

		for _, p := range toFlash {
			levels[p.X][p.Y] = 0
		}
		toFlash = []utils.Point{}
	}
	return result
}

func flash(levels [][]int64, r int64, c int64) []utils.Point {
	lastR := int64(len(levels) - 1)
	lastC := int64(len(levels[r]) - 1)

	toFlash := []utils.Point{}
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i == r && j == c {
				continue
			} else if i < 0 || i > lastR {
				continue
			} else if j < 0 || j > lastC {
				continue
			}
			value := levels[i][j]
			value++
			levels[i][j] = value
			if value == 10 {
				toFlash = append(toFlash, utils.Point{i, j})
			}
		}
	}

	return toFlash
}
