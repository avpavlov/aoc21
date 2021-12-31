package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day11/test.txt", 1656, solve_d11_1},
		{"day11/input.txt", -1, solve_d11_1},
		{"day11/test.txt", 195, solve_d11_2},
		{"day11/input.txt", -1, solve_d11_2},
	})
}
