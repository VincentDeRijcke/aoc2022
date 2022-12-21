package main

import (
	"testing"
)

var example = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`
var example1 = `root: a + b
humn: 5
a: ten - humn
b: 1
ten: 10
`
var example2 = `root: a + b
humn: 5
a: humn - ten
b: 1
ten: 10
`
var example3 = `root: a + b
humn: 50
a: humn / ten
b: 2
ten: 10
`
var example4 = `root: a + b
humn: 5
a: ten / humn
b: 2
ten: 10
`

func Test_resolve(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want1 int64
		want2 int64
	}{
		{name: "Example", args: args{input: example}, want1: 152, want2: 301},
		{name: "ex1", args: args{input: example1}, want1: 6, want2: 9},
		{name: "ex2", args: args{input: example2}, want1: -4, want2: 11},
		{name: "ex3", args: args{input: example3}, want1: 7, want2: 20},
		{name: "ex4", args: args{input: example4}, want1: 4, want2: 5},
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
