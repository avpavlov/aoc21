package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day18/test.txt", 4140, solve_d18_1},
		{"day18/input.txt", -1, solve_d18_1},
		{"day18/test.txt", 3993, solve_d18_2},
		{"day18/input.txt", -1, solve_d18_2},
	})
}
