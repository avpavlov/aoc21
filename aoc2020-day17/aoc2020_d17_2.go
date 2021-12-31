package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_aoc2020_d17_2(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var alive = map[utils.Point4D]bool{}
	for y, line := range lines {
		for x, cell := range strings.Split(line, "") {
			if cell == "#" {
				alive[utils.Point4D{x, y, 0, 0}] = true
			}
		}
	}

	for step := 0; step < 6; step++ {
		var survivors = map[utils.Point4D]bool{}
		var newBorn = map[utils.Point4D]int{}
		for cell, _ := range alive {
			var neighbors int
			for i := 0; i < 81; i++ {
				x := cell.X + (i % 3) - 1
				y := cell.Y + ((i / 3) % 3) - 1
				z := cell.Z + ((i / 9) % 3) - 1
				w := cell.W + ((i / 27) % 3) - 1
				var p = utils.Point4D{x, y, z, w}
				if p == cell {
					continue
				}
				if _, ok := alive[p]; ok {
					neighbors++
				} else {
					newBorn[p]++
				}
			}
			if neighbors >= 2 && neighbors <= 3 {
				survivors[cell] = true
			}
		}
		for cell, neighbors := range newBorn {
			if neighbors == 3 {
				survivors[cell] = true
			}
		}
		alive = survivors
	}

	return int64(len(alive))
}
