package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day20/test.txt", 35, solve_d20_1},
		{"day20/input.txt", -1, solve_d20_1},
		{"day20/test.txt", 3351, solve_d20_2},
		{"day20/input.txt", -1, solve_d20_2},
	})
}
