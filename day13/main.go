package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day13/test.txt", 17, solve_d13_1},
		{"day13/input.txt", -1, solve_d13_1},
	})
	solve_d13_2("day13/input.txt")
}
