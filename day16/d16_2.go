package main

import (
	"aoc2021/utils"
	"strings"
)

func solve_d16_2(filename string) int64 {
	var mapping = map[string]string{"0": "0000", "1": "0001", "2": "0010", "3": "0011", "4": "0100", "5": "0101", "6": "0110", "7": "0111", "8": "1000", "9": "1001", "A": "1010", "B": "1011", "C": "1100", "D": "1101", "E": "1110", "F": "1111"}
	var lines = utils.ReadLines(filename)

	var binaryMessage = ""
	for _, letter := range strings.Split(lines[0], "") {
		binaryMessage += mapping[letter]
	}

	var packets = parse(&binaryMessage)

	verifyRemainder(binaryMessage)

	var result int64
	for _, packet := range *packets {
		result += packet.value()
	}
	return result
}

func (packet *Packet) value() int64 {
	if packet.typeId == 4 {
		return packet.literal
	}
	result := packet.subpackets[0].value()
	switch packet.typeId {
	case 0:
		for _, packet := range packet.subpackets[1:] {
			result += packet.value()
		}
	case 1:
		for _, packet := range packet.subpackets[1:] {
			result *= packet.value()
		}
	case 2:
		for _, packet := range packet.subpackets[1:] {
			if v := packet.value(); v < result {
				result = v
			}
		}
	case 3:
		for _, packet := range packet.subpackets[1:] {
			if v := packet.value(); v > result {
				result = v
			}
		}
	case 5:
		if result > packet.subpackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	case 6:
		if result < packet.subpackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	case 7:
		if result == packet.subpackets[1].value() {
			result = 1
		} else {
			result = 0
		}
	}
	return result
}
