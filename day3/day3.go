package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findBackpackErrorPriority(backpack string) int {
	size := len(backpack)
	compartment1, compartment2 := backpack[0:size/2], backpack[size/2:]
	for _, item := range compartment1 {
		if strings.ContainsRune(compartment2, item) {
			if item < 'a' {
				return int(item - 'A' + 27)
			} else {
				return int(item - 'a' + 1)
			}
		}
	}
	return 0
}

func findGroupTagPriority(group []string) int {
	backpack1, backpack2, backpack3 := group[0], group[1], group[2]
	for _, item := range backpack1 {
		if strings.ContainsRune(backpack2, item) && strings.ContainsRune(backpack3, item) {
			if item < 'a' {
				return int(item - 'A' + 27)
			} else {
				return int(item - 'a' + 1)
			}
		}
	}
	return 0
}

func sumErrorsPriorities(game io.Reader) (int, int) {
	var scanner = bufio.NewScanner(game)

	sum1, sum2 := 0, 0
	group := make([]string, 0, 3)
	for scanner.Scan() {
		backpack := strings.TrimSpace(scanner.Text())
		if backpack != "" {
			sum1 += findBackpackErrorPriority(backpack)
			if len(group) < 3 {
				group = append(group, backpack)
			}
			if len(group) == 3 {
				sum2 += findGroupTagPriority(group)
				group = group[0:0]
			}
		}
	}

	return sum1, sum2
}

func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day3/input.txt")
	check(err)
	part1, part2 := sumErrorsPriorities(bytes.NewReader(content))

	fmt.Println("Sum Priorities Part 1:", part1)
	fmt.Println("Sum Priorities Part 2:", part2)
}
