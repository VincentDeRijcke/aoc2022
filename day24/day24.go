package main

import (
	"aoc_go22/utils"
	_ "embed"
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
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := strings.Split(input, "\n")
	grid := utils.Runes(lines[:len(lines)-1])

	valley := parse(grid)
	resultPart1 = part1(valley, 1, valley.entry, valley.exit)
	resultPart2 = part2(valley, resultPart1)

	return
}

func part1(valley *Valley, min int, start Pos, goal Pos) int {

	pos, end := map[string]Pos{start.String(): start}, false
	for m := min; !end; m++ {
		fmt.Println("Minute", m)
		pos, end = move(pos, valley, m, goal)
		if end {
			return m
		}
	}

	return 0
}

func part2(valley *Valley, part1End int) int {
	time := part1(valley, part1End+1, valley.exit, valley.entry)
	time = part1(valley, time+1, valley.entry, valley.exit)
	return time
}
func move(start map[string]Pos, valley *Valley, m int, goal Pos) (map[string]Pos, bool) {
	next := make(map[string]Pos)
	for _, pos := range start {
		for _, mv := range []Pos{pos, {pos.x, pos.y - 1}, {pos.x, pos.y + 1}, {pos.x - 1, pos.y}, {pos.x + 1, pos.y}} {
			if mv == goal {
				return nil, true
			} else if valley.isFree(mv, m) {
				next[mv.String()] = mv
			}
		}
	}

	return next, false
}

type Pos struct {
	x, y int
}

func (p Pos) String() string {
	return fmt.Sprint("(", p.x, p.y, ")")
}

type Valley struct {
	rows, cols  int
	vert, hori  [][]Blizzard
	entry, exit Pos
}

func (v *Valley) isFree(p Pos, min int) bool {
	if p == v.entry || p == v.exit {
		return true
	}
	if p.y > 0 && p.x > 0 && p.x < v.cols-1 && p.y < v.rows-1 {
		for _, b := range v.hori[p.y] {
			if p.x == b.pos(min) {
				return false
			}
		}
		for _, b := range v.vert[p.x] {
			if p.y == b.pos(min) {
				return false
			}
		}
		return true
	}

	return false
}

type Blizzard struct {
	from, max, dir int
}

func (b Blizzard) pos(min int) int {
	dist := b.dir * min % b.max
	pos := b.from + dist
	if pos > b.max {
		pos -= b.max
	} else if pos < 1 {
		pos += b.max
	}

	return pos
}

func parse(grid [][]rune) *Valley {
	rows, cols := utils.GridSizes(grid)
	valley := Valley{rows, cols, make([][]Blizzard, cols), make([][]Blizzard, rows), Pos{1, 0}, Pos{cols - 2, rows - 1}}

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			switch grid[r][c] {
			case '>':
				valley.hori[r] = append(valley.hori[r], Blizzard{c, cols - 2, 1})
			case '<':
				valley.hori[r] = append(valley.hori[r], Blizzard{c, cols - 2, -1})
			case '^':
				valley.vert[c] = append(valley.vert[c], Blizzard{r, rows - 2, -1})
			case 'v':
				valley.vert[c] = append(valley.vert[c], Blizzard{r, rows - 2, 1})
			}
		}
	}

	return &valley
}
