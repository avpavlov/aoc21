package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day12/test.txt", 226, solve_d12_1},
		{"day12/input.txt", -1, solve_d12_1},
		{"day12/test.txt", 3509, solve_d12_2},
		{"day12/input.txt", -1, solve_d12_2},
	})
}
