package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day06/test.txt", 5934, solve_d06_1},
		{"day06/input.txt", -1, solve_d06_1},
		{"day06/test.txt", 26984457539, solve_d06_2},
		{"day06/input.txt", -1, solve_d06_2},
	})
}
