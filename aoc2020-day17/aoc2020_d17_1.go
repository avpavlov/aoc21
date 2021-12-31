package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_aoc2020_d17_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var alive = map[utils.Point3D]bool{}
	for y, line := range lines {
		for x, cell := range strings.Split(line, "") {
			if cell == "#" {
				alive[utils.Point3D{x, y, 0}] = true
			}
		}
	}

	for step := 0; step < 6; step++ {
		var survivors = map[utils.Point3D]bool{}
		var newBorn = map[utils.Point3D]int{}
		for cell, _ := range alive {
			var neighbors int
			for i := 0; i < 27; i++ {
				x := cell.X + (i % 3) - 1
				y := cell.Y + ((i / 3) % 3) - 1
				z := cell.Z + ((i / 9) % 3) - 1
				var p = utils.Point3D{x, y, z}
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
