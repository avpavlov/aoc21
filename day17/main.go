package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day17/test.txt", 45, solve_d17_1},
		{"day17/input.txt", -1, solve_d17_1},
		{"day17/test.txt", 112, solve_d17_2},
		{"day17/input.txt", -1, solve_d17_2},
	})
}
