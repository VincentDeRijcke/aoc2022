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

const (
	AIR          = '.'
	ROCK         = '#'
	MOVING       = '@'
	CHAMBER_SIZE = 10_000
)

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	instructions := []byte(strings.TrimSpace(input))

	chamber := newChamber(CHAMBER_SIZE)
	resultPart1 = part2(&chamber, 2022, instructions)

	chamber = newChamber(CHAMBER_SIZE)
	resultPart2 = part2(&chamber, 1_000_000_000_000, instructions)

	return
}

var shapes = []Shape{
	parseShape("..@@@@.."),
	/////////////
	parseShape(
		"...@.\n" +
			"..@@@\n" +
			"...@."),
	/////////////
	parseShape(
		"....@..\n" +
			"....@..\n" +
			"..@@@.."),
	/////////////
	parseShape(
		"..@....\n" +
			"..@....\n" +
			"..@....\n" +
			"..@...."),
	/////////////
	parseShape(
		"..@@...\n" +
			"..@@..."),
}

func parseShape(s string) Shape {
	shape := make(Shape, 0, 5)
	grid := utils.Reverse(utils.StringToRunes(s))

	for y, row := range grid {
		for x, r := range row {
			if r == '@' {
				shape = append(shape, Point{x, y})
			}
		}
	}

	return shape
}

type Record struct{ rockType, instruction, index, maxY int }

func part2(chamber *Chamber, rockCount int, instructions []byte) int {
	records := make([]Record, 0, 1000)
	i := 0
	maxY := -1
	doLoopDetection := true

	for r := 0; r < rockCount; r++ {
		fallSize := 0
		rock := Shape(utils.SliceCopy(shapes[r%5]))
		y := maxY + 4

		for moving := true; moving; i++ {
			if i == len(instructions) {
				i = 0
			}
			left := instructions[i] == '<'
			// Move
			rock.move(left, chamber, y)
			// Fall
			if y > 0 && rock.fall(chamber, y-1) {
				fallSize++
				if fallSize >= CHAMBER_SIZE {
					panic("Chamber too small")
				}
				y--
			} else {
				y += rock.settle(chamber, y)
				maxY = utils.Max(maxY, y)
				moving = false
			}
		}
		// Looping analysis
		// Thanks Keith Jones
		if doLoopDetection {
			current := Record{r % 5, i, r, maxY}
			foundPrevious := -1
			for h := len(records) - 1; h >= 0; h-- {
				if records[h].rockType == current.rockType && records[h].instruction == current.instruction {
					foundPrevious = h
					break
				}
			}
			records = append(records, current)
			if foundPrevious > 0 {
				start := records[foundPrevious]
				loopSize := current.index - start.index
				heightIncrease := current.maxY - start.maxY

				if foundPrevious-loopSize >= 0 {
					startPrevious := records[foundPrevious-loopSize]
					sameStage := startPrevious.rockType == start.rockType && startPrevious.instruction == start.instruction
					previousHeightIncrease := start.maxY - startPrevious.maxY
					sameIncrease := heightIncrease == previousHeightIncrease

					if sameIncrease && sameStage && rockCount-r >= loopSize {
						additionalLoops := (rockCount - r) / loopSize
						r += additionalLoops * loopSize
						maxY += additionalLoops * heightIncrease
						chamber.offsetY += additionalLoops * heightIncrease
						doLoopDetection = false
					}
				}
			}
		}
	}
	return maxY + 1
}

type Point struct{ x, y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Point) Down() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) Left() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x + 1, p.y}
}

type Shape []Point

func (s Shape) move(left bool, destination *Chamber, y int) bool {
	for _, point := range s {
		target := point.Left()
		if !left {
			target = point.Right()
		}
		if target.x < 0 || target.x > 6 || destination.get(target.x, y+target.y) != AIR {
			return false
		}
	}
	if left {
		for i, point := range s {
			s[i] = point.Left()
		}
	} else {
		for i, point := range s {
			s[i] = point.Right()
		}
	}

	return true
}

func (s Shape) fall(destination *Chamber, y int) bool {
	for _, point := range s {
		if destination.get(point.x, y+point.y) != AIR {
			return false
		}
	}

	return true
}

func (s Shape) settle(destination *Chamber, y int) int {
	maxY := 0
	for _, point := range s {
		destination.set(point.x, y+point.y, ROCK)
		maxY = utils.Max(maxY, point.y)
	}

	return maxY
}

type Chamber struct {
	grid           [][]rune
	floor, offsetY int
}

func newChamber(size int) Chamber {
	chamber := Chamber{utils.GridRunes(AIR, 7, size), 0, 0}

	return chamber
}
func (c *Chamber) get(x, y int) rune {
	actualY := c.getY(y)
	return c.grid[actualY][x]
}

func (c *Chamber) set(x int, y int, r rune) {
	actualY := c.getY(y)
	c.grid[actualY][x] = r
}

func (c *Chamber) getY(y int) int {
	actualY := y - c.offsetY + c.floor
	if actualY >= len(c.grid) {
		actualY -= len(c.grid)
		if actualY >= c.floor {
			c.grow(actualY - c.floor + 1)
		}
	}

	return actualY
}
func (c *Chamber) grow(growth int) {
	for i := 0; i < growth; i++ {
		for x := 0; x < 7; x++ {
			c.grid[c.floor][x] = AIR
		}
		c.floor++
		c.offsetY++
		if c.floor == len(c.grid) {
			c.floor = 0
		}
	}
}
func (c Chamber) String() string {
	return fmt.Sprintf("floor: %d, offset: %d\n%s", c.floor, c.offsetY, utils.RunesToString(utils.Reverse(c.grid)))
}
