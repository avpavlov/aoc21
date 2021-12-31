package main

import (
	. "aoc2021/utils"
	"fmt"
	"strings"
)

func solve_d19_2(filename string) int64 {
	var lines = ReadLines(filename)
	var scanners []*Scanner
	var inputScanner *Scanner
	for _, line := range lines {
		if line == "" {
			continue
		} else if line[:3] == "---" {
			inputScanner = &Scanner{
				id:             strings.Split(line, " ")[2],
				center:         Point3D{0, 0, 0},
				beacons:        make(map[string]Beacon, 30),
				groupsByCorner: map[Corner][]*BeaconCornerGroup{},
				rotations:      map[string]*Scanner{},
			}
			scanners = append(scanners, inputScanner)
			continue
		} else {
			coords := ParseInt64s(strings.Split(line, ","))
			beacon := Beacon{
				id:    fmt.Sprintf("%s-b%d", inputScanner.id, len(inputScanner.beacons)),
				point: Point3D{int(coords[0]), int(coords[1]), int(coords[2])},
			}
			inputScanner.beacons[beacon.id] = beacon
		}
	}

	for _, scanner := range scanners {
		scanner.findGroups()
	}

	matched := append([]*Scanner{}, scanners[0])
	unmatched := scanners[1:]
loop:
	for len(unmatched) > 0 {
		for i, candidate := range unmatched {
			for rotationId := range rotations {
				rotatedCandidate := candidate.rotate(rotationId)
				for _, scanner := range matched {
					if scanner.match(rotatedCandidate) {
						matched = append(matched, rotatedCandidate)
						unmatched = append(unmatched[:i], unmatched[i+1:]...)
						continue loop
					}
				}
			}
		}
		panic("Cannot match")
	}

	var result int
	for _, scanner1 := range matched {
		for _, scanner2 := range matched {
			d := scanner1.center.DistanceTo(scanner2.center)
			if r := AbsInt(d.X) + AbsInt(d.Y) + AbsInt(d.Z); r > result {
				result = r
			}
		}
	}

	return int64(result)
}
