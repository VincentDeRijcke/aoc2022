package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strings"
)

var unprocessedStackLines [][]rune
var stacks1 [][]rune
var stacks2 [][]rune

// From
//
//	"    [D]    "
//	"[N] [C]    "
//	"[Z] [M] [P]"
//	" 1   2   3 "
//
// To
//
//	"ZN"
//	"MCD"
//	"P"
func buildStacks(lines [][]rune) [][]rune {
	lines = utils.Reverse(lines)
	_, indexes := utils.SliceFilterWithIndexes(lines[0], utils.IsDigit)

	stack := make([][]rune, len(indexes))
	for _, line := range lines[1:] {
		for i, index := range indexes {
			if line[index] != ' ' {
				stack[i] = append(stack[i], line[index])
			}
		}
	}

	return stack
}

func move(line string) {
	fields, err := utils.Atois(strings.FieldsFunc(line, utils.IsNotDigit))
	utils.MaybePanic(err)
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
		if init {
			if line != "" {
				unprocessedStackLines = append(unprocessedStackLines, []rune(line))
			} else {
				stacks1 = buildStacks(unprocessedStackLines)
				stacks2 = buildStacks(unprocessedStackLines)
				init = false
			}
		} else {
			if line != "" {
				move(line)
			}
		}
	}

	return stackTop(stacks1), stackTop(stacks2)
}
func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day5/input.txt")
	utils.MaybePanic(err)
	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
}
