package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day09/test.txt", 15, solve_d09_1},
		{"day09/input.txt", -1, solve_d09_1},
		{"day09/test.txt", 1134, solve_d09_2},
		{"day09/input.txt", -1, solve_d09_2},
	})
}
