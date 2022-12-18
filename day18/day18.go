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
	grid := render(cubes)
	resultPart1 = part1(grid)

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

func render(cubes []*Cube) [][][]bool {
	grid := make([][][]bool, 20)
	for z := 0; z < len(grid); z++ {
		grid[z] = make([][]bool, 20)
		for y := 0; y < len(grid[z]); y++ {
			grid[z][y] = make([]bool, 20)
		}
	}
	for _, cube := range cubes {
		grid[cube.z][cube.y][cube.x] = true
	}

	return grid
}

type Cube struct{ x, y, z int }

func part1(grid [][][]bool) int {
	surface := 0
	for z, level := range grid {
		for y, row := range level {
			for x, cube := range row {
				if cube {
					if x == 0 || !grid[z][y][x-1] {
						surface++
					}
					if x == len(row)-1 || !grid[z][y][x+1] {
						surface++
					}
					if y == 0 || !grid[z][y-1][x] {
						surface++
					}
					if y == len(level)-1 || !grid[z][y+1][x] {
						surface++
					}
					if z == 0 || !grid[z-1][y][x] {
						surface++
					}
					if z == len(grid)-1 || !grid[z+1][y][x] {
						surface++
					}
				}
			}
		}
	}

	return surface
}
