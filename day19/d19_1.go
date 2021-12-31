package main

import (
	. "aoc2021/utils"
	"fmt"
	Sort "sort"
	"strings"
)

func solve_d19_1(filename string) int64 {
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

	allPoints := map[Point3D]bool{}
	for _, scanner := range matched {
		for _, beacon := range scanner.beacons {
			allPoints[beacon.point] = true
		}
	}
	return int64(len(allPoints))
}

type Beacon struct {
	id    string
	point Point3D
}

var corners = []Corner{
	{-1_000, -1_000, -1_000},
	{-1_000, -1_000, 1_000},
	{-1_000, 1_000, -1_000},
	{-1_000, 1_000, 1_000},
	{1_000, -1_000, -1_000},
	{1_000, -1_000, 1_000},
	{1_000, 1_000, -1_000},
	{1_000, 1_000, 1_000},
}

type Rotation func(Point3D) Point3D

var rotations = map[string]Rotation{
	"flt":   func(p Point3D) Point3D { return p },
	"f-tl":  func(p Point3D) Point3D { return Point3D{p.X, -p.Z, p.Y} },
	"f-l-t": func(p Point3D) Point3D { return Point3D{p.X, -p.Y, -p.Z} },
	"ft-l":  func(p Point3D) Point3D { return Point3D{p.X, p.Z, -p.Y} },

	"-f-lt":  func(p Point3D) Point3D { return Point3D{-p.X, -p.Y, p.Z} },
	"-ftl":   func(p Point3D) Point3D { return Point3D{-p.X, p.Z, p.Y} },
	"-fl-t":  func(p Point3D) Point3D { return Point3D{-p.X, p.Y, -p.Z} },
	"-f-t-l": func(p Point3D) Point3D { return Point3D{-p.X, -p.Z, -p.Y} },

	"-lft":  func(p Point3D) Point3D { return Point3D{-p.Y, p.X, p.Z} },
	"tfl":   func(p Point3D) Point3D { return Point3D{p.Z, p.X, p.Y} },
	"lf-t":  func(p Point3D) Point3D { return Point3D{p.Y, p.X, -p.Z} },
	"-tf-l": func(p Point3D) Point3D { return Point3D{-p.Z, p.X, -p.Y} },

	"l-ft":   func(p Point3D) Point3D { return Point3D{p.Y, -p.X, p.Z} },
	"-t-fl":  func(p Point3D) Point3D { return Point3D{-p.Z, -p.X, p.Y} },
	"-l-f-t": func(p Point3D) Point3D { return Point3D{-p.Y, -p.X, -p.Z} },
	"t-f-l":  func(p Point3D) Point3D { return Point3D{p.Z, -p.X, -p.Y} },

	"t-lf":  func(p Point3D) Point3D { return Point3D{p.Z, -p.Y, p.X} },
	"ltf":   func(p Point3D) Point3D { return Point3D{p.Y, p.Z, p.X} },
	"-tlf":  func(p Point3D) Point3D { return Point3D{-p.Z, p.Y, p.X} },
	"-l-tf": func(p Point3D) Point3D { return Point3D{-p.Y, -p.Z, p.X} },

	"l-t-f":  func(p Point3D) Point3D { return Point3D{p.Y, -p.Z, -p.X} },
	"-t-l-f": func(p Point3D) Point3D { return Point3D{-p.Z, -p.Y, -p.X} },
	"-lt-f":  func(p Point3D) Point3D { return Point3D{-p.Y, p.Z, -p.X} },
	"tl-f":   func(p Point3D) Point3D { return Point3D{p.Z, p.Y, -p.X} },
}

type Scanner struct {
	id             string
	center         Point3D
	beacons        map[string]Beacon
	groupsByCorner map[Corner][]*BeaconCornerGroup
	rotations      map[string]*Scanner
}

type BeaconCornerGroup struct {
	id      string
	beacons []*Beacon
}

func (scanner *Scanner) findGroups() {
	xs := map[int]bool{}
	ys := map[int]bool{}
	zs := map[int]bool{}
	for _, beacon := range scanner.beacons {
		xs[beacon.point.X] = true
		ys[beacon.point.Y] = true
		zs[beacon.point.Z] = true
	}

	for _, corner := range corners {
		var minSize = len(scanner.beacons) + 1
		var minSizeGroups map[string]*BeaconCornerGroup
		for x := range xs {
			for y := range ys {
				for z := range zs {
					cube := Cube{
						X1: MinInt(corner.X, x),
						Y1: MinInt(corner.Y, y),
						Z1: MinInt(corner.Z, z),
						X2: MaxInt(corner.X, x),
						Y2: MaxInt(corner.Y, y),
						Z2: MaxInt(corner.Z, z),
					}
					beaconsInCube := findBeaconsInCube(&scanner.beacons, cube)
					if gs := len(beaconsInCube); gs < 12 || gs > minSize {
						continue
					} else if gs < minSize {
						minSize = len(beaconsInCube)
						minSizeGroups = map[string]*BeaconCornerGroup{}
					}
					group := BeaconCornerGroup{beacons: beaconsInCube}
					sortBeacons(&group)
					minSizeGroups[group.id] = &group
				}
			}
		}
		groups := []*BeaconCornerGroup{}
		for _, group := range minSizeGroups {
			groups = append(groups, group)
		}
		scanner.groupsByCorner[corner] = groups
	}
}

func findBeaconsInCube(beacons *map[string]Beacon, cube Cube) []*Beacon {
	var result []*Beacon
	for _, beacon := range *beacons {
		if cube.X1 <= beacon.point.X && beacon.point.X <= cube.X2 && cube.Y1 <= beacon.point.Y && beacon.point.Y <= cube.Y2 && cube.Z1 <= beacon.point.Z && beacon.point.Z <= cube.Z2 {
			result = append(result, &Beacon{
				id:    beacon.id,
				point: beacon.point,
			})
		}
	}
	return result
}

func sortBeacons(group *BeaconCornerGroup) {
	Sort.Slice(group.beacons, func(i, j int) bool {
		p1 := group.beacons[i].point
		p2 := group.beacons[j].point
		return p1.X < p2.X || (p1.X == p2.X && p1.Y < p2.Y) || (p1.X == p2.X && p1.Y == p2.Y && p1.Z < p2.Z)
	})
	id := ""
	for _, beacon := range group.beacons {
		id += "+" + beacon.id
	}
	group.id = id
}

func (scanner *Scanner) rotate(rotationId string) *Scanner {
	if rotatedScanner, found := scanner.rotations[rotationId]; found {
		return rotatedScanner
	} else if rotationId == "flt" {
		return scanner
	}
	var rotation = rotations[rotationId]
	var rotatedScanner = Scanner{
		id:             fmt.Sprintf("%s_r%s", scanner.id, rotationId),
		beacons:        make(map[string]Beacon),
		groupsByCorner: map[Corner][]*BeaconCornerGroup{},
	}
	// copy & rotate beacons
	for id, beacon := range scanner.beacons {
		rotatedScanner.beacons[id] = Beacon{
			id:    id,
			point: rotation(beacon.point),
		}
	}
	// copy & rotate beacon groupsByCorner (incl. corners)
	for corner, groups := range scanner.groupsByCorner {
		rotatedGroups := make([]*BeaconCornerGroup, 0, len(groups))
		for _, group := range groups {
			rotatedGroup := BeaconCornerGroup{
				beacons: make([]*Beacon, 0, len(group.beacons)),
			}
			for _, beacon := range group.beacons {
				rotatedGroup.beacons = append(rotatedGroup.beacons, &Beacon{
					id:    beacon.id,
					point: rotation(beacon.point),
				})
			}
			sortBeacons(&rotatedGroup)
			rotatedGroups = append(rotatedGroups, &rotatedGroup)
		}
		rotatedScanner.groupsByCorner[rotation(corner)] = rotatedGroups
	}
	scanner.rotations[rotationId] = &rotatedScanner
	return &rotatedScanner
}

func (scanner *Scanner) match(candidate *Scanner) bool {
	for _, corner := range corners {
		oppositeCorner := Corner{
			X: -corner.X,
			Y: -corner.Y,
			Z: -corner.Z,
		}
		for _, group := range scanner.groupsByCorner[corner] {
		candidateGroupLoop:
			for _, candidateGroup := range candidate.groupsByCorner[oppositeCorner] {
				if len(group.beacons) != len(candidateGroup.beacons) {
					continue
				}
				first := group.beacons[0].point
				candidateFirst := candidateGroup.beacons[0].point
				for i := 1; i < len(group.beacons); i++ {
					next := group.beacons[i].point
					candidateNext := candidateGroup.beacons[i].point
					if first.DistanceTo(next) != candidateFirst.DistanceTo(candidateNext) {
						continue candidateGroupLoop
					}
				}
				// hooray!
				candidate.fixPosition(first.DistanceTo(candidateFirst))
				return true
			}
		}
	}
	return false
}

func (scanner *Scanner) fixPosition(offset Distance3D) {
	scanner.center.OffsetBy(offset)
	// update all beacons
	for id, beacon := range scanner.beacons {
		beacon.point.OffsetBy(offset)
		scanner.beacons[id] = beacon
	}
	// update all beacon groupsByCorner (EXCL. corners)
	for _, groups := range scanner.groupsByCorner {
		for _, group := range groups {
			for _, beacon := range group.beacons {
				beacon.point.OffsetBy(offset)
			}
		}
	}
	scanner.rotations = nil
}
