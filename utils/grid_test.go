package utils

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	type args[T any] struct {
		in [][]T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[rune]{
		{name: "Nil", args: args[rune]{in: nil}, want: nil},
		{name: "1Line", args: args[rune]{in: StringToRunes("ABC")}, want: StringToRunes("A\nB\nC")},
		{name: "2Line", args: args[rune]{in: StringToRunes("" +
			"ABC \n" +
			"DEFX")}, want: StringToRunes("" +
			"AD\n" +
			"BE\n" +
			"CF\n" +
			" X")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut := Transpose(tt.args.in)
			if !reflect.DeepEqual(gotOut, tt.want) {
				t.Errorf("Transpose() = %v, want %v", RunesToString(gotOut), RunesToString(tt.want))
			}
		})
	}
}
