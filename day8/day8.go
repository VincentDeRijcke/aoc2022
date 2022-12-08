package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strings"
	"time"
)

func scan(survey []string) (visibleCount int, maxScore int) {
	rows := utils.Runes(survey)
	columns := utils.Transpose(rows)

	for i, _ := range rows {
		for j, _ := range rows[i] {
			height := rows[i][j]

			up := utils.Reverse(columns[j][:i])
			down := columns[j][i+1:]

			left := utils.Reverse(rows[i][:j])
			right := rows[i][j+1:]

			if visible(height, up) || visible(height, down) || visible(height, left) || visible(height, right) {
				visibleCount++
			}

			totalScore := score(height, up) * score(height, down) * score(height, left) * score(height, right)
			if totalScore > maxScore {
				maxScore = totalScore
			}
		}
	}

	return
}

func visible(height rune, line []rune) bool {
	for _, other := range line {
		if other >= height {
			return false
		}
	}

	return true
}

func score(height rune, line []rune) int {
	result := 0
	for _, other := range line {
		result++
		if other >= height {
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
