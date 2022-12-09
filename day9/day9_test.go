package main

import "testing"

var example = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

var largerExample = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
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
		{name: "Example", args: args{input: example}, want1: 13, want2: 1},
		{name: "LargerExample", args: args{input: largerExample}, want1: 88, want2: 36},
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
