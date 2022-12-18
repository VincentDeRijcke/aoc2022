package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	GRID_SIZE = 20 // Eyeballed in input
	EMPTY     = ' '
	AIR       = '.'
	STONE     = '#'
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
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)
	cubes := utils.SliceMap(lines, parseCube)

	sortCubes(cubes)
	resultPart1 = part1(render(cubes, AIR))
	resultPart2 = part2(render(cubes, EMPTY))

	return
}

func parseCube(line string) *Cube {
	cube := Cube{}
	fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)

	return &cube
}

func sortCubes(cubes []*Cube) {
	byZYX := func(i, j int) bool {
		return cubes[i].z < cubes[j].z || cubes[i].y < cubes[j].y || cubes[i].x < cubes[j].x
	}
	sort.Slice(cubes, byZYX)
}

func render(cubes []*Cube, r rune) [][][]rune {
	grid := make([][][]rune, GRID_SIZE)
	for z := 0; z < len(grid); z++ {
		grid[z] = utils.GridRunes(r, GRID_SIZE, GRID_SIZE)
	}
	for _, cube := range cubes {
		grid[cube.z][cube.y][cube.x] = STONE
	}

	return grid
}

type Cube struct{ x, y, z int }

func part1(grid [][][]rune) int {
	surface := 0
	for z, level := range grid {
		for y, row := range level {
			for x, cube := range row {
				if cube == STONE {
					if x == 0 || grid[z][y][x-1] == AIR {
						surface++
					}
					if x == len(row)-1 || grid[z][y][x+1] == AIR {
						surface++
					}
					if y == 0 || grid[z][y-1][x] == AIR {
						surface++
					}
					if y == len(level)-1 || grid[z][y+1][x] == AIR {
						surface++
					}
					if z == 0 || grid[z-1][y][x] == AIR {
						surface++
					}
					if z == len(grid)-1 || grid[z+1][y][x] == AIR {
						surface++
					}
				}
			}
		}
	}

	return surface
}

func part2(grid [][][]rune) int {
	for filling, up := true, true; filling; up = !up {
		filling = false
		start, step := 0, 1
		if !up {
			start, step = len(grid)-1, -1
		}
		for z := start; z >= 0 && z < len(grid); z += step {
			level := grid[z]
			for y, row := range level {
				for x, _ := range row {
					if grid[z][y][x] == EMPTY {
						if x == 0 || grid[z][y][x-1] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
						if x == len(row)-1 || grid[z][y][x+1] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
						if y == 0 || grid[z][y-1][x] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
						if y == len(level)-1 || grid[z][y+1][x] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
						if z == 0 || grid[z-1][y][x] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
						if z == len(grid)-1 || grid[z+1][y][x] == AIR {
							grid[z][y][x] = AIR
							filling = true
						}
					}
				}
			}
		}
	}

	return part1(grid)
}
