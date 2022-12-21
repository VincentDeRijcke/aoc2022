package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
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

	resultPart1 = part1(utils.SliceMap(lines, parse))

	return
}
func parse(line string) *Monkey {
	name, job, _ := strings.Cut(line, ": ")
	fields := strings.Fields(job)

	if len(fields) == 1 {
		return &Monkey{name: name, done: true, val: utils.Atoi(fields[0])}
	} else {
		return &Monkey{name: name, a: fields[0], op: fields[1], b: fields[2]}
	}
}

type Monkey struct {
	name     string
	done     bool
	val      int
	a, op, b string
}

func (m *Monkey) calc(monkeys map[string]*Monkey) bool {
	if m.done {
		return true
	}
	a, b := monkeys[m.a], monkeys[m.b]
	if a != nil && a.done && b != nil && b.done {
		m.done = true
		switch m.op {
		case "+":
			m.val = a.val + b.val
		case "-":
			m.val = a.val - b.val
		case "/":
			m.val = a.val / b.val
		case "*":
			m.val = a.val * b.val
		}

		return true
	}

	return false
}

func part1(monkeys []*Monkey) int {
	var root *Monkey
	dones := make(map[string]*Monkey)

	for root == nil || !root.done {
		next := make([]*Monkey, 0, len(monkeys))
		for _, monkey := range monkeys {
			if monkey.calc(dones) {
				dones[monkey.name] = monkey
			} else {
				next = append(next, monkey)
			}
			if root == nil && monkey.name == "root" {
				root = monkey
			}
		}
		monkeys = next
	}

	return root.val
}
