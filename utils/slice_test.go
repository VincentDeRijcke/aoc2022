package utils

import (
	"reflect"
	"testing"
)

func Test_splitSlice(t *testing.T) {
	type args struct {
		slice []string
		sep   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{slice: nil, sep: ";"}, want: nil},
		{name: "blank", args: args{slice: []string{}, sep: ";"}, want: []string{}},
		{name: "nosplit", args: args{slice: []string{"a,b", "c,d"}, sep: ";"}, want: []string{"a,b", "c,d"}},
		{name: "abcd", args: args{slice: []string{"a,b", "c,d"}, sep: ","}, want: []string{"a", "b", "c", "d"}},
		{name: "abcd", args: args{slice: []string{"a,b,c", "d"}, sep: ","}, want: []string{"a", "b", "c", "d"}},
		{name: "abcd", args: args{slice: []string{"a,b,c,d"}, sep: ","}, want: []string{"a", "b", "c", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitSlice(tt.args.slice, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splits(t *testing.T) {
	type args struct {
		s    string
		seps []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Blank", args: args{s: "", seps: []string{";"}}, want: []string{""}},
		{name: "Blank", args: args{s: "", seps: []string{";", ","}}, want: []string{""}},
		{name: "CSV", args: args{s: "a,b,c,d", seps: []string{";", ",", "\t"}}, want: []string{"a", "b", "c", "d"}},
		{name: "CSV", args: args{s: "a;b;c;d", seps: []string{";", ",", "\t"}}, want: []string{"a", "b", "c", "d"}},
		{name: "CSV", args: args{s: "a,b;c,d", seps: []string{";", ",", "\t"}}, want: []string{"a", "b", "c", "d"}},
		{name: "CSV", args: args{s: "\ta,b;c,d", seps: []string{";", ",", "\t"}}, want: []string{"", "a", "b", "c", "d"}},
		{name: "Weird", args: args{s: "abcd", seps: []string{"b", "c"}}, want: []string{"a", "", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Splits(tt.args.s, tt.args.seps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Splits() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_reverse(t *testing.T) {
	type args[T any] struct {
		in []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "Nil", args: args[int]{in: nil}, want: nil},
		{name: "Empty", args: args[int]{in: []int{}}, want: []int{}},
		{name: "NotEmpty", args: args[int]{in: []int{1}}, want: []int{1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2}}, want: []int{2, 1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2, 3}}, want: []int{3, 2, 1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2, 3, 4}}, want: []int{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
			if len(got) > 1 {
				if got[0] == tt.args.in[0] {
					t.Errorf("ReverseInPlace() = Inplace reversed")
				}
			}
		})
	}
}
func Test_reverseInPlace(t *testing.T) {
	type args[T any] struct {
		in []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{name: "Nil", args: args[int]{in: nil}, want: nil},
		{name: "Empty", args: args[int]{in: []int{}}, want: []int{}},
		{name: "NotEmpty", args: args[int]{in: []int{1}}, want: []int{1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2}}, want: []int{2, 1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2, 3}}, want: []int{3, 2, 1}},
		{name: "NotEmpty", args: args[int]{in: []int{1, 2, 3, 4}}, want: []int{4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseInPlace(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInPlace() = %v, want %v", got, tt.want)
			}
			if len(got) > 0 {
				if got[0] != tt.args.in[0] {
					t.Errorf("ReverseInPlace() = Not inplace reversed")
				}
			}
		})
	}
}
