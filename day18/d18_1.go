package main

import (
	"aoc2021/utils"
	"container/list"
	"fmt"
)

func solve_d18_1(filename string) int64 {
	var lines = utils.ReadLines(filename)
	var sum *SnailfishNum
	for _, line := range lines {
		copy := line
		num := parseSnailfishNum(&line)
		if p := num.printSnailfishNum(); p != copy {
			panic(fmt.Sprintf("Diff %s %s", copy, p))
		}

		if sum == nil {
			sum = num
		} else {
			sum = sum.addSnailfishNum(num)
		}
		for sum.reduceSnailfishNum() {
		}
	}

	return sum.magnitude()
}

type SnailfishNum struct {
	is_regular bool
	regular    int64
	left       *SnailfishNum
	right      *SnailfishNum
	parent     *SnailfishNum
	level      int
}

func parseSnailfishNum(text *string) *SnailfishNum {
	num := SnailfishNum{}
	first := utils.Cut(text, 1)
	if first == "[" {
		num.left = parseSnailfishNum(text)
		comma := utils.Cut(text, 1)
		if comma != "," {
			panic(fmt.Sprintf("Expected comma, got %s. Length of remainder %d, remainder %s", comma, len(*text), *text))
		}
		num.right = parseSnailfishNum(text)
		close := utils.Cut(text, 1)
		if close != "]" {
			panic(fmt.Sprintf("Expected ], got %s. Length of remainder %d, remainder %s", close, len(*text), *text))
		}
	} else if utils.IsDigit(first[0]) {
		num.is_regular = true
		num.regular = utils.ParseInt64(first + utils.CutDigits(text))
	} else {
		panic(fmt.Sprintf("Expected comma, [ or number, got %s. Length of remainder %d, remainder %s", first, len(*text), *text))
	}
	return &num
}

func (num *SnailfishNum) printSnailfishNum() string {
	if num.is_regular {
		return fmt.Sprintf("%d", num.regular)
	} else {
		return fmt.Sprintf("[%s,%s]", num.left.printSnailfishNum(), num.right.printSnailfishNum())
	}
}

func (num *SnailfishNum) reduceSnailfishNum() bool {
	//fmt.Println(num.printSnailfishNum())
	all := list.New()
	num.linearize(all, 0)
	//println(num.printSnailfishNum())
	//for e := all.Front(); e != nil; e = e.Next() {
	//	num := e.Value.(*SnailfishNum)
	//	fmt.Printf("Check is_reg=%t lev=%d num=%s\n", num.is_regular, num.level, num.printSnailfishNum())
	//}
	if reduce(all, true) || reduce(all, false) {
		return true
	} else {
		//println("No changes")
		return false
	}
}

func (num *SnailfishNum) linearize(list *list.List, level int) {
	num.level = level
	if num.is_regular {
		list.PushBack(num)
	} else {
		num.left.linearize(list, level+1)
		num.left.parent = num

		list.PushBack(num)

		num.right.linearize(list, level+1)
		num.right.parent = num
	}
}

func reduce(list *list.List, explode bool) bool {
	for e := list.Front(); e != nil; e = e.Next() {
		num := e.Value.(*SnailfishNum)
		//fmt.Printf("Check is_reg=%t lev=%d num=%s\n", num.is_regular, num.level, num.printSnailfishNum())
		if num.is_regular && num.regular >= 10 && !explode {
			//fmt.Printf("split %s\n", num.printSnailfishNum())
			num.is_regular = false
			num.left = &SnailfishNum{
				is_regular: true,
				regular:    num.regular / 2,
				level:      num.level + 1,
			}
			num.right = &SnailfishNum{
				is_regular: true,
				regular:    num.regular - num.regular/2,
				level:      num.level + 1,
			}
			list.InsertBefore(num.left, e)
			list.InsertAfter(num.right, e)
			list.Remove(e)
			return true
		} else if !num.is_regular && num.level >= 4 && explode {
			//fmt.Printf("explode %s\n", num.printSnailfishNum())
			if !num.left.is_regular || !num.right.is_regular {
				panic(fmt.Sprintf("Left or right or both are not regular %s %s level=%d", num.left.printSnailfishNum(), num.right.printSnailfishNum(), num.level))
			}
			for back := e.Prev().Prev(); back != nil; back = back.Prev() {
				if back.Value.(*SnailfishNum).is_regular {
					back.Value.(*SnailfishNum).regular += num.left.regular
					//println("changed at left", back.Value.(*SnailfishNum).regular-num.left.regular, "to", back.Value.(*SnailfishNum).regular)
					break
				}
			}

			for forward := e.Next().Next(); forward != nil; forward = forward.Next() {
				if forward.Value.(*SnailfishNum).is_regular {
					forward.Value.(*SnailfishNum).regular += num.right.regular
					//println("changed at right", forward.Value.(*SnailfishNum).regular-num.right.regular, "to", forward.Value.(*SnailfishNum).regular)
					break
				}
			}
			num.is_regular = true
			num.regular = 0
			num.left = nil
			num.right = nil
			return true
		}
	}

	return false
}

func (num1 *SnailfishNum) addSnailfishNum(num2 *SnailfishNum) *SnailfishNum {
	num := SnailfishNum{}
	num.left = num1
	num.right = num2
	return &num
}

func (num *SnailfishNum) magnitude() int64 {
	if num.is_regular {
		return num.regular
	} else {
		return 3*num.left.magnitude() + 2*num.right.magnitude()
	}
}
