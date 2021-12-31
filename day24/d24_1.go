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

const registerX = 0
const registerY = 1
const registerZ = 2
const registerW = 3

const operationAdd = 4
const operationMul = 5
const operationDiv = 6
const operationMod = 7
const operationEql = 8
const operationInp = 9

var registers = map[string]byte{"x": registerX, "y": registerY, "z": registerZ, "w": registerW}
var commands = map[string]byte{"inp": operationInp, "mul": operationMul, "add": operationAdd, "div": operationDiv, "mod": operationMod, "eql": operationEql}

func get(state *State, register byte) value_t {
	switch register {
	case registerX:
		return state.x
	case registerY:
		return state.y
	case registerZ:
		return state.z
	case registerW:
		return state.w
	default:
		panic("Unknown register")
	}
}

func set(state *State, register byte, value value_t) {
	switch register {
	case registerX:
		state.x = value
	case registerY:
		state.y = value
	case registerZ:
		state.z = value
	case registerW:
		state.w = value
	default:
		panic("Unknown register")
	}
}

func calc(operation byte, v1, v2 value_t) value_t {
	switch operation {
	case operationAdd:
		res := v1 + v2
		if (res - v2) != v1 {
			panic(fmt.Sprintf("res=%d op1=%d op2=%d", res, v1, v2))
		}
		return res
	case operationMul:
		res := v1 * v2
		if v2 != 0 && (res/v2) != v1 {
			panic(fmt.Sprintf("res=%d op1=%d op2=%d", res, v1, v2))
		}
		return res
	case operationDiv:
		return v1 / v2
	case operationMod:
		return v1 % v2
	case operationEql:
		if v1 == v2 {
			return 1
		} else {
			return 0
		}
	default:
		panic("Unknown command")
	}
}

func parseOperands(reg string, regOrConst string) (register1 byte, useRegister2 bool, register2 byte, value2 value_t) {
	operand1, isRegister := registers[reg]
	if !isRegister {
		panic(reg)
	}
	if operand2, isRegister := registers[regOrConst]; isRegister {
		return operand1, true, operand2, 0
	} else {
		return operand1, false, 0, value_t(ParseInt64(regOrConst))
	}
}

func solve_d24(filename string, reduceFn func(v1, v2 int64) int64) int64 {
	states := []State{{0, 0, 0, 0}}
	statesSerials := []int64{0}
	var lines = ReadLines(filename)
	for i, line := range lines {
		parts := strings.Split(line, " ")
		log.Printf("%d of %d %s, states=%d", i+1, len(lines), line, len(states))

		command := commands[parts[0]]

		if command == operationInp {
			newStates := reduceStates(&states, &statesSerials, reduceFn)
			states = make([]State, 0, len(newStates)*9)
			statesSerials = make([]int64, 0, len(newStates)*9)

			register1 := registers[parts[1]]
			for state, serial := range newStates {
				for digit := value_t(1); digit <= 9; digit++ {
					set(&state, register1, digit)
					states = append(states, state)
					statesSerials = append(statesSerials, serial*10+int64(digit))
				}
			}
		} else {
			register1, useRegister2, register2, value2 := parseOperands(parts[1], parts[2])
			for index, state := range states {
				v1 := get(&state, register1)
				var v2 value_t
				if useRegister2 {
					v2 = get(&state, register2)
				} else {
					v2 = value2
				}
				set(&state, register1, calc(command, v1, v2))
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
