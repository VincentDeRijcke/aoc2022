package main

import (
	"strings"
	"testing"
)

func Test_scoreGame(t *testing.T) {
	type args struct {
		strategy  string
		expected1 int
		expected2 int
	}
	const fullExample = `A Y
						 B X
						 C Z`
	tests := []struct {
		name string
		args args
	}{
		{name: "Blank", args: args{strategy: "", expected1: 0, expected2: 0}},
		{name: "Example Line 1", args: args{strategy: "A Y", expected1: 8, expected2: 4}},
		{name: "Example Line 2", args: args{strategy: "B X", expected1: 1, expected2: 1}},
		{name: "Example Line 3", args: args{strategy: "C Z", expected1: 6, expected2: 7}},
		{name: "Example Full", args: args{strategy: fullExample, expected1: 15, expected2: 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			actual1, actual2 := scoreGame(strings.NewReader(args.strategy))
			if actual1 != args.expected1 {
				t.Errorf("Score = %d; expected %d", actual1, args.expected1)
			} else if actual2 != args.expected2 {
				t.Errorf("Score = %d; expected %d", actual2, args.expected2)
			}
		})
	}
}
