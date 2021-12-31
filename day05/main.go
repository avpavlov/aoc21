package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day05/test.txt", 5, solve_d05_1},
		{"day05/input.txt", -1, solve_d05_1},
		{"day05/test.txt", 12, solve_d05_2},
		{"day05/input.txt", -1, solve_d05_2},
	})
}
