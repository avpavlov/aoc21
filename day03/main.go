package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day03/test.txt", 198, solve_d03_1},
		{"day03/input.txt", -1, solve_d03_1},
		{"day03/test.txt", 230, solve_d03_2},
		{"day03/input.txt", -1, solve_d03_2},
	})
}
