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
	AIR    = '.'
	ROCK   = '#'
	MOVING = '@'
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

	rocks, height := getRocks(2022)
	chamber := utils.GridRunes(AIR, 7, height+10)
	resultPart1 = part1(chamber, rocks, instructions)
	//rocks, height := getRocks(1_000_000_000_000)
	//chamber := utils.GridRunes(AIR, 7, height+10)
	//resultPart1 = part1(chamber, rocks, instructions)

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

func getRocks(num int) ([]Shape, int) {
	maxHeight := 0
	rocks := make([]Shape, num)
	for i := 0; i < num; i++ {
		rocks[i] = utils.SliceCopy(shapes[i%5])
		maxHeight += len(rocks[i])
	}

	return rocks, maxHeight
}

func part1(chamber [][]rune, rocks []Shape, instructions []byte) int {
	i := 0
	maxY := -1

	for _, rock := range rocks {
		y := maxY + 4

		//fmt.Println("Rock", r+1, "falling")
		//view(chamber[:y+5], rock, y)
		//fmt.Println()
		for moving := true; moving; i++ {
			if i == len(instructions) {
				i = 0
			}
			left := instructions[i] == '<'
			// Move
			rock.move(left, chamber[y:])

			//if left {
			//	fmt.Println(i+1, ": Move Left ")
			//} else {
			//	fmt.Println(i+1, ": Move Right ")
			//}
			//view(chamber[:y+4], rock, y)
			// Fall
			if y > 0 && rock.fall(chamber[y-1:]) {
				y--
				//fmt.Println("Falls 1 unit")
				//view(chamber[:y+4], rock, y)
			} else {
				y += rock.settle(chamber[y:])
				maxY = utils.Max(maxY, y)
				moving = false

				//fmt.Println("Rock", r+1, "come to rest maxY:", maxY)
				//fmt.Println(utils.RunesToString(utils.Reverse(chamber[:maxY+1])))
				//fmt.Println()
			}
		}
	}

	return maxY + 1
}

func view(chamber [][]rune, rock Shape, rockY int) {
	chamber = utils.GridCopy(chamber)
	for _, p := range rock {
		chamber[p.y+rockY][p.x] = MOVING
	}
	fmt.Println(utils.RunesToString(utils.Reverse(chamber)))
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

func (s Shape) move(left bool, destination [][]rune) bool {
	for _, point := range s {
		target := point.Left()
		if !left {
			target = point.Right()
		}
		if target.x < 0 || target.x > 6 || destination[target.y][target.x] != AIR {
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

func (s Shape) fall(destination [][]rune) bool {
	for _, point := range s {
		if destination[point.y][point.x] != AIR {
			return false
		}
	}

	return true
}

func (s Shape) settle(destination [][]rune) int {
	maxY := 0
	for _, point := range s {
		destination[point.y][point.x] = ROCK
		maxY = utils.Max(maxY, point.y)
	}

	return maxY
}
