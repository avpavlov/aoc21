package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day25/test.txt", 58, solve_d25_1},
		{"day25/input.txt", -1, solve_d25_1},
	})
}
