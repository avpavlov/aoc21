package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day16/test1.txt", 16, solve_d16_1},
		{"day16/test2.txt", 12, solve_d16_1},
		{"day16/test3.txt", 23, solve_d16_1},
		{"day16/test4.txt", 31, solve_d16_1},
		{"day16/input.txt", -1, solve_d16_1},
		{"day16/test5.txt", 1, solve_d16_2},
		{"day16/input.txt", -1, solve_d16_2},
	})
}
