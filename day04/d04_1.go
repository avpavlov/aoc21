package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d04_1(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var boards = []Board{}
	for line := 2; line < len(lines); {
		boards = append(boards, parseBoard(lines, line))
		line += S + 1 // empty line
	}

	var numbers = utils.ParseInt64s(strings.Split(lines[0], ","))
	for _, number := range numbers {
		for _, board := range boards {
			if board.play(number) {
				return number * board.multiplyUnplayed()
			}
		}
	}

	panic("cannot be here")
}

func parseBoard(lines []string, startLine int) Board {
	var board = Board{}
	board.init()
	for i := 0; i < S; i++ {
		var line = lines[startLine+i]
		var numbers = utils.ParseInt64s(strings.Split(utils.ReplaceAll(strings.TrimSpace(line), "  ", " "), " "))
		for j, number := range numbers {
			board.rows[i][number] = true
			board.cols[j][number] = true
		}
	}
	return board
}

const S = 5

type vector map[int64]bool
type Board struct {
	rows [S]vector
	cols [S]vector
}

func (b *Board) init() {
	(*b).rows = [S]vector{}
	(*b).cols = [S]vector{}
	for i := 0; i < S; i++ {
		(*b).rows[i] = vector{}
		(*b).cols[i] = vector{}
	}
}

func (b *Board) play(number int64) bool {
	for i := 0; i < S; i++ {
		delete((*b).rows[i], number)
		delete((*b).cols[i], number)
	}
	for i := 0; i < S; i++ {
		if len((*b).rows[i]) == 0 || len((*b).cols[i]) == 0 {
			return true
		}
	}
	return false
}

func (b *Board) multiplyUnplayed() int64 {
	var res int64 = 0
	for i := 0; i < S; i++ {
		for k := range (*b).rows[i] {
			res += k
		}
	}
	return res
}
