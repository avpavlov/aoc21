package main

import (
	"aoc2021/utils"
	"log"
	"strings"
)

func solve_d08_2(filename string) int64 {
	var lines = utils.ReadLines(filename)

	var result int64 = 0
	for _, line := range lines {
		var parts = strings.Split(line, " | ")
		d1, d7, d4, len5, len6, d8 := groupByLen(strings.Split(strings.TrimSpace(parts[0]), " "))
		var a = minus(d7, d1)
		var cf = minus(d7, a)
		var bd = minus(d4, cf)
		var g, _ = findAndMinus(len5, a+d4)
		var d9 = utils.SortLetters(a + bd + cf + g)
		var e, _ = findAndMinus(len6, d9)
		var f, _ = findAndMinus(len6, a+bd+e+g)
		var c = minus(cf, f)
		var d5 = utils.SortLetters(a + bd + f + g)
		var d6 = utils.SortLetters(a + bd + e + f + g)
		var b, d0 = findAndMinus(len6, a+cf+e+g)
		var d = minus(bd, b)
		var d2 = utils.SortLetters(a + c + d + e + g)
		var d3 = utils.SortLetters(a + c + d + f + g)
		//fmt.Printf("a=%s b=%s c=%s d=%s e=%s f=%s g=%s cf=%s bd=%s\n", a, b, c, d, e, f, g, cf, bd)
		m := map[string]int{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7, d8: 8, d9: 9}
		//fmt.Println(m)
		var display = strings.Split(strings.TrimSpace(parts[1]), " ")
		var num = 0
		for _, d := range display {
			num = num*10 + m[utils.SortLetters(d)]
		}
		//fmt.Printf("%s %d\n", display, num)
		result += int64(num)
	}
	return result
}

func findAndMinus(candidates []string, sub string) (string, string) {
	var c string
	for _, s := range candidates {
		c = minus(s, sub)
		if len(c) == 1 {
			return c, s
		}
	}
	log.Panicf("Cannot be here %s %s\n", candidates, sub)
	return "", ""
}

func groupByLen(seq []string) (string, string, string, []string, []string, string) {
	var s2 string
	var s3 string
	var s4 string
	var s5 = []string{}
	var s6 = []string{}
	var s7 string
	for _, s := range seq {
		s = utils.SortLetters(s)
		switch len(s) {
		case 2:
			s2 = s
			break
		case 3:
			s3 = s
			break
		case 4:
			s4 = s
			break
		case 5:
			s5 = append(s5, s)
			break
		case 6:
			s6 = append(s6, s)
			break
		case 7:
			s7 = s
			break
		}
	}
	return s2, s3, s4, s5, s6, s7
}

func minus(s1 string, s2 string) string {
	for i := range s2 {
		s1 = strings.ReplaceAll(s1, string([]byte{s2[i]}), "")
	}
	return s1
}
