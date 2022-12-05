package utils

import (
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
func Atois(slice []string) ([]int, error) {
	return SliceMapErr(slice, strconv.Atoi)
}

func Runes(strings []string) [][]rune {
	return SliceMap(strings, func(s string) []rune { return []rune(s) })
}

func RunesToString(runes [][]rune) string {
	return strings.Join(Strings(runes), "\n")
}

func StringToRunes(s string) [][]rune {
	return Runes(strings.Split(s, "\n"))
}

func Strings(runes [][]rune) []string {
	return SliceMap(runes, func(r []rune) string { return string(r) })
}

func TransposeString(s string) string {
	out, err := Transpose(StringToRunes(s))
	MaybePanic(err)
	return RunesToString(out)
}
