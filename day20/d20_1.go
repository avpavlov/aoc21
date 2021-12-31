package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d20_1(filename string) int64 {
	return solve_d20(filename, 2)
}

func solve_d20_2(filename string) int64 {
	return solve_d20(filename, 50)
}

func solve_d20(filename string, maxSteps int) int64 {
	var lines = utils.ReadLines(filename)
	var bright = &utils.Image{}
	var imageImprov = map[int]int{}
	for y, line := range lines {
		if y == 0 {
			for x, cell := range strings.Split(line, "") {
				if cell == "#" {
					imageImprov[x] = 1
				} else if cell != "." {
					panic("Unknown cell " + cell)
				}
			}
		} else if y == 1 {
			continue
		} else {
			for x, cell := range strings.Split(line, "") {
				if cell == "#" {
					(*bright)[utils.Point{int64(x), int64(y - 2)}] = 1
				} else if cell != "." {
					panic("Unknown cell " + cell)
				}
			}
		}
	}

	var infinities = []int{imageImprov[0], 0}
	for step := 0; step < maxSteps; step++ {
		bright = bright.ImproveImage(&imageImprov, infinities[(step+1)%2])
	}

	var result int
	for _, bit := range *bright {
		result += bit
	}
	return int64(result)
}
