package main

import (
	"aoc2021/utils"
	"math"
	"strings"
)

func solve_d17_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	_, coords := utils.SplitPair(lines[0], ": ")
	targetXRangeStr, targetYRangeStr := utils.SplitPair(coords, ", ")
	targetXRange := utils.ParseInt64s(strings.Split(targetXRangeStr[2:], ".."))
	targetXMin := targetXRange[0]
	targetXMax := targetXRange[1]
	targetYRange := utils.ParseInt64s(strings.Split(targetYRangeStr[2:], ".."))
	targetYMin := targetYRange[0]
	targetYMax := targetYRange[1]

	xCanditates := map[int64]bool{}
	for L := targetXMin; L <= targetXMax; L++ {
		var x = int64(math.Round((-1 + math.Sqrt(float64(1+8*L))) / 2))
		if checkX := (x * (x + 1)) / 2; checkX < targetXMin || checkX > targetXMax {
			continue
		}
		xCanditates[x] = true
	}

	freeFallPoints := map[int64]bool{}
	ffp := int64(0)
	for i := int64(1); ffp < 1_000_000; i++ {
		freeFallPoints[ffp] = true
		ffp += i
	}

	var result int64 = 0
	//for X := range xCanditates {
	for y := int64(1); y < 1_000_000; y++ {
		height := y * (y + 1) / 2
		//fmt.Printf("X=%d Y=%d height=%d\n", X, Y, height)
		if height > result {
			for H := targetYMin; H <= targetYMax; H++ {
				ffp := height - H
				//fmt.Printf("H=%d ffp=%d %t\n", H, ffp, freeFallPoints[ffp])
				if freeFallPoints[ffp] {
					result = height
					break
				}
			}
		}
	}
	//}

	return result
}
