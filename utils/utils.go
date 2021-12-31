package utils

import (
	"fmt"
	"os"
	Sort "sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	check(err)
	return strings.Split(string(dat), "\n")
}

func SplitPair(text string, delim string) (string, string) {
	var pair = strings.Split(text, delim)
	return pair[0], pair[1]
}

func ParseInt64(s string) int64 {
	var m, err = strconv.ParseInt(s, 10, 32)
	check(err)
	return m
}

func ParseInt64s(strings []string) []int64 {
	var result = []int64{}
	for _, s := range strings {
		result = append(result, ParseInt64(s))
	}
	return result
}

func ReplaceAll(s string, old string, new string) string {
	for true {
		s2 := strings.ReplaceAll(s, old, new)
		if s2 == s {
			return s
		}
		s = s2
	}
	panic("cannot be here")
}

func BitsToNumber(s string) int {
	var n = 0
	for _, c := range s {
		n *= 2
		if c == '1' {
			n++
		}
	}
	return n
}

func MaxByLen(left []string, right []string) []string {
	if len(left) >= len(right) {
		return left
	} else {
		return right
	}
}

func MinByLen(left []string, right []string) []string {
	if len(left) >= len(right) {
		return right
	} else {
		return left
	}
}

func SplitByCondition(collection []string, condFn func(string) bool) ([]string, []string) {
	var truthy = []string{}
	var falsy = []string{}
	for _, line := range collection {
		if condFn(line) {
			truthy = append(truthy, line)
		} else {
			falsy = append(falsy, line)
		}
	}
	return truthy, falsy
}

func MinMax(n1, n2 int64) (int64, int64) {
	if n1 < n2 {
		return n1, n2
	} else {
		return n2, n1
	}
}

func SortInt64s(arr []int64) {
	Sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
}

func Sign(n int64) int64 {
	if n == 0 {
		return 0
	} else if n < 0 {
		return -1
	} else {
		return 1
	}
}

type Point struct {
	X int64
	Y int64
}

func ToPoint(s string) Point {
	var parts = strings.Split(s, ",")
	return Point{ParseInt64(parts[0]), ParseInt64(parts[1])}
}

func SortLetters(word string) string {
	s := []rune(word)
	Sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func Cut(message *string, count int64) string {
	result := (*message)[:count]
	*message = (*message)[count:]
	return result
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func CutDigits(message *string) string {
	for index := range *message {
		if !IsDigit((*message)[index]) {
			if index == 0 {
				return ""
			}
			result := (*message)[:index]
			*message = (*message)[index:]
			return result
		}
	}
	result := *message
	*message = ""
	return result
}

type Point3D struct {
	X, Y, Z int
}

type Point4D struct {
	X, Y, Z, W int
}

type Image map[Point]int

func (image *Image) CanvasBounds() (minX int64, minY int64, maxX int64, maxY int64) {
	for p := range *image {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}
	return
}

func (dots *Image) PrintCanvas(dot string, empty string) {
	var minX, minY, maxX, maxY = dots.CanvasBounds()

	var canvas = make([][]string, maxY+1-minY)
	for i := range canvas {
		canvas[i] = make([]string, maxX+1-minX)
	}

	for p, v := range *dots {
		if v == 0 {
			canvas[p.Y-minY][p.X-minX] = empty
		} else {
			canvas[p.Y-minY][p.X-minX] = dot
		}
	}

	for _, row := range canvas {
		var out = ""
		for _, cell := range row {
			if cell == dot {
				out += cell
			} else {
				out += empty
			}
		}
		fmt.Println(out)
	}
}

func MinInt64(v1 int64, v2 int64) int64 {
	if v1 <= v2 {
		return v1
	} else {
		return v2
	}
}
func MaxInt64(v1 int64, v2 int64) int64 {
	if v1 >= v2 {
		return v1
	} else {
		return v2
	}
}
func MinInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v1
	} else {
		return v2
	}
}
func MaxInt(v1 int, v2 int) int {
	if v1 >= v2 {
		return v1
	} else {
		return v2
	}
}
func AbsInt(v1 int) int {
	if v1 >= 0 {
		return v1
	} else {
		return -v1
	}
}
func AbsInt64(v1 int64) int64 {
	if v1 >= 0 {
		return v1
	} else {
		return -v1
	}
}

type Corner = Point3D
type Distance3D struct {
	X, Y, Z int
}

func (point *Point3D) OffsetBy(offset Distance3D) {
	point.X += offset.X
	point.Y += offset.Y
	point.Z += offset.Z
}

func (source Point3D) DistanceTo(target Point3D) Distance3D {
	return Distance3D{
		source.X - target.X,
		source.Y - target.Y,
		source.Z - target.Z,
	}
}

func (image *Image) ImproveImage(imageImprov *map[int]int, infinity int) *Image {
	nextImage := Image{}
	var minX, minY, maxX, maxY = image.CanvasBounds()
	var extra = int64(1)
	for row := minY - extra; row <= maxY+extra; row++ {
		for column := minX - extra; column <= maxX+extra; column++ {
			cell := Point{column, row}
			var mask int
			for i := int64(0); i < 9; i++ {
				x := cell.X + (i % 3) - 1
				y := cell.Y + ((i / 3) % 3) - 1
				var p = Point{x, y}
				mask *= 2
				if x < minX || x > maxX || y < minY || y > maxY {
					mask += infinity
				} else {
					mask += (*image)[p]
				}
			}
			nextImage[cell] = (*imageImprov)[mask]
		}
	}
	return &nextImage
}

type Cube struct {
	X1, X2, Y1, Y2, Z1, Z2 int
}

func (left *Cube) Contains(right *Cube) bool {
	return left.X1 <= right.X1 && right.X2 <= left.X2 && left.Y1 <= right.Y1 && right.Y2 <= left.Y2 && left.Z1 <= right.Z1 && right.Z2 <= left.Z2
}

func (left *Cube) Intersects(right *Cube) bool {
	return !(right.X2 < left.X1 || left.X2 < right.X1 || right.Y2 < left.Y1 || left.Y2 < right.Y1 || right.Z2 < left.Z1 || left.Z2 < right.Z1)
}

func (left *Cube) Subtract(right *Cube) []*Cube {
	cube1 := Cube{
		X1: left.X1,
		X2: right.X1 - 1,
		Y1: left.Y1,
		Y2: left.Y2,
		Z1: left.Z1,
		Z2: left.Z2,
	}
	cube2 := Cube{
		X1: MaxInt(right.X1, left.X1),
		X2: MinInt(right.X2, left.X2),
		Y1: left.Y1,
		Y2: left.Y2,
		Z1: left.Z1,
		Z2: left.Z2,
	}
	cube3 := Cube{
		X1: right.X2 + 1,
		X2: left.X2,
		Y1: left.Y1,
		Y2: left.Y2,
		Z1: left.Z1,
		Z2: left.Z2,
	}

	cube4 := Cube{
		X1: cube2.X1,
		X2: cube2.X2,
		Y1: cube2.Y1,
		Y2: right.Y1 - 1,
		Z1: cube2.Z1,
		Z2: cube2.Z2,
	}
	cube5 := Cube{
		X1: cube2.X1,
		X2: cube2.X2,
		Y1: MaxInt(right.Y1, left.Y1),
		Y2: MinInt(right.Y2, left.Y2),
		Z1: cube2.Z1,
		Z2: cube2.Z2,
	}
	cube6 := Cube{
		X1: cube2.X1,
		X2: cube2.X2,
		Y1: right.Y2 + 1,
		Y2: cube2.Y2,
		Z1: cube2.Z1,
		Z2: cube2.Z2,
	}

	cube7 := Cube{
		X1: cube5.X1,
		X2: cube5.X2,
		Y1: cube5.Y1,
		Y2: cube5.Y2,
		Z1: cube5.Z1,
		Z2: right.Z1 - 1,
	}
	cube9 := Cube{
		X1: cube5.X1,
		X2: cube5.X2,
		Y1: cube5.Y1,
		Y2: cube5.Y2,
		Z1: right.Z2 + 1,
		Z2: cube5.Z2,
	}

	result := []*Cube{}
	result = appendValidCube(cube1, result)
	result = appendValidCube(cube3, result)
	result = appendValidCube(cube4, result)
	result = appendValidCube(cube6, result)
	result = appendValidCube(cube7, result)
	result = appendValidCube(cube9, result)
	return result
}

func appendValidCube(cube Cube, array []*Cube) []*Cube {
	if cube.X1 <= cube.X2 && cube.Y1 <= cube.Y2 && cube.Z1 <= cube.Z2 {
		return append(array, &cube)
	} else {
		return array
	}
}
