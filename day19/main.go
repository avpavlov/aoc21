package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day19/test.txt", 79, solve_d19_1},
		{"day19/input.txt", -1, solve_d19_1},
		{"day19/test.txt", 3621, solve_d19_2},
		{"day19/input.txt", -1, solve_d19_2},
	})
}
