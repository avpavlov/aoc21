package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d07_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var crabs = map[int64]int64{}
	var positions = utils.ParseInt64s(strings.Split(lines[0], ","))
	var minP int64
	var maxP int64
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
	}
	//fmt.Println(crabs)

	var minCost int64 = 0
	for p := minP; p <= maxP; p++ {
		lc := leftCost(crabs, p-1, minP)
		//fmt.Printf("p=%d lc=%d from=%d to=%d\n", p, lc, p-1, minP)
		rc := rightCost(crabs, p+1, maxP)
		//fmt.Printf("p=%d rc=%d from=%d to=%d\n", p, rc, p+1, maxP)

		currentCost := lc + rc
		if p == minP || minCost > currentCost {
			minCost = currentCost
			//fmt.Printf("minCost=%d at=%d\n", minCost, p)
		}
	}

	return minCost
}

func leftCost(crabs map[int64]int64, from int64, to int64) int64 {
	var result int64 = 0
	var steps int64 = 0
	for p := from; p >= to; p-- {
		var c = crabs[p]
		steps += (from - p + 1)
		result += c * steps
	}
	return result
}

func rightCost(crabs map[int64]int64, from int64, to int64) int64 {
	var result int64 = 0
	var steps int64 = 0
	for p := from; p <= to; p++ {
		var c = crabs[p]
		steps += (p - from + 1)
		result += c * steps
	}
	return result
}
