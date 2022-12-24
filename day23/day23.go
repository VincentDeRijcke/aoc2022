package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"image"
	"image/color"
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
	scan := utils.Runes(lines[:len(lines)-1])

	resultPart1 = part1(scan)
	resultPart2 = part2(scan)

	return
}

const (
	N    = -1
	S    = 1
	W    = 0
	E    = 2
	ELF  = '#'
	OPEN = '.'
)

var rounds = [][]int{{N, S, W, E}, {S, W, E, N}, {W, E, N, S}, {E, N, S, W}}

type Ground struct {
	n, s, e, w int
	elves      []*Elf
	from, to   *ElfPos
}

func (g Ground) String() string {
	return fmt.Sprintf("(%d, %d) -> (%d, %d)\n", g.n, g.w, g.s, g.e) + g.from.String(g.n, g.s, g.w, g.e) //+ "\nMoves\n" + g.to.String(g.n, g.s, g.w, g.e)
}

func (p *ElfPos) String(n, s, w, e int) string {
	rows := p.pos[n+p.offset : s+p.offset+1]
	ss := utils.SliceMap(rows, func(row []*Elf) string {
		str := []rune(strings.Repeat(".", e-w+1))
		for i := 0; i < len(str); i++ {
			if row[w+p.offset+i] != nil {
				str[i] = ELF
			}
		}
		return string(str)
	})

	return strings.Join(ss, "\n")
}

func (p *ElfPos) Runes() [][]rune {
	return utils.GridMap(p.pos, func(e *Elf) rune {
		if e != nil {
			return ELF
		}
		return OPEN
	})
}

type Elf struct {
	x, y            int
	moving, blocked bool
	toX, toY        int
}

type ElfPos struct {
	offset int
	pos    [][]*Elf
}

func (p *ElfPos) get(x, y int) *Elf {
	return p.pos[y+p.offset][x+p.offset]
}
func (p *ElfPos) set(x int, y int, elf *Elf) {
	p.pos[y+p.offset][x+p.offset] = elf
}

func part1(scan [][]rune) int {
	ground := newGround(scan)
	//fmt.Println("== Initial State ==")
	//fmt.Println(*ground)

	for i := 0; i < 10; i++ {
		round := rounds[i%4]
		if consider(ground, round) {
			move(ground)
		}
		//fmt.Println("== End of Round ", i+1, " ==")
		//fmt.Println(*ground)
	}

	return ((ground.s - ground.n + 1) * (ground.e - ground.w + 1)) - len(ground.elves)
}

func part2(scan [][]rune) int {
	ground := newGround(scan)
	colors := map[rune]color.Color{OPEN: color.Black, ELF: utils.Elf}
	imgs := make([]*image.Paletted, 0, 2000)
	imgs = append(imgs, utils.RunesToBlocks(ground.from.Runes(), colors, 4, 4))

	for i := 0; true; i++ {
		round := rounds[i%4]
		if consider(ground, round) {
			move(ground)
			imgs = append(imgs, utils.RunesToBlocks(ground.from.Runes(), colors, 4, 4))
		} else {
			fmt.Println("== Final of Round ", i+1, " ==")
			//fmt.Println(*ground)
			imgs = append(imgs, utils.RunesToBlocks(ground.from.Runes(), colors, 4, 4))
			utils.SaveGif(utils.Animate(imgs), "part2.gif")
			return i + 1
		}
		if i%1000 == 0 {
			fmt.Println("== End of Round ", i+1, " ==")
			//fmt.Println(*ground)
		}
	}

	return 0
}

func consider(ground *Ground, directions []int) bool {
	oneNeedsMoving := false

	for _, dir := range directions {
		for _, elf := range ground.elves {
			x, y := elf.x, elf.y
			toX, toY := x, y

			needMoving := ground.from.get(x+1, y) != nil ||
				ground.from.get(x-1, y) != nil ||
				ground.from.get(x, y+1) != nil ||
				ground.from.get(x+1, y+1) != nil ||
				ground.from.get(x-1, y+1) != nil ||
				ground.from.get(x, y-1) != nil ||
				ground.from.get(x+1, y-1) != nil ||
				ground.from.get(x-1, y-1) != nil

			oneNeedsMoving = oneNeedsMoving || needMoving
			moving := !elf.moving && needMoving

			if moving {
				switch dir {
				case N, S:
					toY += dir
					if ground.from.get(toX, toY) == nil &&
						ground.from.get(toX-1, toY) == nil &&
						ground.from.get(toX+1, toY) == nil {
						elf.moving = true
						elf.toX, elf.toY = toX, toY
					}
				case W, E:
					toX += dir - 1
					if ground.from.get(toX, toY) == nil &&
						ground.from.get(toX, toY-1) == nil &&
						ground.from.get(toX, toY+1) == nil {
						elf.moving = true
						elf.toX, elf.toY = toX, toY
					}
				}
				if elf.moving {
					if ground.to.get(toX, toY) == nil {
						ground.to.set(toX, toY, elf)
					} else {
						ground.to.get(toX, toY).blocked = true
						elf.blocked = true
					}
				}
			}
		}
	}

	return oneNeedsMoving
}
func move(ground *Ground) {
	n, s, w, e := ground.s, ground.n, ground.e, ground.w
	for _, elf := range ground.elves {
		if elf.moving {
			if !elf.blocked {
				ground.from.set(elf.x, elf.y, nil)
				ground.from.set(elf.toX, elf.toY, elf)
				elf.x, elf.y = elf.toX, elf.toY
			}
			elf.moving = false
			elf.blocked = false
			ground.to.set(elf.toX, elf.toY, nil)
		}
		n = utils.Min(n, elf.toY)
		s = utils.Max(s, elf.toY)
		w = utils.Min(w, elf.toX)
		e = utils.Max(e, elf.toX)
	}
	ground.n, ground.s, ground.w, ground.e = n, s, w, e
}

func newGround(scan [][]rune) *Ground {
	rows, cols := utils.GridSizes(scan)
	pad := 100 // Tried a few times until large enough
	grid := Ground{0, rows - 1, cols - 1, 0, make([]*Elf, 0, rows*cols), newElfPos(cols, rows, pad), newElfPos(cols, rows, pad)}

	for y, row := range scan {
		for x, val := range row {
			if val == ELF {
				elf := &Elf{x: x, y: y}
				grid.elves = append(grid.elves, elf)
				grid.from.set(x, y, elf)
			}
		}
	}

	return &grid
}

func newElfPos(cols int, rows int, offset int) *ElfPos {
	grid := make([][]*Elf, rows+2*offset)
	for i := 0; i < rows+2*offset; i++ {
		grid[i] = make([]*Elf, cols+2*offset)
	}

	return &ElfPos{offset, grid}
}
