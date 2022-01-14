package main

import (
	. "aoc2021/utils"
	"fmt"
	"log"
	"strings"
)

func solve_d24_1(filename string) int64 {
	return solve_d24(filename, MaxInt64)
}

func solve_d24_2(filename string) int64 {
	return solve_d24(filename, MinInt64)
}

type value_t int64

type State struct {
	x, y, z, w value_t
}

type Register struct {
	get func(state *State) value_t
	set func(state *State, val value_t)
}

var registers = map[string]Register{
	"x": {get: func(state *State) value_t { return state.x }, set: func(state *State, val value_t) { state.x = val }},
	"y": {get: func(state *State) value_t { return state.y }, set: func(state *State, val value_t) { state.y = val }},
	"z": {get: func(state *State) value_t { return state.z }, set: func(state *State, val value_t) { state.z = val }},
	"w": {get: func(state *State) value_t { return state.w }, set: func(state *State, val value_t) { state.w = val }},
}

type Command func(v1, v2 value_t) value_t

var commands = map[string]Command{
	"mul": func(v1, v2 value_t) value_t {
		res := v1 * v2
		if v2 != 0 && (res/v2) != v1 {
			panic(fmt.Sprintf("res=%d op1=%d op2=%d", res, v1, v2))
		}
		return res
	},
	"add": func(v1, v2 value_t) value_t {
		res := v1 + v2
		if (res - v2) != v1 {
			panic(fmt.Sprintf("res=%d op1=%d op2=%d", res, v1, v2))
		}
		return res
	},
	"div": func(v1, v2 value_t) value_t { return v1 / v2 },
	"mod": func(v1, v2 value_t) value_t { return v1 % v2 },
	"set": func(v1, v2 value_t) value_t { return v2 },
	"eql": func(v1, v2 value_t) value_t {
		if v1 == v2 {
			return 1
		} else {
			return 0
		}
	},
	"neq": func(v1, v2 value_t) value_t {
		if v1 != v2 {
			return 1
		} else {
			return 0
		}
	},
}

func parseOperands(reg string, regOrConst string) (register1 *Register, useRegister2 bool, register2 *Register, value2 value_t) {
	operand1, isRegister := registers[reg]
	if !isRegister {
		panic(reg)
	}
	if operand2, isRegister := registers[regOrConst]; isRegister {
		return &operand1, true, &operand2, 0
	} else {
		return &operand1, false, nil, value_t(ParseInt64(regOrConst))
	}
}

func solve_d24(filename string, reduceFn func(v1, v2 int64) int64) int64 {
	var lines = ReadLines(filename)
	tokenizedLines := make([][]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		tokenizedLines = append(tokenizedLines, parts)
	}
	lastToken := len(tokenizedLines) - 1
	states := []State{{0, 0, 0, 0}}
	statesSerials := []int64{0}
	for i := 0; i < len(tokenizedLines); i++ {
		parts := tokenizedLines[i]
		if i != lastToken {
			nextParts := tokenizedLines[i+1]
			if parts[1] == nextParts[1] {
				if parts[0] == "mul" && parts[2] == "0" && nextParts[0] == "add" {
					parts = nextParts
					parts[0] = "set"
					i++
				} else if parts[0] == "eql" && nextParts[0] == "eql" && nextParts[2] == "0" {
					parts[0] = "neq"
					i++
				}
			}
		}
		log.Printf("%d of %d %v, states=%d", i+1, len(lines), parts, len(states))

		if parts[0] == "add" && parts[2] == "0" {
			continue
		} else if parts[0] == "inp" {
			newStates := reduceStates(&states, &statesSerials, reduceFn)
			states = make([]State, 0, len(newStates)*9)
			statesSerials = make([]int64, 0, len(newStates)*9)

			register1 := registers[parts[1]]
			for state, serial := range newStates {
				for digit := value_t(1); digit <= 9; digit++ {
					register1.set(&state, digit)
					states = append(states, state)
					statesSerials = append(statesSerials, serial*10+int64(digit))
				}
			}
		} else {
			command := commands[parts[0]]
			register1, useRegister2, register2, value2 := parseOperands(parts[1], parts[2])
			for index, state := range states {
				v1 := register1.get(&state)
				var v2 value_t
				if useRegister2 {
					v2 = register2.get(&state)
				} else {
					v2 = value2
				}
				register1.set(&state, command(v1, v2))
				states[index] = state
			}
		}
	}

	var initialized bool
	var result int64
	for index, state := range states {
		serial := statesSerials[index]
		if state.z == 0 && (!initialized || result != reduceFn(result, serial)) {
			result = serial
			initialized = true
		}
	}

	return result
}

func reduceStates(states *[]State, statesSerials *[]int64, reduceFn func(v1, v2 int64) int64) map[State]int64 {
	newStates := make(map[State]int64, len(*states))
	for i, state := range *states {
		serial := (*statesSerials)[i]
		if currentSerial, exist := newStates[state]; !exist || currentSerial != reduceFn(currentSerial, serial) {
			newStates[state] = serial
		}
	}
	return newStates
}
