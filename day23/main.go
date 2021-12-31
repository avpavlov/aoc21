package main

import . "aoc2021/utils"

func main() {
	Run([]TestCase{
		{"day23/test1.txt", 12521, solve_d23_1},
		{"day23/input1.txt", -1, solve_d23_1},
		{"day23/test2.txt", 44169, solve_d23_2},
		{"day23/input2.txt", -1, solve_d23_2},
	})
}
