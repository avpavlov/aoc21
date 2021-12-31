package main

import (
	. "aoc2021/utils"
	"container/heap"
	"strings"
)

type AmphipodType byte

const Any = AmphipodType(0)
const A = AmphipodType('A')
const B = AmphipodType('B')
const C = AmphipodType('C')
const D = AmphipodType('D')

var spots = map[Point]AmphipodType{
	{0, 0}:  Any,
	{1, 0}:  Any,
	{2, 1}:  A,
	{2, 2}:  A,
	{3, 0}:  Any,
	{4, 1}:  B,
	{4, 2}:  B,
	{5, 0}:  Any,
	{6, 1}:  C,
	{6, 2}:  C,
	{7, 0}:  Any,
	{8, 1}:  D,
	{8, 2}:  D,
	{9, 0}:  Any,
	{10, 0}: Any,
}

var costs = map[AmphipodType]int64{A: 1, B: 10, C: 100, D: 1000}

var HallwaySize = 11
var RoomSize = 2

type GameState struct {
	cost      int64
	amphipods map[Point]AmphipodType
	previous  *GameState
}

func solve_d23_1(filename string) int64 {
	var lines = ReadLines(filename)
	var game = GameState{
		cost:      0,
		amphipods: make(map[Point]AmphipodType, RoomSize*4),
	}
	for y, line := range lines {
		if y >= 2 && y < 2+RoomSize {
			for x, s := range strings.Split(line[3:10], "#") {
				game.amphipods[Point{int64(x+1) * 2, int64(y - 1)}] = AmphipodType(s[0])
			}
		}
	}

	gameStateCosts := map[string]int64{}

	pq := PriorityQueue{
		Items: []interface{}{},
		Comparator: func(i, j interface{}) bool {
			return (i).(*GameState).cost < (j).(*GameState).cost
		},
	}
	push := func(gs *GameState) bool {
		key := gs.key()
		if existingCost, found := gameStateCosts[key]; !found || existingCost > gs.cost {
			gameStateCosts[key] = gs.cost
			heap.Push(&pq, gs)
			return true
		} else {
			return false
		}
	}
	pop := func() *GameState {
		return heap.Pop(&pq).(*GameState)
	}

	heap.Init(&pq)
	push(&game)
	var minimalCost int64 = 999_999_999_999
states:
	for pq.Len() > 0 {
		gs := pop()
		for position, kind := range gs.amphipods {
			for destination, acceptedKind := range spots {
				if gs.canEnterSpot(position, kind, destination, acceptedKind) {
					amphipods, gameOver := gs.move(position, destination)
					newGameState := GameState{
						cost:      gs.cost + costs[kind]*(AbsInt64(position.X-destination.X)+position.Y+destination.Y),
						amphipods: amphipods,
						previous:  gs,
					}
					if gameOver {
						minimalCost = MinInt64(minimalCost, newGameState.cost)
						//if minimalCost == newGameState.cost {
						//	fmt.Println("------------------------------------")
						//	newGameState.printChain()
						//}
						continue states
					} else if push(&newGameState) {
					}
				}
			}
		}
	}
	return minimalCost
}
