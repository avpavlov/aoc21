package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day10/test.txt", 26397, solve_d10_1},
		{"day10/input.txt", -1, solve_d10_1},
		{"day10/test.txt", 288957, solve_d10_2},
		{"day10/input.txt", -1, solve_d10_2},
	})
}
