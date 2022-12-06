package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
)

//func isMarker(s []rune) bool {
//	for i := 1; i < len(s); i++ {
//		if utils.Contains(s[0:i], s[i]) {
//			return false
//		}
//	}
//
//	return true
//}
//
//func indexMarker(runes []rune, size int, start int) (from, to int) {
//	for i := start; i < len(runes)-size+1; i++ {
//		if isMarker(runes[i : i+size]) {
//			return i, i + size
//		}
//	}
//	return -1, -1
//}

func indexMarker(runes []rune, size int, start int) (from, to int) {
	if len(runes)-start < size {
		return -1, -1
	}

	from = start
	to = start + 1
	for i := to; i < len(runes) && to-from < size; {
		found := utils.LastIndex(runes[from:to], runes[i])

		if found >= 0 {
			from += found + 1
		}
		i++
		to = i
	}
	return from, to
}

func resolve(input string) (resultPart1 int, resultPart2 int) {
	resultPart1 = -1
	resultPart2 = -1

	signal := []rune(input)

	if _, resultPart1 = indexMarker(signal, 4, 0); resultPart1 >= 0 {
		_, resultPart2 = indexMarker(signal, 14, resultPart1-4)
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
