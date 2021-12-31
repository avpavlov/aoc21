package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func solve_d13_2(filename string) {
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
			fold(&dots, strings.HasSuffix(parts[0], "y"), utils.ParseInt64(parts[1]))
		}
	}

	var maxX = int64(0)
	var maxY = int64(0)
	for p := range dots {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	var canvas = make([][]string, maxY+1)
	for i := range canvas {
		canvas[i] = make([]string, maxX+1)
	}

	for p := range dots {
		canvas[p.Y][p.X] = "##"
	}

	for _, row := range canvas {
		var out = ""
		for _, cell := range row {
			if cell == "##" {
				out += cell
			} else {
				out += "  "
			}
		}
		fmt.Println(out)
	}
}
