package main

import (
	"testing"
)

var example = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
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
		{name: "Example", args: args{input: example}, want1: 18, want2: 54},
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

func TestBlizzard_pos(t *testing.T) {
	tests := []struct {
		name     string
		blizzard Blizzard
		min      int
		want     int
	}{
		{name: "#>....#-0", blizzard: Blizzard{1, 5, 1}, min: 0, want: 1},
		{name: "#>x...#-1", blizzard: Blizzard{1, 5, 1}, min: 1, want: 2},
		{name: "#>...x#-4", blizzard: Blizzard{1, 5, 1}, min: 4, want: 5},
		{name: "#>....#-5", blizzard: Blizzard{1, 5, 1}, min: 5, want: 1},
		{name: "#>x...#-6", blizzard: Blizzard{1, 5, 1}, min: 6, want: 2},
		{name: "#x...>#-1", blizzard: Blizzard{5, 5, 1}, min: 4, want: 4},
		{name: "#....<#-0", blizzard: Blizzard{5, 5, -1}, min: 0, want: 5},
		{name: "#x...<#-0", blizzard: Blizzard{5, 5, -1}, min: 4, want: 1},
		{name: "#....<#-5", blizzard: Blizzard{5, 5, -1}, min: 5, want: 5},
		{name: "#x...<#-9", blizzard: Blizzard{5, 5, -1}, min: 9, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.blizzard.pos(tt.min); got != tt.want {
				t.Errorf("pos() = %v, want %v", got, tt.want)
			}
		})
	}
}
