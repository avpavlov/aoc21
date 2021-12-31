package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day08/test.txt", 26, solve_d08_1},
		{"day08/input.txt", -1, solve_d08_1},
		{"day08/test.txt", 61229, solve_d08_2},
		{"day08/input.txt", -1, solve_d08_2},
	})
}
