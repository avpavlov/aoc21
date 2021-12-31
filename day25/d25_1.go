package main

import (
	"aoc2021/utils"
)

func solve_d25_1(filename string) int64 {
	const empty1 = byte('.')
	const empty2 = byte(' ')
	const south = byte('v')
	const east = byte('>')
	var lines = utils.ReadLines(filename)
	var field = make([][]byte, len(lines))
	for y, line := range lines {
		row := make([]byte, len(line))
		field[y] = row
		for x := 0; x < len(line); x++ {
			spot := line[x]
			if spot == south || spot == east || spot == empty1 {
				row[x] = spot
			} else {
				panic("Unknown spot")
			}
		}
	}

	var touched = make([][]int, len(field))
	for y, row := range field {
		touched[y] = make([]int, len(row))
	}

	var maxY = len(field)
	var maxX = len(field[0])
	var steps = 0
	for {
		moved := 0

		empty, nextEmpty := empty1, empty2
		for y, row := range field {
			for x, spot := range row {
				if spot != empty {
					continue
				}
				left := (x - 1 + maxX) % maxX
				if row[left] == east && touched[y][left] <= steps {
					row[x] = east
					row[left] = nextEmpty
					moved++
					touched[y][x] = steps + 1
				} else {
					row[x] = nextEmpty
				}
			}
		}

		empty, nextEmpty = empty2, empty1
		for y, row := range field {
			for x, spot := range row {
				if spot != empty {
					continue
				}
				top := (y - 1 + maxY) % maxY
				if field[top][x] == south && touched[top][x] <= steps {
					row[x] = south
					field[top][x] = nextEmpty
					moved++
					touched[y][x] = steps + 1
				} else {
					row[x] = nextEmpty
				}
			}
		}

		if moved == 0 {
			return int64(steps + 1)
		} else {
			steps++
		}
	}
}
