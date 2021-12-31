package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day14/test.txt", 1588, solve_d14_1},
		{"day14/input.txt", -1, solve_d14_1},
		{"day14/test.txt", 2188189693529, solve_d14_2},
		{"day14/input.txt", -1, solve_d14_2},
	})
}
