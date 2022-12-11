package utils

import (
	"reflect"
	"testing"
)

func Test_atois(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Nil", args: args{slice: nil}, want: nil},
		{name: "Empty", args: args{slice: []string{}}, want: []int{}},
		{name: "NoErr", args: args{slice: []string{"1", "-1", "100", "-100"}}, want: []int{1, -1, 100, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Atois(tt.args.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Atois() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransposeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Nil", args: args{s: ""}, want: ""},
		{name: "SingleChar", args: args{s: "a"}, want: "a"},
		{name: "OneLine", args: args{s: "a.b"}, want: "" +
			"a\n" +
			".\n" +
			"b"},
		{name: "MultiLine", args: args{s: "" +
			"a.b\n" +
			"c.d"}, want: "" +
			"ca\n" +
			"..\n" +
			"db"},
		{name: "Triangle", args: args{s: "" +
			"1   \n" +
			"22  \n" +
			"333 \n" +
			"4444"}, want: "" +
			"4321\n" +
			"432 \n" +
			"43  \n" +
			"4   "},
		{name: "Pyramid", args: args{s: "" +
			"1   \n" +
			"22  \n" +
			"333 \n" +
			"4444\n" +
			"333 \n" +
			"22  \n" +
			"1   "}, want: "" +
			"1234321\n" +
			" 23432 \n" +
			"  343  \n" +
			"   4   "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateString(tt.args.s); got != tt.want {
				t.Errorf("RotateString() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
