package utils

import (
	"reflect"
	"testing"
)

var grid = StringToRunes("" +
	".1.\n" +
	"121\n" +
	".1.\n" +
	".0.")

func TestGrid_Move(t *testing.T) {
	type args struct {
		c int
		r int
	}
	type testCase[I any] struct {
		name string
		args args
		want I
	}
	tests := []testCase[rune]{
		{name: "normal", args: args{c: 0, r: 0}, want: '.'},
		{name: "normal", args: args{c: 1, r: 1}, want: '2'},
		{name: "normal", args: args{c: 2, r: 2}, want: '.'},
		{name: "normal", args: args{c: 1, r: 3}, want: '0'},
		{name: "normal", args: args{c: 2, r: 3}, want: '.'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g = NewGrid(grid)
			if got := g.Move(tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExplorer_Up(t *testing.T) {
	type args struct {
		c int
		r int
	}
	type testCase[I any] struct {
		name      string
		args      args
		wantUp    string
		wantDown  string
		wantLeft  string
		wantRight string
	}
	tests := []testCase[rune]{
		{name: "r0", args: args{c: 0, r: 0}, wantUp: "", wantDown: "1..", wantLeft: "", wantRight: "1."},
		{name: "r0", args: args{c: 1, r: 0}, wantUp: "", wantDown: "210", wantLeft: ".", wantRight: "."},
		{name: "r0", args: args{c: 2, r: 0}, wantUp: "", wantDown: "1..", wantLeft: "1.", wantRight: ""},

		{name: "r1", args: args{c: 0, r: 1}, wantUp: ".", wantDown: "..", wantLeft: "", wantRight: "21"},
		{name: "r1", args: args{c: 1, r: 1}, wantUp: "1", wantDown: "10", wantLeft: "1", wantRight: "1"},
		{name: "r1", args: args{c: 2, r: 1}, wantUp: ".", wantDown: "..", wantLeft: "12", wantRight: ""},

		{name: "r2", args: args{c: 0, r: 2}, wantUp: "1.", wantDown: ".", wantLeft: "", wantRight: "1."},
		{name: "r2", args: args{c: 1, r: 2}, wantUp: "21", wantDown: "0", wantLeft: ".", wantRight: "."},
		{name: "r2", args: args{c: 2, r: 2}, wantUp: "1.", wantDown: ".", wantLeft: "1.", wantRight: ""},

		{name: "r3", args: args{c: 0, r: 3}, wantUp: ".1.", wantDown: "", wantLeft: "", wantRight: "0."},
		{name: "r3", args: args{c: 1, r: 3}, wantUp: "121", wantDown: "", wantLeft: ".", wantRight: "."},
		{name: "r3", args: args{c: 2, r: 3}, wantUp: ".1.", wantDown: "", wantLeft: "0.", wantRight: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g = NewGrid(grid)
			g.Move(tt.args.r, tt.args.c)

			e := g.Explore()
			if tt.wantUp == "" && (e.Next(Up) || e.Up()) {
				t.Errorf("Up() did not overflow")
			}
			if tt.wantDown == "" && (e.Next(Down) || e.Down()) {
				t.Errorf("Down() did not overflow")
			}
			if tt.wantLeft == "" && (e.Next(Left) || e.Left()) {
				t.Errorf("Left() did not overflow")
			}
			if tt.wantRight == "" && (e.Next(Right) || e.Right()) {
				t.Errorf("Right() did not overflow")
			}

			mover := func(e Explorer[rune], direction int) string {
				res := []rune("")
				for e.Next(direction) {
					res = append(res, e.Current())
				}

				return string(res)
			}

			gotUp := mover(g.Explore(), Up)
			if tt.wantUp != gotUp {
				t.Errorf("Next(Up) gotValue = %v, want %v", gotUp, tt.wantUp)
			}
			gotDown := mover(g.Explore(), Down)
			if tt.wantDown != gotDown {
				t.Errorf("Next(Down) gotValue = %v, want %v", gotDown, tt.wantDown)
			}
			gotLeft := mover(g.Explore(), Left)
			if tt.wantLeft != gotLeft {
				t.Errorf("Next(Left) gotValue = %v, want %v", gotLeft, tt.wantLeft)
			}
			gotRight := mover(g.Explore(), Right)
			if tt.wantRight != gotRight {
				t.Errorf("Next(Right) gotValue = %v, want %v", gotRight, tt.wantRight)
			}
		})
	}
}
