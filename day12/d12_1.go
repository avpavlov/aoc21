package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d12_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var caves = map[string]*Cave{}
	for _, line := range lines {
		var parts = strings.Split(line, "-")
		from, ok := caves[parts[0]]
		if !ok {
			var c = Cave{id: parts[0], big: strings.ToUpper(parts[0]) == parts[0], end: parts[0] == "end", edges: make([]*Cave, 0)}
			from = &c
			caves[c.id] = from
		}
		to, ok := caves[parts[1]]
		if !ok {
			var c = Cave{id: parts[1], big: strings.ToUpper(parts[1]) == parts[1], end: parts[1] == "end", edges: make([]*Cave, 0)}
			to = &c
			caves[c.id] = to
		}
		from.edges = append(from.edges, to)
		if !to.end && from.id != "start" {
			to.edges = append(to.edges, from)
		}
	}

	var game = CavesGame{caves: &caves, result: 0}
	var start = caves["start"]
	visitCave(nil, start, &game)
	return game.result
}

func visitCave(prev *Cave, current *Cave, game *CavesGame) {
	if current.visited {
		return
	}
	if current.end {
		game.result++
		return
	}

	if !current.big {
		current.visited = true
	}

	for _, edge := range current.edges {
		if current.big && edge.big && prev != nil && prev.id == edge.id {
			continue
		}
		visitCave(current, edge, game)
	}

	defer func() { current.visited = false }()
}

type CavesGame struct {
	caves             *map[string]*Cave
	smallVisitedTwice bool
	result            int64
}

type Cave struct {
	id      string
	big     bool
	end     bool
	edges   []*Cave
	visited bool
}
