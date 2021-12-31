package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d21_1(filename string) int64 {
	type Player struct {
		pos   int64
		score int64
		moves int64
	}
	players := []Player{}
	var lines = utils.ReadLines(filename)
	for _, line := range lines {
		p := Player{
			pos: utils.ParseInt64(strings.Split(line, "starting position: ")[1]),
		}
		players = append(players, p)
	}

	dice := int64(1)
	pi := 0
	for {
		p := &players[pi%2]
		moves := rollDice(&dice) + rollDice(&dice) + rollDice(&dice)
		p.pos = (p.pos-1+moves)%10 + 1
		p.score += p.pos
		p.moves += 3
		if p.score >= 1000 {
			break
		}
		pi++
	}

	loser, _ := utils.MinMax(players[0].score, players[1].score)
	return loser * (players[0].moves + players[1].moves)
}

func rollDice(dice *int64) int64 {
	d := *dice
	*dice = (d % 100) + 1
	return d
}
