package utils

import (
	"errors"
	"strings"
)

// SplitSlice splits a splice of strings using the same separator.
// Return the merge result of all the Splits.
func SplitSlice(slice []string, sep string) []string {
	if len(slice) == 0 {
		return slice
	}

	var res []string
	for _, s := range slice {
		res = append(res, strings.Split(s, sep)...)
	}

	return res
}

// Splits splits a string by all the separators
func Splits(s string, seps ...string) []string {
	res := []string{s}
	for _, sep := range seps {
		res = SplitSlice(res, sep)
	}

	return res
}

// SliceMap maps all the element of a slice into another slice
func SliceMap[I, O any](slice []I, f func(I) O) (res []O) {
	res, _ = SliceMapErr(slice, func(i I) (O, error) {
		return f(i), nil
	})

	return
}

// SliceMapErr is a more general version of SliceMap with an error returning mapping function
func SliceMapErr[I, O any](slice []I, f func(I) (O, error)) (res []O, err error) {
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

// SliceFilter filters a slice of all the elements not matching the predicate
func SliceFilter[T any](slice []T, predicate func(T) bool) (res []T) {
	res, _ = SliceFilterWithIndexes(slice, predicate)

	return
}

// SliceFilterWithIndexes filters a slice of all the elements not matching the predicate.
// Also returns the indexes of all the indices not filtered in the original slice
func SliceFilterWithIndexes[T any](slice []T, predicate func(T) bool) (res []T, indexes []int) {
	if slice == nil {
		return nil, nil
	}
	if len(slice) == 0 {
		return []T{}, []int{}
	}
	res = make([]T, 0, len(slice))
	indexes = make([]int, 0, len(slice))

	for i, v := range slice {
		if predicate(v) {
			res = append(res, v)
			indexes = append(indexes, i)
		}
	}

	return res, indexes
}

// Reverse reverses a slice. A new slice is created.
func Reverse[T any](in []T) []T {
	if in == nil {
		return nil
	}
	out := make([]T, len(in))

	last := len(in) - 1
	for i := 0; i < len(in); i++ {
		out[last-i] = in[i]
	}

	return out
}

// ReverseInPlace in place reverse of a slice
func ReverseInPlace[T any](in []T) []T {
	i, l := 0, len(in)-1
	if l > 0 {
		for i <= l {
			tmp := in[i]
			in[i] = in[l]
			in[l] = tmp

			i++
			l--
		}
	}

	return in
}

// Transpose lines to columns
// Each line should have same size
// From
//
//	N,C,X
//	Z,M,P
//	1,2,3
//
// To
//
//	1,Z,N
//	2,M,C
//	3,P,X
func Transpose[T any](in [][]T) (out [][]T, err error) {
	if in == nil {
		return nil, nil
	}
	if len(in) == 0 {
		return in, nil
	}

	out = make([][]T, len(in[0]))
	in = Reverse(in)
	for _, line := range in {
		if len(out) != len(line) {
			return nil, errors.New("not all line have same size")
		}
		for i, v := range line {
			out[i] = append(out[i], v)
		}
	}

	return out, nil
}
func Contains[T comparable](slice []T, want T) bool {
	for _, got := range slice {
		if got == want {
			return true
		}
	}
	return false
}
