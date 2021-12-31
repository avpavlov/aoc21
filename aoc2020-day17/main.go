package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"aoc2020-day17/test.txt", 112, solve_aoc2020_d17_1},
		{"aoc2020-day17/input.txt", -1, solve_aoc2020_d17_1},
		{"aoc2020-day17/test.txt", 848, solve_aoc2020_d17_2},
		{"aoc2020-day17/input.txt", -1, solve_aoc2020_d17_2},
	})
}
