package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day15/test.txt", 40, solve_d15_1},
		{"day15/input.txt", -1, solve_d15_1},
		{"day15/test.txt", 315, solve_d15_2},
		{"day15/input.txt", -1, solve_d15_2},
	})
}
