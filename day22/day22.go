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
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	grid, moves := parse(lines)

	resultPart1 = part1(grid, moves)
	resultPart2 = part2(grid, moves)

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

func (p *Position) String() string {
	return fmt.Sprint(*p)
}
func (p *Position) password() int {
	return password(p.r, p.c, p.d)
}
func password(row, col, dir int) int {
	return 1000*(row) + 4*(col) + dir
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

func (p *Position) move3d(grid [][]rune, folding []*Position) {
	nextR, nextC := p.r, p.c

	switch p.d {
	case R:
		nextC = p.c + 1
		if grid[p.r][nextC] == EMPTY {
			p.fold(grid, folding)
		}
	case L:
		nextC = p.c - 1
		if grid[p.r][nextC] == EMPTY {
			p.fold(grid, folding)
		}
	case U:
		nextR = p.r - 1
		if grid[nextR][p.c] == EMPTY {
			p.fold(grid, folding)
		}
	case D:
		nextR = p.r + 1
		if grid[nextR][p.c] == EMPTY {
			p.fold(grid, folding)
		}
	}
	if grid[nextR][nextC] == OPEN {
		p.r = nextR
		p.c = nextC
	}
}

func (p *Position) fold(grid [][]rune, folding []*Position) {
	nextPos := folding[p.password()]
	if nextPos == nil {
		panic(errors.New("No folding for " + p.String()))
	}
	switch r := grid[nextPos.r][nextPos.c]; r {
	case OPEN:
		p.r = nextPos.r
		p.c = nextPos.c
		p.d = nextPos.d
	case WALL: // no move
	default:
		panic(errors.New("Folding for " + p.String() + " is " + nextPos.String() + " but is '" + string(r) + "'"))
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
func part2(grid [][]rune, moves []*Move) int {
	pos := &Position{1, strings.IndexRune(string(grid[1]), OPEN), R}
	rules := folding(grid)

	for _, move := range moves {
		pos.turn(move.turn)
		for i := 0; i < move.dist; i++ {
			pos.move3d(grid, rules)
		}
	}

	return pos.password()
}

func folding(grid [][]rune) []*Position {
	instructions := make([]*Position, 300_000)
	if len(grid) < 50 {
		// Test cube
		//      - - 1 -
		//	    2 3 4 -
		//	    - - 5 6
		for i := 0; i < 4; i++ {
			// L1 -> D3
			instructions[password(1+i, 9, L)] = &Position{5, 5 + i, D}
			// U3 --> R1
			instructions[password(5, 5+i, U)] = &Position{1 + i, 9, R}
			// R1 --> L6 (inverted)
			instructions[password(1+i, 12, R)] = &Position{12 - i, 16, L}
			// R6 (inverted) --> L1
			instructions[password(12-i, 16, R)] = &Position{1 + i, 12, L}
			// U1 --> D2 (inverted)
			instructions[password(1, 9+i, U)] = &Position{5, 5 - i, D}
			// U2 (inverted) --> D1
			instructions[password(5, 5-i, U)] = &Position{1, 9 + i, D}

			// L2 --> U6 (inverted)
			instructions[password(5+i, 1, L)] = &Position{12, 16 - i, U}
			// D6 (inverted) --> R2
			instructions[password(12, 16-i, D)] = &Position{5 + i, 1, R}
			// D2 --> U5 (inverted)
			instructions[password(8, 1+i, L)] = &Position{12, 12 - i, U}
			// D5 (inverted) --> U2
			instructions[password(12, 12-i, D)] = &Position{8, 1 + i, U}

			// R4 --> D6 (inverted)
			instructions[password(5+i, 12, R)] = &Position{9, 16 - i, D}
			// U6 (inverted) --> L4
			instructions[password(9, 16-i, U)] = &Position{5 + i, 12, L}

			// D3 --> R5 (inverted)
			instructions[password(8, 5+i, D)] = &Position{12 - i, 9, R}
			// L5 (inverted) --> U3
			instructions[password(12-i, 9, L)] = &Position{8, 5 + i, U}

		}
	} else {
		// my cube
		//
		//  c  c c  c c  c c
		//          1 1  1 1
		//     5 5  0 0  5 5
		//  1  0 1  0 1  0 1
		//  ---- 1111 2222 r1
		//  ---- 1111 2222
		//  ---- 1111 2222
		//  ---- 1111 2222 r50
		//  ---- 3333 ---- r51
		//  ---- 3333 ----
		//  ---- 3333 ----
		//  ---- 3333 ---- r100
		//  4444 5555 ---- r101
		//  4444 5555 ----
		//  4444 5555 ----
		//  4444 5555 ---- r150
		//  6666 ---- ---- r151
		//  6666 ---- ----
		//  6666 ---- ----
		//  6666 ---- ---- r200
		for i := 0; i < 50; i++ {
			// U1  --> R6  // L6  --> D1
			instructions[password(1, 51+i, U)] = &Position{151 + i, 1, R}
			instructions[password(151+i, 1, L)] = &Position{1, 51 + i, D}
			// L1  --> R4i // L4i --> R1
			instructions[password(1+i, 51, L)] = &Position{150 - i, 1, R}
			instructions[password(150-i, 1, L)] = &Position{1 + i, 51, R}
			// U2  --> U6  // D6  --> D2
			instructions[password(1, 101+i, U)] = &Position{200, 1 + i, U}
			instructions[password(200, 1+i, D)] = &Position{1, 101 + i, D}
			// R2  --> L5i // R5i --> L2
			instructions[password(1+i, 150, R)] = &Position{150 - i, 100, L}
			instructions[password(150-i, 100, R)] = &Position{1 + i, 150, L}
			// D2  --> L3  // R3  --> U2
			instructions[password(50, 101+i, D)] = &Position{51 + i, 100, L}
			instructions[password(51+i, 100, R)] = &Position{50, 101 + i, U}
			// L3  --> D4  // U4  --> R3
			instructions[password(51+i, 51, L)] = &Position{101, 1 + i, D}
			instructions[password(101, 1+i, U)] = &Position{51 + i, 51, R}
			// D5  --> L6  // R6  --> U5
			instructions[password(150, 51+i, D)] = &Position{151 + i, 50, L}
			instructions[password(151+i, 50, R)] = &Position{150, 51 + i, U}
		}
	}

	return instructions
}
