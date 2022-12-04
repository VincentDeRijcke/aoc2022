package main

import (
	"testing"
)

func Test_scoreGame(t *testing.T) {
	type args struct {
		input string
		want1 int
		want2 int
	}
	const fullExample = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	tests := []struct {
		name string
		args args
	}{
		{name: "Blank", args: args{input: "", want1: 0, want2: 0}},
		{name: "Ex1", args: args{input: "6-6,4-6", want1: 1, want2: 1}},
		{name: "Example Full", args: args{input: fullExample, want1: 2, want2: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			got1, got2 := resolve(args.input)
			if got1 != args.want1 {
				t.Errorf("Part1 %d; expected %d", got1, args.want1)
			} else if got2 != args.want2 {
				t.Errorf("Part2 %d; expected %d", got2, args.want2)
			}
		})
	}
}
