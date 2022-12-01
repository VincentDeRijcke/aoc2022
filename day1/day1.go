package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var elfIndex = 1
var elves = make([]*elf, 300)[0:0]

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type elf struct {
	index    int
	calories int
}

func newElf() *elf {
	elf := new(elf)
	elf.index = elfIndex
	elf.calories = 0

	elfIndex++

	return elf
}

func main() {
	fmt.Println("Reading input")
	var content, err = os.ReadFile("./day1/input.txt")
	check(err)
	var scanner = bufio.NewScanner(bytes.NewReader(content))

	elf := newElf()
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Println("Elf", elf.index, "hold", elf.calories, "calories")
			elves = append(elves, elf)
			elf = newElf()
		} else {
			var c, err = strconv.Atoi(text)
			check(err)
			elf.calories += c
		}
	}
	fmt.Println("Found", len(elves), "elves")
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	fmt.Println("The 3 best elves are:")
	totalCalories := 0
	for _, e := range elves[0:3] {
		fmt.Println("Elf", e.index, "hold", e.calories, "calories")
		totalCalories += e.calories
	}
	fmt.Println("Total calories for the best elves:", totalCalories)
}
