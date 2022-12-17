package main

import (
	"fmt"
	"testing"
)

var example = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>
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
		{name: "Example", args: args{input: example}, want1: 3068, want2: 1_514_285_714_288},
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

func Test_chamber_set(t *testing.T) {
	chamber := newChamber(10)
	fmt.Println(chamber)
	for r := 'a'; r <= 'z'; r++ {
		chamber.set(int(r%7), int(r-'a'), r)
		fmt.Println(chamber)
		fmt.Println()
	}
	fmt.Println("Get Stuff")
	fmt.Println(string(chamber.get(1, 16)))
	fmt.Println(string(chamber.get(3, 25)))
	fmt.Println(string(chamber.get(0, 25)))
	fmt.Println(string(chamber.get(0, 0)))
	fmt.Println(chamber)
}
