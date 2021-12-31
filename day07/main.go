package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day07/test.txt", 37, solve_d07_1},
		{"day07/input.txt", -1, solve_d07_1},
		{"day07/test.txt", 168, solve_d07_2},
		{"day07/input.txt", -1, solve_d07_2},
	})
}
