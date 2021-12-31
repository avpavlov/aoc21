package main

import "aoc2021/utils"

func solve_d03_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	ones, zeroes := utils.SplitByCondition(lines, splitByBitFn(0))
	var oxygens = utils.MaxByLen(ones, zeroes)
	var co2s = utils.MinByLen(ones, zeroes)

	var bits = len(lines[0])

	for bit := 1; len(oxygens) > 1 && bit < bits; bit++ {
		oxygens = utils.MaxByLen(utils.SplitByCondition(oxygens, splitByBitFn(bit)))
	}
	for bit := 1; len(co2s) > 1 && bit < bits; bit++ {
		co2s = utils.MinByLen(utils.SplitByCondition(co2s, splitByBitFn(bit)))
	}

	var oxygen = utils.BitsToNumber(oxygens[0])
	var co2 = utils.BitsToNumber(co2s[0])

	return int64(oxygen * co2)
}

func splitByBitFn(bit int) func(string) bool {
	return func(line string) bool { return line[bit] == '1' }
}
