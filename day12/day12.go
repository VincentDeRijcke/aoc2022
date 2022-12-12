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

type Pos struct{ x, y int }

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Path []Pos

func (path Path) String() string {
	return strings.Join(utils.SliceMap(path, func(pos Pos) string { return pos.String() }), "->")
}

func findElevation(heightmap [][]rune, target rune) []Pos {
	var found []Pos
	for y, row := range heightmap {
		for x, height := range row {
			if height == target {
				found = append(found, Pos{x, y})
			}
		}
	}
	return found
}

func findPath(heightmap [][]rune, start rune) int {
	rows, cols := utils.GridSizes(heightmap)
	visited := utils.GridRunes('.', cols, rows)

	var rootPaths []Path
	roots := findElevation(heightmap, start)
	for _, root := range roots {
		visited[roots[0].y][roots[0].x] = start
		rootPaths = append(rootPaths, Path{root})
	}

	path := explore(heightmap, visited, rootPaths)

	if path != nil {
		fmt.Println(sprintPath(path, heightmap))
		return len(path) - 1
	}

	return 0
}

func sprintPath(path Path, heightmap [][]rune) string {
	rows, cols := utils.GridSizes(heightmap)
	pathmap := utils.GridRunes('.', cols, rows)

	for _, pos := range path {
		pathmap[pos.y][pos.x] = heightmap[pos.y][pos.x]
	}

	return utils.RunesToString(pathmap)
}
func sprintPaths(paths []Path, heightmap [][]rune, visited [][]rune) string {
	pathmap := utils.GridCopy(heightmap)

	for i, row := range visited {
		for j, v := range row {
			if v != '.' {
				pathmap[i][j] = '.'
			}
		}
	}

	for _, path := range paths {
		for i, pos := range path {
			pathmap[pos.y][pos.x] = '⊙'
			if i == len(path)-1 {
				pathmap[pos.y][pos.x] = '█'
			}
		}
	}

	return utils.RunesToString(pathmap)
}

func explore(heightmap [][]rune, visited [][]rune, paths []Path) Path {
	nextPaths := make([]Path, 0, 100)

	fmt.Println("Step:", len(paths[0]), " Candidates:", len(paths))
	fmt.Println(sprintPaths(paths, heightmap, visited))

	for _, path := range paths {
		current := path[len(path)-1]
		x, y, h := current.x, current.y, heightmap[current.y][current.x]
		if h == 'S' {
			h = 'a'
		}

		for _, next := range []*Pos{step(x+1, y, h, heightmap), step(x-1, y, h, heightmap), step(x, y+1, h, heightmap), step(x, y-1, h, heightmap)} {
			if next != nil && visited[next.y][next.x] == '.' {
				visited[next.y][next.x] = '#'
				nextPath := make(Path, len(path), len(path)+1)
				copy(nextPath, path)
				nextPath = append(nextPath, *next)
				if heightmap[next.y][next.x] == 'E' {
					visited[next.y][next.x] = 'E'
					return nextPath
				} else {
					nextPaths = append(nextPaths, nextPath)
				}
			}

		}
	}

	if len(nextPaths) > 0 {
		return explore(heightmap, visited, nextPaths)
	}

	return nil
}

func step(x int, y int, h rune, heightmap [][]rune) *Pos {
	if x >= 0 && y >= 0 && y < len(heightmap) && x < len(heightmap[0]) {
		dh := heightmap[y][x]
		if dh == 'E' {
			dh = 'z'
		}
		if dh <= h+1 {
			return &Pos{x, y}
		}
	}

	return nil
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	heightmap := utils.Runes(lines)
	resultPart1 = findPath(heightmap, 'S')
	resultPart2 = findPath(heightmap, 'a')
	if resultPart2 > resultPart1 {
		resultPart2 = resultPart1
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
