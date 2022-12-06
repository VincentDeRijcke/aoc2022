package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
)

func contains(runes []rune, x rune) bool {
	for _, r := range runes {
		if r == x {
			return true
		}
	}
	return false
}
func isMarker(s []rune) bool {
	for i := 1; i < len(s); i++ {
		if contains(s[0:i], s[i]) {
			return false
		}
	}

	return true
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	signal := []rune(input)
	for i := 0; i < len(signal)-3; i++ {
		if isMarker(signal[i : i+4]) {
			resultPart1 = i + 4
			break
		}
	}
	for i := 0; i < len(signal)-13; i++ {
		if isMarker(signal[i : i+14]) {
			resultPart2 = i + 14
			break
		}
	}

	return
}
func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day6/input.txt")
	utils.MaybePanic(err)
	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
}
