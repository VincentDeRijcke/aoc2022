package main

import (
	"reflect"
	"testing"
)

var example = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
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
		{name: "Example", args: args{input: example}, want1: 26, want2: 56000011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := resolve(tt.args.input, 10)
			if got1 != tt.want1 {
				t.Errorf("resolve() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("resolve() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func TestSensor_cover(t *testing.T) {
	tests := []struct {
		name string
		y    int
		want Cover
	}{
		{name: "Sensor at x=8, y=7: closest beacon is at x=2, y=10", y: 10, want: Cover{2, 14}},
		{name: "Sensor at x=8, y=7: closest beacon is at x=2, y=10", y: 16, want: Cover{8, 8}},
		{name: "Sensor at x=8, y=7: closest beacon is at x=2, y=10", y: 0, want: Cover{6, 10}},
		{name: "Sensor at x=8, y=7: closest beacon is at x=2, y=10", y: -2, want: Cover{8, 8}},
	}
	for _, tt := range tests {
		s := parseSensor(tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := s.cover(tt.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cover() = %v, want %v", got, tt.want)
			}
		})
	}
}
