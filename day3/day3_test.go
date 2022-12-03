package main

import (
	"strings"
	"testing"
)

func Test_scoreGame(t *testing.T) {
	type args struct {
		input     string
		expected1 int
		expected2 int
	}
	const fullExample = `vJrwpWtwJgWrhcsFMMfFFhFp
						 jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
                         PmmdzqPrVvPwwTWBwg
						 wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
						 ttgJtRGJQctTZtZT
                         CrZsJsPPZsGzwwsLwLmpwMDw`
	const group1 = `vJrwpWtwJgWrhcsFMMfFFhFp
						 jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
                         PmmdzqPrVvPwwTWBwg`
	const group2 = `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
					ttgJtRGJQctTZtZT
					CrZsJsPPZsGzwwsLwLmpwMDw`
	tests := []struct {
		name string
		args args
	}{
		{name: "Blank", args: args{input: "", expected1: 0, expected2: 0}},
		{name: "Example Line 1", args: args{input: "vJrwpWtwJgWrhcsFMMfFFhFp", expected1: 16, expected2: 0}},
		{name: "Example Group 1", args: args{input: group1, expected1: 16 + 38 + 42, expected2: 18}},
		{name: "Example Group 2", args: args{input: group2, expected1: 22 + 20 + 19, expected2: 52}},
		{name: "Example Full", args: args{input: fullExample, expected1: 157, expected2: 70}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			actual1, actual2 := sumErrorsPriorities(strings.NewReader(args.input))
			if actual1 != args.expected1 {
				t.Errorf("Score = %d; expected %d", actual1, args.expected1)
			} else if actual2 != args.expected2 {
				t.Errorf("Score = %d; expected %d", actual2, args.expected2)
			}
		})
	}
}
