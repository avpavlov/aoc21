package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day21/test.txt", 739785, solve_d21_1},
		{"day21/input.txt", -1, solve_d21_1},
		{"day21/test.txt", 444356092776315, solve_d21_2},
		{"day21/input.txt", -1, solve_d21_2},
	})
}
