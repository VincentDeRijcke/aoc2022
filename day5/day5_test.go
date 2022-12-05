package main

import (
	"reflect"
	"testing"
)

const fullExample = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_resolve(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want1 string
		want2 string
	}{
		{name: "Example Full", args: args{input: fullExample}, want1: "CMZ", want2: "MCD"},
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

func Test_buildStacks(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{name: "Example", args: args{lines: "    [D]    \n" +
			"[N] [C]    \n" +
			"[Z] [M] [P]\n" +
			" 1   2   3 "},
			want: [][]rune{
				/*1*/ []rune("ZN"),
				/*2*/ []rune("MCD"),
				/*3*/ []rune("P"),
			}},
		{name: "Real", args: args{lines: "    [H]         [H]         [V]    \n" +
			"    [V]         [V] [J]     [F] [F]\n" +
			"    [S] [L]     [M] [B]     [L] [J]\n" +
			"    [C] [N] [B] [W] [D]     [D] [M]\n" +
			"[G] [L] [M] [S] [S] [C]     [T] [V]\n" +
			"[P] [B] [B] [P] [Q] [S] [L] [H] [B]\n" +
			"[N] [J] [D] [V] [C] [Q] [Q] [M] [P]\n" +
			"[R] [T] [T] [R] [G] [W] [F] [W] [L]\n" +
			" 1   2   3   4   5   6   7   8   9 "},
			want: [][]rune{
				/*1*/ []rune("RNPG"),
				/*2*/ []rune("TJBLCSVH"),
				/*3*/ []rune("TDBMNL"),
				/*4*/ []rune("RVPSB"),
				/*5*/ []rune("GCQSWMVH"),
				/*6*/ []rune("WQSCDBJ"),
				/*7*/ []rune("FQL"),
				/*8*/ []rune("WMHTDLFV"),
				/*9*/ []rune("LPBVMJF"),
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStacks(sliceMap(splits(tt.args.lines, "\n"), func(s string) []rune { return []rune(s) })); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
