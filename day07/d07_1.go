package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d07_1(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var crabs = map[int64]int64{}
	var positions = utils.ParseInt64s(strings.Split(lines[0], ","))
	var minP int64
	var maxP int64
	var costToZero int64 = 0
	var total int64 = 0
	for i, p := range positions {
		crabs[p]++
		if i == 0 {
			minP = p
			maxP = p
		} else if p > maxP {
			maxP = p
		} else if p < minP {
			minP = p
		}
		costToZero += p
		total++
	}

	var minCost = costToZero - minP*total
	//fmt.Printf("minCost=%d minP=%d costToZero=%d total=%d\n", minCost, minP, costToZero, total)
	var leftCost int64 = 0
	var totalLeft int64 = 0
	var rightCost int64 = 0
	var totalRight int64 = 0
	//fmt.Println(crabs)
	for p := minP; p <= maxP; p++ {
		var c = crabs[p]

		if p == minP {
			leftCost = 0
			totalLeft = 0
			rightCost = costToZero - minP*total
			totalRight = total - c
		} else {
			totalLeft += crabs[p-1]
			leftCost += totalLeft
			rightCost -= totalRight
			totalRight -= c
		}

		//fmt.Printf("p=%d c=%d totalLeft=%d totalRight=%d leftCost=%d rightCost=%d\n", p, c, totalLeft, totalRight, leftCost, rightCost)
		if c+totalRight+totalLeft != total {
			panic("Totals mismatched")
		}
		currentCost := leftCost + rightCost
		if minCost > currentCost {
			//fmt.Printf("p=%d minCost=%d currentCost=%d\n", p, minCost, currentCost)
			minCost = currentCost
		}
	}

	return minCost
}
