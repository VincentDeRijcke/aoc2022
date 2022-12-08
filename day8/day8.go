package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func newForestMaps(survey []string) ([][]rune, [][]bool, [][]int) {
	forest := make([][]rune, len(survey))
	visibility := make([][]bool, len(survey))
	score := make([][]int, len(survey))
	for i, r := range survey {
		forest[i] = make([]rune, len(r))
		visibility[i] = make([]bool, len(r))
		score[i] = make([]int, len(r))
		for j, c := range r {
			forest[i][j] = c
		}
	}

	return forest, visibility, score
}

func scanRow(forest [][]rune, i int, visibility [][]bool, scores [][]int) {
	columns := len(forest[0])

	var max rune
	for j := 0; j < columns; j++ {
		height := forest[i][j]
		if j == 0 || height > max {
			max = height
			visibility[i][j] = true
		}

		score := 0
		for k := j - 1; k >= 0; k-- {
			score++
			if forest[i][k] >= height {
				break
			}
		}
		scores[i][j] = score
	}
	for j := columns - 1; j >= 0; j-- {
		height := forest[i][j]
		if j == columns-1 || height > max {
			max = height
			visibility[i][j] = true
		}
		score := 0
		for k := j + 1; k < columns; k++ {
			score++
			if forest[i][k] >= height {
				break
			}
		}
		scores[i][j] *= score
	}
}

func scanColumn(forest [][]rune, j int, visibility [][]bool, scores [][]int) {
	rows := len(forest)

	var max rune
	for i := 0; i < rows; i++ {
		height := forest[i][j]

		if i == 0 || height > max {
			max = height
			visibility[i][j] = true
		}
		score := 0
		for k := i - 1; k >= 0; k-- {
			score++
			if forest[k][j] >= height {
				break
			}
		}
		scores[i][j] *= score
	}
	for i := rows - 1; i >= 0; i-- {
		height := forest[i][j]
		if i == rows-1 || height > max {
			max = height
			visibility[i][j] = true
		}
		score := 0
		for k := i + 1; k < rows; k++ {
			score++
			if forest[k][j] >= height {
				break
			}
		}
		scores[i][j] *= score
	}
}

func scan(forest [][]rune, visibility [][]bool, scores [][]int) (int, int) {
	rows := len(forest)
	columns := len(forest[0])

	// Scan Rows
	for i := 0; i < rows; i++ {
		scanRow(forest, i, visibility, scores)
	}

	// Scan Columns
	for j := 0; j < columns; j++ {
		scanColumn(forest, j, visibility, scores)
	}

	visibleCount := 0
	for _, row := range visibility {
		for _, visible := range row {
			if visible {
				visibleCount++
			}
		}
	}

	maxScore := 0
	for _, row := range scores {
		for _, score := range row {
			if score > maxScore {
				maxScore = score
			}
		}
	}

	//fmt.Println(forest)
	//fmt.Println(visibility)
	//fmt.Println(scores)

	return visibleCount, maxScore
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	survey := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)
	forest, visibility, score := newForestMaps(survey)
	resultPart1, resultPart2 = scan(forest, visibility, score)

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
