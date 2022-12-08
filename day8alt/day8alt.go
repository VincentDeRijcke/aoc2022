package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func scan(survey []string) (visibleCount int, maxScore int) {
	forest := utils.Runes(survey)
	grid := utils.NewGrid(forest)
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			height := grid.Move(i, j)
			if visible(height, grid.Explore(), utils.Up) || visible(height, grid.Explore(), utils.Down) || visible(height, grid.Explore(), utils.Left) || visible(height, grid.Explore(), utils.Right) {
				visibleCount++
			}

			totalScore := score(height, grid.Explore(), utils.Up) * score(height, grid.Explore(), utils.Down) * score(height, grid.Explore(), utils.Left) * score(height, grid.Explore(), utils.Right)
			if totalScore > maxScore {
				maxScore = totalScore
			}
		}
	}

	return
}

func visible(height rune, explorer utils.Explorer[rune], direction int) bool {
	for explorer.Next(direction) {
		if explorer.Current() >= height {
			return false
		}
	}

	return true
}

func score(height rune, explorer utils.Explorer[rune], direction int) int {
	result := 0
	for explorer.Next(direction) {
		result++
		if explorer.Current() >= height {
			break
		}
	}

	return result
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	survey := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)

	resultPart1, resultPart2 = scan(survey)

	return
}

func main() {
	var content, err = os.ReadFile("./day8/input.txt")
	utils.MaybePanic(err)

	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
