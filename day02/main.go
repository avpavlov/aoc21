package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day02/test.txt", 150, solve_d02_1},
		{"day02/input.txt", -1, solve_d02_1},
		{"day02/test.txt", 900, solve_d02_2},
		{"day02/input.txt", -1, solve_d02_2},
	})
}
