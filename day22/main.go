package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day22/test1.txt", 590784, solve_d22_1},
		{"day22/input.txt", -1, solve_d22_1},
		{"day22/test2.txt", 2758514936282235, solve_d22_2},
		{"day22/input.txt", -1, solve_d22_2},
	})
}
