package main

import (
	"aoc_go22/utils"
	_ "embed"
	"errors"
	"fmt"
	"strconv"
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
func resolve(input string) (resultPart1 int64, resultPart2 int64) {
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)

	resultPart1 = part1(utils.SliceMap(lines, parse), "root", false).val
	resultPart2 = part2(utils.SliceMap(lines, parse))

	return
}
func parse(line string) *Monkey {
	name, job, _ := strings.Cut(line, ": ")
	fields := strings.Fields(job)

	var monkey *Monkey
	if len(fields) == 1 {
		monkey = &Monkey{name: name, done: true, val: int64(utils.Atoi(fields[0]))}
	} else {
		monkey = &Monkey{name: name, a: fields[0], op: fields[1], b: fields[2]}
	}
	if monkey.name == "humn" {
		monkey.human = true
	}

	return monkey
}

type Monkey struct {
	name             string
	done, human      bool
	val              int64
	a, op, b         string
	humanA, humanB   *Monkey
	monkeyA, monkeyB int64
}

func (m *Monkey) calc(monkeys map[string]*Monkey, varHuman bool) bool {
	if m.done {
		return true
	}
	a, b := monkeys[m.a], monkeys[m.b]
	if a != nil && a.done && b != nil && b.done {
		m.done = true
		if a.human {
			m.human = true
			m.humanA = a
			m.monkeyB = b.val
		}
		if b.human {
			m.human = true
			m.humanB = b
			m.monkeyA = a.val
		}
		if b.human && a.human {
			panic(errors.New("Oopsie, this going to be hard"))
		}
		if !m.human || !varHuman {
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
		}

		return true
	}

	return false
}

func (m *Monkey) solve(i int64) int64 {
	if !m.human {
		panic(errors.New("nothing to solve dude"))
	}
	if m.name == "humn" {
		return i
	}

	// m.solve(i) == i = m.a 'op' m.b
	switch m.op {
	case "+":
		if m.humanA != nil {
			return m.humanA.solve(i - m.monkeyB)
		} else {
			return m.humanB.solve(i - m.monkeyA)
		}
	case "-":
		if m.humanA != nil {
			// i = human - j =>
			return m.humanA.solve(i + m.monkeyB)
		} else {
			// i = j - human =>
			return m.humanB.solve(m.monkeyA - i)
		}
	case "/":
		if m.humanA != nil {
			// i = human / j =>
			return m.humanA.solve(i * m.monkeyB)
		} else {
			// i = j / human
			return m.humanB.solve(m.monkeyA / i)
		}
	case "*":
		if m.humanA != nil {
			return m.humanA.solve(i / m.monkeyB)
		} else {
			return m.humanB.solve(i / m.monkeyA)
		}
	}
	panic(errors.New("WTH"))
}

func (m *Monkey) print() string {
	if m.name == "humn" {
		return "x"
	}

	if m.humanA != nil {
		return fmt.Sprint("(", m.humanA.print(), m.op, m.monkeyB, ")")
	} else if m.humanB != nil {
		return fmt.Sprint("(", m.monkeyA, m.op, m.humanB.print(), ")")
	} else {
		return strconv.FormatInt(m.val, 10)
	}
}

func part1(monkeys []*Monkey, rootName string, humanVar bool) *Monkey {
	var root *Monkey
	dones := make(map[string]*Monkey)

	for root == nil || !root.done {
		next := make([]*Monkey, 0, len(monkeys))
		for _, monkey := range monkeys {
			if monkey.calc(dones, humanVar) {
				dones[monkey.name] = monkey
			} else {
				next = append(next, monkey)
			}
			if root == nil && monkey.name == rootName {
				root = monkey
			}
		}
		monkeys = next
	}

	return root
}

func part2(monkeys []*Monkey) int64 {
	var root, humn *Monkey

	for _, monkey := range monkeys {
		if root == nil && monkey.name == "root" {
			root = monkey
		}
		if humn == nil && monkey.name == "humn" {
			humn = monkey
		}
	}

	a, b := part1(monkeys, root.a, true), part1(monkeys, root.b, true)
	//fmt.Println(root.a, " is human?", a.human, "val=", a.val)
	//fmt.Println(root.b, " is human?", b.human, "val=", b.val)
	//fmt.Println(a.print(), "=", b.print())

	if b.human && a.human {
		panic(errors.New("Not great"))
	}

	if a.human {
		return a.solve(b.val)
	} else {
		return b.solve(a.val)
	}
}
