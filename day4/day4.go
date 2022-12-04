package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitSlice(slice []string, sep string) []string {
	var res []string
	for _, s := range slice {
		res = append(res, strings.Split(s, sep)...)
	}

	return res
}

func splits(s string, seps ...string) []string {
	res := []string{s}
	for _, sep := range seps {
		res = splitSlice(res, sep)
	}

	return res
}

func atois(slice []string) ([]int, error) {
	res := make([]int, len(slice))
	for i, s := range slice {
		var err error
		res[i], err = strconv.Atoi(s)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func getRanges(line string) (l1 int, h1 int, l2 int, h2 int) {
	ints, err := atois(splits(line, ",", "-"))
	check(err)

	return ints[0], ints[1], ints[2], ints[3]
}

func resolve(input string) (int, int) {
	lines := strings.Split(input, "\n")
	resultPart1, resultPart2 := 0, 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			l1, h1, l2, h2 := getRanges(line)
			if (l1 >= l2 && h1 <= h2) || (l2 >= l1 && h2 <= h1) {
				resultPart1++
			}
			if !(l1 > h2 || h1 < l2 || l2 > h1 || h2 < l1) {
				resultPart2++
			}
		}
	}

	return resultPart1, resultPart2
}
func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day4/input.txt")
	check(err)
	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
}
