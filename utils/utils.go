package utils

import (
	"os"
	"strconv"
	"strings"
)

// MaybePanic panics if err is not nil
func MaybePanic(e error) {
	if e != nil {
		panic(e)
	}
}

// Atois converts a slice of string into a slice of int
func Atois(slice []string) (ints []int) {
	ints, err := SliceMapErr(slice, strconv.Atoi)
	MaybePanic(err)
	return
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	MaybePanic(err)

	return i
}

func Runes(strings []string) [][]rune {
	return SliceMap(strings, func(s string) []rune { return []rune(s) })
}

func Strings(runes [][]rune) []string {
	return SliceMap(runes, func(r []rune) string { return string(r) })
}

func RunesToString(runes [][]rune) string {
	return strings.Join(Strings(runes), "\n")
}

func StringToRunes(s string) [][]rune {
	return Runes(strings.Split(s, "\n"))
}

func RunesTrimSpace(runes []rune) []rune {
	return []rune(strings.TrimSpace(string(runes)))
}

// RotateString rotates from
//
//	...x...
//	..xxx..
//	.xxxxx.
//
// To
//
//	...
//	x..
//	xx.
//	xxx
//	xx.
//	x..
//	...
func RotateString(s string) string {
	return RunesToString(Transpose(Reverse(StringToRunes(s))))
}

func IsNotEmpty(s string) bool {
	return len(s) > 0
}

func Max(x int, ys ...int) int {
	max := x
	for _, y := range ys {
		if y > max {
			max = y
		}
	}
	return max
}

func Min(x int, ys ...int) int {
	min := x
	for _, y := range ys {
		if y < min {
			min = y
		}
	}
	return min
}

func Diff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func IsInt(a any) bool {
	switch a.(type) {
	case int:
		return true
	default:
		return false
	}
}

func StringToFile(file, content string) {
	f, err := os.Create(file)
	MaybePanic(err)
	defer f.Close()

	_, err = f.WriteString(content)
	MaybePanic(err)
}
