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
	lines = lines[:len(lines)-1]

	grid, moves := parse(lines)

	resultPart1 = part1(grid, moves)

	return
}
func parse(lines []string) ([][]rune, []*Move) {
	moves := parseMoves(lines[len(lines)-1])
	lines = lines[:len(lines)-2]
	rows, cols := len(lines), utils.Max(0, utils.SliceMap(lines, func(line string) int { return len(line) })...)

	grid := utils.GridRunes(EMPTY, cols+2, rows+2)
	for r, row := range lines {
		for c, tile := range row {
			grid[r+1][c+1] = tile
		}
	}

	return grid, moves
}
func parseMoves(line string) []*Move {
	fields := make([]string, 1, len(line))
	for i, r := range line {
		if r == 'R' || r == 'L' {
			fields = append(fields, "")
		}
		fields[len(fields)-1] += string([]rune{r})
		if (r == 'R' || r == 'L') && i < len(line)-1 {
			fields = append(fields, "")
		}
	}

	return utils.SliceMap(fields, func(field string) *Move {
		if field == "R" {
			return &Move{1, 0}
		} else if field == "L" {
			return &Move{-1, 0}
		}
		return &Move{0, utils.Atoi(field)}
	})
}

const (
	EMPTY = ' '
	OPEN  = '.'
	WALL  = '#'
	R     = 0
	D     = 1
	L     = 2
	U     = 3
)

type Move struct {
	turn, dist int
}

type Position struct {
	r, c, d int
}

func (p *Position) password() int {
	return 1000*(p.r) + 4*(p.c) + p.d
}

func (p *Position) turn(dir int) {
	p.d += dir
	if p.d > U {
		p.d = R
	} else if p.d < 0 {
		p.d = U
	}
}
func (p *Position) move(grid [][]rune) {
	nextR, nextC := p.r, p.c

	switch p.d {
	case R:
		nextC = p.c + 1
		if grid[p.r][nextC] == EMPTY {
			for nextC = p.c; grid[p.r][nextC] != EMPTY; nextC-- {
			}
			nextC++
		}
	case L:
		nextC = p.c - 1
		if grid[p.r][nextC] == EMPTY {
			for nextC = p.c; grid[p.r][nextC] != EMPTY; nextC++ {
			}
			nextC--
		}
	case U:
		nextR = p.r - 1
		if grid[nextR][p.c] == EMPTY {
			for nextR = p.r; grid[nextR][p.c] != EMPTY; nextR++ {
			}
			nextR--
		}
	case D:
		nextR = p.r + 1
		if grid[nextR][p.c] == EMPTY {
			for nextR = p.r; grid[nextR][p.c] != EMPTY; nextR-- {
			}
			nextR++
		}
	}
	if grid[nextR][nextC] == OPEN {
		p.r = nextR
		p.c = nextC
	}
}

func part1(grid [][]rune, moves []*Move) int {
	//fmt.Println(utils.RunesToString(grid))
	pos := &Position{1, strings.IndexRune(string(grid[1]), OPEN), R}

	for _, move := range moves {
		pos.turn(move.turn)
		for i := 0; i < move.dist; i++ {
			pos.move(grid)
		}
	}

	return pos.password()
}
