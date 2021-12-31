package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d04_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var boards = []Board{}
	for line := 2; line < len(lines); {
		boards = append(boards, parseBoard(lines, line))
		line += S + 1 // empty line
	}

	var numbers = utils.ParseInt64s(strings.Split(lines[0], ","))
	for _, number := range numbers {
		var newBoards = []Board{}
		for _, board := range boards {
			if board.play(number) {
				if len(boards) == 1 {
					return number * board.multiplyUnplayed()
				}
			} else {
				newBoards = append(newBoards, board)
			}
		}
		boards = newBoards
	}

	panic("cannot be here")
}
