package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d21_2(filename string) int64 {
	type Player struct {
		pos int64
	}
	players := []Player{}
	var lines = utils.ReadLines(filename)
	for _, line := range lines {
		p := Player{
			pos: utils.ParseInt64(strings.Split(line, "starting position: ")[1]),
		}
		players = append(players, p)
	}

	dice := make([]byte, 27)
	for i := byte(0); i < 27; i++ {
		d1 := (i % 3) + 1
		d2 := ((i / 3) % 3) + 1
		d3 := ((i / 9) % 3) + 1
		dice[i] = d1 + d2 + d3
	}

	wins := [2]int64{0, 0}

	type Uni struct {
		move1  bool
		pos1   byte
		score1 byte
		pos2   byte
		score2 byte
	}
	getPlayer := func(uni *Uni) (pos byte, score byte) {
		if uni.move1 {
			return uni.pos1, uni.score1
		} else {
			return uni.pos2, uni.score2
		}
	}
	nextUni := func(uni *Uni, pos byte, score byte) Uni {
		if uni.move1 {
			return Uni{false, pos, score, uni.pos2, uni.score2}
		} else {
			return Uni{true, uni.pos1, uni.score1, pos, score}
		}
	}
	universes := map[Uni]int64{}
	universes[Uni{true, byte(players[0].pos), 0, byte(players[1].pos), 0}] = 1
	for len(universes) > 0 {
		for uni, count := range universes {
			delete(universes, uni)
			pos, score := getPlayer(&uni)
			for _, d := range dice {
				nextPos := (pos-1+d)%10 + 1
				nextScore := score + nextPos
				if nextScore >= 21 {
					if uni.move1 {
						wins[0] += count
					} else {
						wins[1] += count
					}
				} else {
					universes[nextUni(&uni, nextPos, nextScore)] += count
				}
			}
		}
	}

	_, winner := utils.MinMax(wins[0], wins[1])
	return winner
}
