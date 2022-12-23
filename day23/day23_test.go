package main

import (
	"testing"
)

var small = `..##.
..#..
.....
..##.
`
var example = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`

func Test_resolve(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want1 int
		want2 int
	}{
		{name: "Small", args: args{input: small}, want1: 25, want2: 0},
		{name: "Example", args: args{input: example}, want1: 110, want2: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := resolve(tt.args.input)
			if got1 != tt.want1 {
				t.Errorf("resolve() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("resolve() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
