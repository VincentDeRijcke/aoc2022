package main

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
			if got := splitSlice(tt.args.slice, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSlice() = %v, want %v", got, tt.want)
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
			if got := splits(tt.args.s, tt.args.seps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atois(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "Nil", args: args{slice: nil}, want: nil, wantErr: false},
		{name: "Empty", args: args{slice: []string{}}, want: []int{}, wantErr: false},
		{name: "NoErr", args: args{slice: []string{"1", "-1", "100", "-100"}}, want: []int{1, -1, 100, -100}, wantErr: false},
		{name: "Err", args: args{slice: []string{"1", "-1", "bou", "-100"}}, want: []int{1, -1, 0, 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := atois(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("atois() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atois() got = %v, want %v", got, tt.want)
			}
		})
	}
}
