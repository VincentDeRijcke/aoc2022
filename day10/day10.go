package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func execute(instructions []string, cycles []int) (int, string) {
	maxc := 240
	X, i := 1, 0
	signal := make([]int, maxc+1)
	screen, col, line := utils.GridRunes('.', 40, 6), 0, 0
	processing := ""

	for c := 1; c <= maxc; c++ {
		//during
		signal[c] = X * c
		//drawing
		if col == X-1 || col == X || col == X+1 {
			screen[line][col] = '#'
		}
		if col == 39 {
			col = 0
			line++
		} else {
			col++
		}

		// ending
		if instructions[i] == "noop" {
			i++
		} else {
			if processing == "" {
				processing = instructions[i]
			} else {
				adding, err := strconv.Atoi(strings.Fields(processing)[1])
				utils.MaybePanic(err)
				X = X + adding
				processing = ""
				i++
			}
		}
	}
	sum := 0
	for _, cycle := range cycles {
		sum += signal[cycle]
	}

	return sum, utils.RunesToString(screen)
}

func resolve(input string) (resultPart1 int, resultPart2 string) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	resultPart1, resultPart2 = execute(lines, []int{20, 60, 100, 140, 180, 220})

	return resultPart1, resultPart2
}

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:\n" + part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
