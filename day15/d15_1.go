package main

import (
	"aoc2021/utils"
	"container/heap"
	"strings"
)

func solve_d15_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var cave = [][]int64{}
	var path = [][]int64{}
	var lenY = int64(len(lines))
	var lenX = int64(len(lines[0]))
	cave = make([][]int64, lenY)
	path = make([][]int64, lenY)
	for y, line := range lines {
		cave[y] = make([]int64, lenX)
		path[y] = make([]int64, lenX)
		for x, risk := range utils.ParseInt64s(strings.Split(line, "")) {
			cave[y][x] = risk
			path[y][x] = -1
		}
	}

	var maxY = int64(len(cave) - 1)
	var maxX = int64(len(cave[0]) - 1)

	type Move struct {
		x, y, accumulatedRisk, priority int64
	}
	pq := utils.PriorityQueue{
		Items: []interface{}{},
		Comparator: func(i, j interface{}) bool {
			return (i).(Move).priority < (j).(Move).priority
		},
	}
	push := func(x, y, accumulatedRisk int64) {
		heap.Push(&pq, Move{x, y, accumulatedRisk, accumulatedRisk + cave[y][x]})
	}
	pop := func() (int64, int64, int64) {
		move := heap.Pop(&pq).(Move)
		return move.x, move.y, move.accumulatedRisk
	}
	heap.Init(&pq)

	push(0, 0, 0)
	for pq.Len() > 0 {
		x, y, accumulatedRisk := pop()
		riskToThisPoint := path[y][x]
		accumulatedRisk += cave[y][x]
		if riskToThisPoint == -1 || riskToThisPoint > accumulatedRisk {
			path[y][x] = accumulatedRisk
			if x < maxX {
				push(x+1, y, accumulatedRisk)
			}
			if y < maxY {
				push(x, y+1, accumulatedRisk)
			}
			if x > 0 {
				push(x-1, y, accumulatedRisk)
			}
			if y > 0 {
				push(x, y-1, accumulatedRisk)
			}
		}
	}

	return path[maxY][maxX] - path[0][0]
}
