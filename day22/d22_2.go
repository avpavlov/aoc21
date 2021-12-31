package main

import (
	. "aoc2021/utils"
	"container/list"
	"strings"
)

func solve_d22_2(filename string) int64 {
	var activeCubes = list.New()
	var lines = ReadLines(filename)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		on := parts[0] == "on"

		parts = strings.Split(parts[1], ",")
		xs := ParseInt64s(strings.Split(parts[0][2:], ".."))
		ys := ParseInt64s(strings.Split(parts[1][2:], ".."))
		zs := ParseInt64s(strings.Split(parts[2][2:], ".."))

		cube := Cube{}
		cube.X1 = int(xs[0])
		cube.X2 = int(xs[1])
		cube.Y1 = int(ys[0])
		cube.Y2 = int(ys[1])
		cube.Z1 = int(zs[0])
		cube.Z2 = int(zs[1])
		if on {
			// on1+on2 -> on2-on1 to the rest of the list
			switchOnModificators := list.New()
			switchOnModificators.PushFront(&cube)

			for e1 := activeCubes.Front(); e1 != nil; e1 = e1.Next() {
				activeCube := e1.Value.(*Cube)
				newSwitchOnModificators := list.New()
				for e2 := switchOnModificators.Front(); e2 != nil; e2 = e2.Next() {
					switchOn := e2.Value.(*Cube)
					if !activeCube.Intersects(switchOn) {
						newSwitchOnModificators.PushFront(switchOn)
					} else if activeCube.Contains(switchOn) {
					} else {
						remainder := switchOn.Subtract(activeCube)
						for _, r := range remainder {
							newSwitchOnModificators.PushFront(r)
							//switchOnModificators.PushFront(r)
						}
					}
				}
				switchOnModificators = newSwitchOnModificators
			}
			activeCubes.PushBackList(switchOnModificators)
		} else {
			switchOff := &cube
			// on+off -> replace on with [on-off], apply off to the rest of the list
			newActiveCubes := list.New()
			for e1 := activeCubes.Front(); e1 != nil; e1 = e1.Next() {
				activeCube := e1.Value.(*Cube)
				if !activeCube.Intersects(switchOff) {
					newActiveCubes.PushFront(activeCube)
				} else if switchOff.Contains(activeCube) {
				} else {
					remainder := activeCube.Subtract(switchOff)
					for _, r := range remainder {
						newActiveCubes.PushFront(r)
					}
				}
			}
			activeCubes = newActiveCubes
		}
	}

	var result int
	for e := activeCubes.Front(); e != nil; e = e.Next() {
		cmd := e.Value.(*Cube)
		result += (cmd.X2 - cmd.X1 + 1) * (cmd.Y2 - cmd.Y1 + 1) * (cmd.Z2 - cmd.Z1 + 1)
	}
	return int64(result)
}
