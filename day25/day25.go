package main

import (
	"aoc_go22/utils"
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
func resolve(input string) (resultPart1 string, resultPart2 int) {
	lines := strings.Split(input, "\n")

	fuel := utils.SliceMap(lines[:len(lines)-1], parseSnafu)
	total := 0
	for _, f := range fuel {
		total += f
	}

	return toSnafu(total), 0
}
func parseSnafu(line string) int {
	digits := []rune(line)
	val := 0
	for _, digit := range digits {
		val = val*5 + value(digit)
	}

	return val
}
func value(r rune) int {
	switch r {
	case '0', '1', '2':
		return int(r - '0')
	case '-':
		return -1
	case '=':
		return -2
	}
	panic(errors.New("unknown digit " + string(r)))
}

func toSnafu(i int) string {
	digits := []rune("012=-")
	var snafu []rune
	for i > 0 {
		r := i % 5
		i = i / 5
		snafu = append(snafu, digits[r])
		if r > 2 {
			i++
		}
	}

	return string(utils.Reverse(snafu))
}
