package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type pos struct {
	i, j int
}

var tailLog map[string]pos

func record(p pos) {
	tailLog[fmt.Sprintf("(%d,%d)", p.i, p.j)] = p
}

func step(p pos, direction string) pos {
	switch direction {
	case "U":
		return pos{p.i, p.j + 1}
	case "D":
		return pos{p.i, p.j - 1}
	case "R":
		return pos{p.i + 1, p.j}
	case "L":
		return pos{p.i - 1, p.j}
	default:
		// no move
		return p
	}
}

func follow(h pos, t pos) pos {
	newt := pos{t.i, t.j}
	di, dj := h.i-t.i, h.j-t.j
	if di == 2 || di == -2 || dj == 2 || dj == -2 {
		// ..T..  .H.
		// H.T.H  ...
		// ..T..  TTT
		//        ...
		//        .H.

		if di > 0 {
			newt = step(newt, "R")
		} else if di < 0 {
			newt = step(newt, "L")
		}
		if dj > 0 {
			newt = step(newt, "U")
		} else if dj < 0 {
			newt = step(newt, "D")
		}
	}

	return newt
}

func move(h pos, t []pos, direction string) (pos, []pos) {
	newh := step(h, direction)

	newt := make([]pos, len(t))
	for k, _ := range t {
		newt[k] = pos{t[k].i, t[k].j}
		if k == 0 {
			newt[k] = follow(newh, newt[k])
		} else {
			newt[k] = follow(newt[k-1], newt[k])
		}
	}

	//printTailLog(tailLog, newh, newt)

	return newh, newt
}

func execMoves(lines []string, cordLength int) int {
	h := pos{0, 0}
	t := utils.SliceMap(make([]pos, cordLength-1), func(_ pos) pos { return pos{0, 0} })

	tailLog = make(map[string]pos, 10000)
	record(t[len(t)-1])

	//fmt.Println("Start:")
	//printTailLog(tailLog, h, t)

	for _, line := range lines {
		instructions := strings.Fields(line)
		steps, err := strconv.Atoi(instructions[1])
		utils.MaybePanic(err)

		//fmt.Println(line)
		for i := 0; i < steps; i++ {
			h, t = move(h, t, instructions[0])
			record(t[len(t)-1])
		}
	}

	fmt.Println("Final:")
	printTailLog(tailLog, h, t)

	return len(tailLog)
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	fmt.Println("Part1")
	resultPart1 = execMoves(lines, 2)
	fmt.Println("Part2")
	resultPart2 = execMoves(lines, 10)

	return resultPart1, resultPart2
}

func printTailLog(log map[string]pos, head pos, tail []pos) {
	mini, minj := utils.Min(0, head.i), utils.Min(0, head.j)
	maxi, maxj := head.i, head.j
	for _, v := range log {
		maxi = utils.Max(maxi, v.i)
		maxj = utils.Max(maxj, v.j)
		mini = utils.Min(mini, v.i)
		minj = utils.Min(minj, v.j)
	}
	for _, t := range tail {
		maxi = utils.Max(maxi, t.i)
		maxj = utils.Max(maxj, t.j)
		mini = utils.Min(mini, t.i)
		minj = utils.Min(minj, t.j)
	}

	grid := runes('.', maxi-mini+1, maxj-minj+1)
	for _, v := range log {
		grid[v.j-minj][v.i-mini] = '#'
	}
	for k := len(tail) - 1; k >= 0; k-- {
		r := rune('1' + k)
		if k == len(tail)-1 {
			r = 'T'
		}
		grid[tail[k].j-minj][tail[k].i-mini] = r
	}
	grid[head.j-minj][head.i-mini] = 'H'
	grid = utils.Reverse(grid)
	fmt.Println(utils.RunesToString(grid))
	fmt.Println()
}

func runes(r rune, cols int, rows int) [][]rune {
	grid := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]rune, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = r
		}
	}

	return grid
}

func main() {
	var content, err = os.ReadFile("./day9/input.txt")
	utils.MaybePanic(err)

	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
