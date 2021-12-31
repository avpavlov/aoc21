package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d13_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var dots = map[utils.Point]int{}
	var parseDots = true
	for _, line := range lines {
		if line == "" {
			parseDots = false
			continue
		}
		if parseDots {
			var parts = utils.ParseInt64s(strings.Split(line, ","))
			var p = utils.Point{parts[0], parts[1]}
			dots[p]++
		} else {
			var parts = strings.Split(line, "=")
			return fold(&dots, strings.HasSuffix(parts[0], "y"), utils.ParseInt64(parts[1]))
		}
	}

	panic("Must not be here")
}

func fold(dots *map[utils.Point]int, foldY bool, foldAt int64) int64 {
	var base = 2 * foldAt
	var toDelete = []utils.Point{}
	for p, v := range *dots {
		if foldY {
			if p.Y < foldAt {
				continue
			} else if p.Y == foldAt {
				toDelete = append(toDelete, p)
			} else {
				toDelete = append(toDelete, p)
				p := utils.Point{p.X, base - p.Y}
				(*dots)[p] += v
			}
		} else {
			if p.X < foldAt {
				continue
			} else if p.X == foldAt {
				toDelete = append(toDelete, p)
			} else {
				toDelete = append(toDelete, p)
				p := utils.Point{base - p.X, p.Y}
				(*dots)[p] += v
			}
		}
	}
	for _, p := range toDelete {
		delete(*dots, p)
	}
	return int64(len(*dots))
}
