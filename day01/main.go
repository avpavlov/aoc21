package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day01/test.txt", 7, solve_d01_1},
		{"day01/input.txt", -1, solve_d01_1},
		{"day01/test.txt", 5, solve_d01_2},
		{"day01/input.txt", -1, solve_d01_2},
	})
}
