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

const (
	SOURCE = '+'
	AIR    = ' '
	ROCK   = '█'
	SAND   = 'o'
)

var COLORS = map[rune]color.Color{
	SOURCE: color.Gray{20},
	AIR:    color.Black,
	ROCK:   color.White,
	SAND:   color.RGBA{194, 178, 128, 0xff},
}

type Point struct{ x, y int }

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
func (p Point) Up() Point {
	return Point{p.x, p.y - 1}
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

type Edge struct{ from, to Point }

func (e *Edge) points() []Point {
	minx, maxx := utils.Min(e.from.x, e.to.x), utils.Max(e.from.x, e.to.x)
	miny, maxy := utils.Min(e.from.y, e.to.y), utils.Max(e.from.y, e.to.y)

	points := make([]Point, 0, utils.Max(maxx-minx, maxy-miny)+1)

	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			points = append(points, Point{x, y})
		}
	}

	return points
}

type Cave struct {
	source     Point
	minx, maxx int
	miny, maxy int
	content    [][]rune
}

func (cave Cave) String() string {
	return fmt.Sprintf("%v\n%s %v\n", Point{cave.minx, cave.miny}, utils.RunesToString(cave.content), Point{cave.maxx, cave.maxy})
}

func (cave *Cave) addPoint(p Point, r rune) {
	cave.content[p.y-cave.miny][p.x-cave.minx] = r
}

func (cave *Cave) getPoint(p Point) rune {
	return cave.content[p.y-cave.miny][p.x-cave.minx]
}

func (cave *Cave) addEdge(e Edge) {
	for _, p := range e.points() {
		cave.addPoint(p, ROCK)
	}
}

func (cave *Cave) addFloor() {
	oc := *cave

	cave.maxy += 2
	cave.maxx = utils.Max(cave.maxx, cave.source.x+cave.maxy)
	cave.minx = utils.Min(cave.minx, cave.source.x-cave.maxy)
	cave.content = utils.GridRunes(AIR, cave.maxx-cave.minx+1, cave.maxy-cave.miny+1)

	for y := oc.miny; y <= oc.maxy; y++ {
		for x := oc.minx; x <= oc.maxx; x++ {
			cave.addPoint(Point{x, y}, oc.getPoint(Point{x, y}))
		}
	}
	for x := cave.minx; x <= cave.maxx; x++ {
		cave.addPoint(Point{x, cave.maxy}, ROCK)
	}
}

func (cave *Cave) isFalling(p Point) bool {
	return p.x < cave.minx || p.x > cave.maxx || p.y > cave.maxy || p.y < cave.miny
}

func (cave *Cave) isSolid(p Point) bool {
	return cave.getPoint(p) != AIR
}
func (cave *Cave) addSand() bool {
	p := cave.source
	for {
		down := p.Down()
		if cave.isFalling(down) {
			return false
		}
		if cave.isSolid(down) {
			left := down.Left()
			if cave.isFalling(left) {
				return false
			}
			if cave.isSolid(left) {
				right := down.Right()
				if cave.isFalling(right) {
					return false
				}
				if cave.isSolid(right) {
					if p == cave.source {
						return false
					}
					cave.addPoint(p, SAND)
					return true
				} else {
					p = right
				}
			} else {
				p = left
			}
		} else {
			p = down
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)
	_ = lines

	var frames = make([]*image.Paletted, 0)
	cave := parse(lines)
	for ; cave.addSand(); resultPart1++ {
		frames = append(frames, utils.RunesToBlocks(cave.content, COLORS, 6, 6))
	}

	frames = append(frames, utils.RunesToBlocks(cave.content, COLORS, 6, 6))
	gif := utils.Animate(frames)
	utils.SaveGif(gif, "day14/part1.gif")

	cave.addFloor()
	frames = make([]*image.Paletted, 0)
	for resultPart2 = resultPart1 + 1; cave.addSand(); resultPart2++ {
		if resultPart2%100 == 0 {
			frames = append(frames, utils.RunesToBlocks(cave.content, COLORS, 6, 6))
		}
	}
	frames = append(frames, utils.RunesToBlocks(cave.content, COLORS, 6, 6))
	gif = utils.Animate(frames)
	utils.SaveGif(gif, "day14/part2.gif")

	return
}

func parse(lines []string) Cave {
	minx, miny := 500, 0
	maxx, maxy := 500, 0
	var edges []Edge
	for _, line := range lines {
		points := utils.SliceMap(strings.Split(line, " -> "), parsePoint)
		for i := 1; i < len(points); i++ {
			edges = append(edges, Edge{points[i-1], points[i]})
			minx = utils.Min(minx, points[i-1].x, points[i].x)
			miny = utils.Min(miny, points[i-1].y, points[i].y)
			maxx = utils.Max(maxx, points[i-1].x, points[i].x)
			maxy = utils.Max(maxy, points[i-1].y, points[i].y)
		}
	}
	source := Point{500, 0}
	cave := Cave{source, minx, maxx, miny, maxy, utils.GridRunes(AIR, maxx-minx+1, maxy-miny+1)}
	cave.addPoint(source, SOURCE)
	for _, edge := range edges {
		cave.addEdge(edge)
	}

	return cave
}

func parsePoint(s string) Point {
	ints := utils.Atois(utils.Splits(s, ","))

	return Point{ints[0], ints[1]}
}
