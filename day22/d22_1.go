package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d22_1(filename string) int64 {
	type Command struct {
		on                     bool
		x1, x2, y1, y2, z1, z2 int64
	}
	var commands = []Command{}
	var lines = utils.ReadLines(filename)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		cmd := Command{}
		cmd.on = parts[0] == "on"

		parts = strings.Split(parts[1], ",")
		xs := utils.ParseInt64s(strings.Split(parts[0][2:], ".."))
		ys := utils.ParseInt64s(strings.Split(parts[1][2:], ".."))
		zs := utils.ParseInt64s(strings.Split(parts[2][2:], ".."))
		cmd.x1 = xs[0]
		cmd.x2 = xs[1]
		cmd.y1 = ys[0]
		cmd.y2 = ys[1]
		cmd.z1 = zs[0]
		cmd.z2 = zs[1]
		commands = append(commands, cmd)
	}

	var result int64
	for x := int64(-50); x <= 50; x++ {
		for y := int64(-50); y <= 50; y++ {
			for z := int64(-50); z <= 50; z++ {
				for c := len(commands) - 1; c >= 0; c-- {
					cmd := commands[c]
					if x >= cmd.x1 && x <= cmd.x2 && y >= cmd.y1 && y <= cmd.y2 && z >= cmd.z1 && z <= cmd.z2 {
						if cmd.on {
							result++
						}
						break
					}
				}
			}
		}
	}
	return result
}
