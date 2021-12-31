package main

import (
	. "aoc2021/utils"
	"container/heap"
	"strings"
)

func solve_d23_2(filename string) int64 {
	spots = map[Point]AmphipodType{
		{0, 0}:  Any,
		{1, 0}:  Any,
		{2, 1}:  A,
		{2, 2}:  A,
		{2, 3}:  A,
		{2, 4}:  A,
		{3, 0}:  Any,
		{4, 1}:  B,
		{4, 2}:  B,
		{4, 3}:  B,
		{4, 4}:  B,
		{5, 0}:  Any,
		{6, 1}:  C,
		{6, 2}:  C,
		{6, 3}:  C,
		{6, 4}:  C,
		{7, 0}:  Any,
		{8, 1}:  D,
		{8, 2}:  D,
		{8, 3}:  D,
		{8, 4}:  D,
		{9, 0}:  Any,
		{10, 0}: Any,
	}
	RoomSize = 4

	var lines = ReadLines(filename)
	var game = GameState{
		cost:      0,
		amphipods: make(map[Point]AmphipodType, RoomSize*4),
	}
	for y, line := range lines {
		if y >= 2 && y < 2+RoomSize {
			for x, s := range strings.Split(line[3:10], "#") {
				game.amphipods[Point{X: int64(x+1) * 2, Y: int64(y - 1)}] = AmphipodType(s[0])
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
	push := func(gs *GameState) {
		key := gs.key()
		if existingCost, found := gameStateCosts[key]; !found || existingCost > gs.cost {
			gameStateCosts[key] = gs.cost
			heap.Push(&pq, gs)
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
		if gs.cost >= minimalCost {
			continue
		}
		for position, kind := range gs.amphipods {
			for destination, acceptedKind := range spots {
				newCost := gs.cost + costs[kind]*(AbsInt64(position.X-destination.X)+position.Y+destination.Y)
				if newCost < minimalCost && gs.canEnterSpot(position, kind, destination, acceptedKind) {
					amphipods, gameOver := gs.move(position, destination)
					if gameOver {
						minimalCost = MinInt64(minimalCost, newCost)
						continue states
					} else {
						newGameState := GameState{
							cost:      newCost,
							amphipods: amphipods,
						}
						push(&newGameState)
					}
				}
			}
		}
	}
	return minimalCost
}

func (game *GameState) canEnterSpot(position Point, kind AmphipodType, destination Point, acceptedKind AmphipodType) bool {
	if acceptedKind != Any && acceptedKind != kind {
		return false
	}
	if position.Y == 0 && destination.Y == 0 {
		return false
	}
	// can exit room
	for y := position.Y - 1; y > 0; y-- {
		if _, occupied := game.amphipods[Point{position.X, y}]; occupied {
			return false
		}
	}
	// can move along hallway
	for x, maxX := MinMax(position.X, destination.X); x <= maxX; x++ {
		if x != position.X {
			if _, occupied := game.amphipods[Point{x, 0}]; occupied {
				return false
			}
		}
	}
	// can enter room
	if destination.Y > 0 {
		for y := int64(1); y <= destination.Y; y++ {
			if _, occupied := game.amphipods[Point{destination.X, y}]; occupied {
				return false
			}
		}
		// the rest is occupied by own kind
		for y := destination.Y + 1; y <= int64(RoomSize); y++ {
			occupiedBy, occupied := game.amphipods[Point{destination.X, y}]
			if !occupied || occupiedBy != kind {
				return false
			}
		}
	}

	return true
}

func (game *GameState) key() string {
	rows := make([][]byte, RoomSize+1)
	for r := 0; r <= RoomSize; r++ {
		rows[r] = make([]byte, HallwaySize)
		for c := 0; c < HallwaySize; c++ {
			rows[r][c] = ' '
		}
	}
	for spot := range spots {
		rows[spot.Y][spot.X] = '.'
	}
	for point, kind := range game.amphipods {
		rows[point.Y][point.X] = byte(kind)
	}
	result := ""
	for _, row := range rows {
		result += string(row) + "\n"
	}
	return result
}

func (game *GameState) move(from Point, to Point) (map[Point]AmphipodType, bool) {
	gameOver := true
	m2 := make(map[Point]AmphipodType, len(game.amphipods))
	for position, kind := range game.amphipods {
		if position == from {
			m2[to] = kind
			gameOver = gameOver && spots[to] == kind
		} else {
			m2[position] = kind
			gameOver = gameOver && spots[position] == kind
		}
	}
	return m2, gameOver
}
