package main

import (
	"aoc2021/utils"
	"fmt"
	"strings"
)

func solve_d16_1(filename string) int64 {
	var mapping = map[string]string{"0": "0000", "1": "0001", "2": "0010", "3": "0011", "4": "0100", "5": "0101", "6": "0110", "7": "0111", "8": "1000", "9": "1001", "A": "1010", "B": "1011", "C": "1100", "D": "1101", "E": "1110", "F": "1111"}
	var lines = utils.ReadLines(filename)

	var binaryMessage = ""
	for _, letter := range strings.Split(lines[0], "") {
		binaryMessage += mapping[letter]
	}

	var packets = parse(&binaryMessage)

	verifyRemainder(binaryMessage)

	return packets.sumVersions()
}

type Packet struct {
	version    int64
	typeId     int64
	literal    int64
	subpackets Packets
}

type Packets []*Packet

func parse(message *string) *Packets {
	var packets Packets
	version := parseBinary(utils.Cut(message, 3))
	typeId := parseBinary(utils.Cut(message, 3))
	if typeId == 4 {
		var literalStr = ""
		for true {
			var groupStr = utils.Cut(message, 5)
			literalStr += groupStr[1:]
			if groupStr[0:1] == "0" {
				break
			}
		}
		packets = append(packets, &Packet{
			version: version,
			typeId:  typeId,
			literal: parseBinary(literalStr),
		})
	} else {
		packet := Packet{
			version: version,
			typeId:  typeId,
		}
		lengthTypeId := parseBinary(utils.Cut(message, 1))
		if lengthTypeId == 0 {
			subpacketsLength := parseBinary(utils.Cut(message, 15))
			for subpackets := utils.Cut(message, subpacketsLength); len(subpackets) > 0; {
				packet.subpackets = append(packet.subpackets, *parse(&subpackets)...)
			}
		} else {
			for subpacketsCount := parseBinary(utils.Cut(message, 11)); subpacketsCount > 0; subpacketsCount-- {
				packet.subpackets = append(packet.subpackets, *parse(message)...)
			}
		}
		packets = append(packets, &packet)
	}

	return &packets
}

func (packets *Packets) sumVersions() int64 {
	var result int64
	for _, packet := range *packets {
		result += packet.version
		result += packet.subpackets.sumVersions()
	}
	return result
}

func parseBinary(binaryText string) int64 {
	var result int64
	for _, bit := range strings.Split(binaryText, "") {
		result *= 2
		if bit == "1" {
			result++
		}
	}
	return result
}

func verifyRemainder(binaryMessage string) {
	if parseBinary(binaryMessage) != 0 {
		fmt.Printf("Non-zero remainder %s\n", binaryMessage)
	}
}
