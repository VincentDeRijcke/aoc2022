package main

import (
	"testing"
)

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
		{name: "EdgeCase", args: args{input: ""}, want1: -1, want2: -1},
		{name: "EdgeCase", args: args{input: "111"}, want1: -1, want2: -1},
		{name: "EdgeCase", args: args{input: "1234"}, want1: 4, want2: -1},
		{name: "EdgeCase", args: args{input: "12345"}, want1: 4, want2: -1},
		{name: "EdgeCase", args: args{input: "1234567890abc"}, want1: 4, want2: -1},
		{name: "EdgeCase", args: args{input: "1234567890abcd"}, want1: 4, want2: 14},
		{name: "Example", args: args{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, want1: 7, want2: 19},
		{name: "Example", args: args{input: "bvwbjplbgvbhsrlpgdmjqwftvncz"}, want1: 5, want2: 23},
		{name: "Example", args: args{input: "nppdvjthqldpwncqszvftbrmjlhg"}, want1: 6, want2: 23},
		{name: "Example", args: args{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, want1: 10, want2: 29},
		{name: "Example", args: args{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, want1: 11, want2: 26},
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

func Test_isMarker(t *testing.T) {
	type args struct {
		s []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Example", args: args{s: []rune("nppd")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMarker(tt.args.s); got != tt.want {
				t.Errorf("isMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
