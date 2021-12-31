package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day04/test.txt", 4512, solve_d04_1},
		{"day04/input.txt", -1, solve_d04_1},
		{"day04/test.txt", 1924, solve_d04_2},
		{"day04/input.txt", -1, solve_d04_2},
	})
}
