package main

import (
	"fmt"
	"os"
	"strings"
)

var lines []string
var stacks1 = [][]rune{
	/*1*/ []rune("RNPG"),
	/*2*/ []rune("TJBLCSVH"),
	/*3*/ []rune("TDBMNL"),
	/*4*/ []rune("RVPSB"),
	/*5*/ []rune("GCQSWMVH"),
	/*6*/ []rune("WQSCDBJ"),
	/*7*/ []rune("FQL"),
	/*8*/ []rune("WMHTDLFV"),
	/*9*/ []rune("LPBVMJF"),
}

var stacks2 = [][]rune{
	/*1*/ []rune("RNPG"),
	/*2*/ []rune("TJBLCSVH"),
	/*3*/ []rune("TDBMNL"),
	/*4*/ []rune("RVPSB"),
	/*5*/ []rune("GCQSWMVH"),
	/*6*/ []rune("WQSCDBJ"),
	/*7*/ []rune("FQL"),
	/*8*/ []rune("WMHTDLFV"),
	/*9*/ []rune("LPBVMJF"),
}

func build(line string) {

}

func move(line string) {
	fields, err := atois(strings.FieldsFunc(line, func(r rune) bool {
		return !(r >= '0' && r <= '9')
	}))
	check(err)
	count, from, to := fields[0], fields[1]-1, fields[2]-1

	for i := 0; i < count; i++ {
		lenFrom := len(stacks1[from])
		stacks1[to] = append(stacks1[to], stacks1[from][lenFrom-1])
		stacks1[from] = stacks1[from][0 : lenFrom-1]
	}
	lenFrom2 := len(stacks2[from])
	start2 := lenFrom2 - count
	stacks2[to] = append(stacks2[to], stacks2[from][start2:lenFrom2]...)
	stacks2[from] = stacks2[from][0:start2]
}

func stackTop(stacks [][]rune) string {
	top := make([]rune, 0, len(stacks))
	for _, s := range stacks {
		if len(s) > 0 {
			top = append(top, s[len(s)-1])
		}
	}
	return string(top)
}
func resolve(input string) (resultPart1 string, resultPart2 string) {
	lines := strings.Split(input, "\n")

	init := true
	for _, line := range lines {
		if line != "" {
			if init {
				build(line)
			} else {
				move(line)
			}
		} else {
			init = false
		}
	}

	return stackTop(stacks1), stackTop(stacks2)
}
func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day5/input.txt")
	check(err)
	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
}
