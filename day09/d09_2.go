package main

import (
	"aoc2021/utils"
	Sort "sort"
	"strings"
)

func solve_d09_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var levels = [][]int64{}
	for _, line := range lines {
		var parts = utils.ParseInt64s(strings.Split(line, ""))
		levels = append(levels, parts)
	}

	var basins = []int64{}
	for r, row := range levels {
		for c, value := range row {
			if value != -1 && value != 9 {
				basins = append(basins, basinSize(&levels, r, c))
			}
		}
	}

	Sort.Slice(basins, func(i, j int) bool { return basins[i] < basins[j] })
	var max3 = basins[len(basins)-3:]
	return max3[0] * max3[1] * max3[2]
}

func basinSize(levels *[][]int64, r int, c int) int64 {
	if r < 0 || c < 0 || r == len(*levels) || c == len((*levels)[r]) {
		return 0
	}
	v := (*levels)[r][c]
	if v == -1 || v == 9 {
		return 0
	}
	(*levels)[r][c] = -1
	return 1 + basinSize(levels, r, c-1) + basinSize(levels, r, c+1) + basinSize(levels, r-1, c) + basinSize(levels, r+1, c)
}
