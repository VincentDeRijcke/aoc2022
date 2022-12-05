package main

import (
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Split a splice of string using the same separator.
// Return the merge result of all the splits.
func splitSlice(slice []string, sep string) []string {
	if len(slice) == 0 {
		return slice
	}

	var res []string
	for _, s := range slice {
		res = append(res, strings.Split(s, sep)...)
	}

	return res
}

// Split a string by all the separators
func splits(s string, seps ...string) []string {
	res := []string{s}
	for _, sep := range seps {
		res = splitSlice(res, sep)
	}

	return res
}

// Convert a slice of string into a slice of int
func atois(slice []string) ([]int, error) {
	return sliceMapErr(slice, strconv.Atoi)
}
func sliceMap[I, O any](slice []I, f func(I) O) (res []O) {
	res, _ = sliceMapErr(slice, func(i I) (O, error) {
		return f(i), nil
	})

	return
}

func sliceMapErr[I, O any](slice []I, f func(I) (O, error)) (res []O, err error) {
	if slice == nil {
		return nil, nil
	}
	if len(slice) == 0 {
		return []O{}, nil
	}
	res = make([]O, len(slice))
	for i, v := range slice {
		res[i], err = f(v)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
