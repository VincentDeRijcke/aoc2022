package main

import (
	"aoc_go22/utils"
	_ "embed"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	ORDERED = iota
	NOT_ORDERED
	CONTINUE
)

type List []any

func ToList(a any) List {
	switch i := a.(type) {
	case List:
		return i
	case int:
		return List{i}[:]
	default:
		panic(errors.New("not a List"))
	}
}

func ordered(left List, right List) int {
	for i := 0; i < len(left); i++ {
		if i == len(right) {
			// Right ran out of items
			return NOT_ORDERED
		}
		if utils.IsInt(left[i]) && utils.IsInt(right[i]) {
			if left[i].(int) < right[i].(int) {
				return ORDERED
			} else if left[i].(int) > right[i].(int) {
				return NOT_ORDERED
			}
			// else continue
		} else {
			res := ordered(ToList(left[i]), ToList(right[i]))
			if res != CONTINUE {
				return res
			}
		}
	}
	if len(left) == len(right) {
		return CONTINUE
	}
	// Left ran out of items
	return ORDERED
}

func parse(s string) List {
	var l List
	var fields []string
	if s[0] == '[' {
		s = s[1 : len(s)-1]
		fields = strings.FieldsFunc(s, splitter())
	} else {
		fields = []string{s}
	}
	for _, f := range fields {
		if f[0] == '[' {
			l = append(l, parse(f))
		} else {
			l = append(l, utils.Atoi(f))
		}
	}

	return l
}
func splitter() func(r rune) bool {
	level := 0
	return func(r rune) bool {
		switch r {
		case '[':
			level++
		case ']':
			level--
		case ',':
			if level == 0 {
				return true
			} else {
				return false
			}
		}
		return false
	}
}
func checkPackets(packets []string) int {
	o := 0
	for i := 1; i <= len(packets)/2; i++ {
		left := parse(packets[(i-1)*2])
		right := parse(packets[(i-1)*2+1])
		if ordered(left, right) == ORDERED {
			o += i
		}
	}

	return o
}

func searchPacket(packets []string, packet string) int {
	for i := 1; i <= len(packets); i++ {
		if packets[i-1] == packet {
			return i
		}
	}

	panic(errors.New("signal disappeared, wtf"))
}

func orderPackets(packets []string) int {
	packets = append(packets, "[[2]]", "[[6]]")
	sort.Slice(packets, func(i int, j int) bool {
		if ordered(parse(packets[i]), parse(packets[j])) == ORDERED {
			return true
		}
		return false
	})
	fmt.Println(strings.Join(packets, "\n"))

	return searchPacket(packets, "[[2]]") * searchPacket(packets, "[[6]]")
}
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)

	resultPart1 = checkPackets(lines)
	resultPart2 = orderPackets(lines)

	return
}

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
