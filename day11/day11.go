package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type monkey struct {
	items       []int
	divider     int
	operator    string
	operand     int
	trueTarget  int
	falseTarget int
}

func (m *monkey) op(level int) int {
	switch m.operator {
	case "*":
		return level * m.operand
	case "+":
		return level + m.operand
	default: // old * old
		return level * level
	}
}
func (m *monkey) test(level int) int {
	if level%m.divider == 0 {
		return m.trueTarget
	}
	return m.falseTarget
}

func parse(lines []string) []*monkey {
	var current *monkey
	monkeys := make([]*monkey, 0, 8)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "Monkey") {
			current = new(monkey)
			monkeys = append(monkeys, current)
		} else if strings.HasPrefix(line, "  Starting items: ") {
			current.items = utils.Atois(strings.Split(strings.TrimPrefix(line, "  Starting items: "), ", "))
		} else if strings.HasPrefix(line, "  Operation: ") {
			fields := strings.Fields(strings.TrimPrefix(line, "  Operation: new = old "))
			right, old := strconv.Atoi(fields[1])
			if old != nil {
				current.operator = "* old"
			} else {
				current.operator = fields[0]
				current.operand = right
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			current.divider, _ = strconv.Atoi(strings.TrimPrefix(line, "  Test: divisible by "))
			current.trueTarget, _ = strconv.Atoi(strings.TrimPrefix(lines[i+1], "    If true: throw to monkey "))
			current.falseTarget, _ = strconv.Atoi(strings.TrimPrefix(lines[i+2], "    If false: throw to monkey "))
		}
	}
	return monkeys
}
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	monkeys := parse(lines)
	part1 := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, _ := range monkeys {
			for _, item := range monkeys[i].items {
				part1[i]++
				item = monkeys[i].op(item) / 3
				target := monkeys[i].test(item)
				monkeys[target].items = append(monkeys[target].items, item)
			}
			monkeys[i].items = nil
		}
	}
	sort.Ints(part1)
	resultPart1 = part1[len(part1)-1] * part1[len(part1)-2]

	monkeys = parse(lines)
	part2 := make([]int, len(monkeys))
	commonDivider := 1
	for _, m := range monkeys {
		commonDivider *= m.divider
	}

	for round := 1; round <= 10000; round++ {
		for i, _ := range monkeys {
			for _, item := range monkeys[i].items {
				part2[i]++
				item = monkeys[i].op(item) % commonDivider
				target := monkeys[i].test(item)
				monkeys[target].items = append(monkeys[target].items, item)
			}
			monkeys[i].items = nil
		}
	}

	sort.Ints(part2)
	resultPart2 = part2[len(part2)-1] * part2[len(part2)-2]

	return resultPart1, resultPart2
}

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
