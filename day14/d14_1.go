package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d14_1(filename string) int64 {
	return solve_d14(filename, 10)
}

func solve_d14_2(filename string) int64 {
	return solve_d14(filename, 40)
}

func solve_d14(filename string, steps int) int64 {
	var lines = utils.ReadLines(filename)
	var polymer = map[string]int64{}
	var counts = map[uint8]int64{}
	var rules = map[string]uint8{}
	for i, line := range lines {
		if i == 0 {
			for j := 0; j < len(line)-1; j++ {
				polymer[line[j:j+2]]++
				counts[line[j]]++
			}
			counts[line[len(line)-1]]++
		} else if i > 1 {
			var parts = strings.Split(line, " -> ")
			rules[parts[0]] = parts[1][0]
		}
	}

	for step := 0; step < steps; step++ {
		var newPolymer = map[string]int64{}
		for pair, count := range polymer {
			if insert, ok := rules[pair]; ok {
				newPolymer[string([]uint8{pair[0], insert})] += count
				newPolymer[string([]uint8{insert, pair[1]})] += count
				counts[insert] += count
			}
		}
		polymer = newPolymer
	}

	var min = int64(-1)
	var max = int64(0)

	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min || min == -1 {
			min = v
		}
	}
	return max - min
}
