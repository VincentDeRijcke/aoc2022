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
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)

	resultPart1 = part1(parse(lines))
	resultPart2 = part2(parse(lines))

	return
}
func parse(lines []string) *File {
	values := utils.Atois(lines)

	file := &File{make([]*Number, 0, len(values)), nil}
	for _, value := range values {
		file.add(value)
	}
	file.postInit()

	return file
}

type Number struct {
	val  int
	prev *Number
	next *Number
}

func (c *Number) shift(pos int) {
	if pos == 0 {
		//fmt.Println("No move")
		return
	}

	t := c.lookup(pos)
	c.prev.next = c.next
	c.next.prev = c.prev
	if pos > 0 {
		n := t.next
		//fmt.Println("Shift", c.val, "between", t.val, "and", n.val)
		c.prev = t
		c.next = n
		n.prev = c
		t.next = c
	} else if pos < 0 {
		p := t.prev
		//fmt.Println("Shift", c.val, "between", p.val, "and", t.val)
		c.prev = p
		c.next = t
		p.next = c
		t.prev = c
	}
}

func (n *Number) lookup(pos int) *Number {
	cur := n
	if pos > 0 {
		for i := 0; i < pos; i++ {
			cur = cur.next
		}
	} else if pos < 0 {
		for i := pos; i < 0; i++ {
			cur = cur.prev
		}
	}

	return cur
}

type File struct {
	numbers []*Number
	start   *Number
}

func (f *File) add(value int) {
	current := &Number{value, nil, nil}
	f.numbers = append(f.numbers, current)
	if len(f.numbers) > 1 {
		current.prev = f.numbers[len(f.numbers)-2]
		current.prev.next = current
	}
	if value == 0 {
		f.start = current
	}
}
func (f *File) postInit() {
	f.numbers[len(f.numbers)-1].next = f.numbers[0]
	f.numbers[0].prev = f.numbers[len(f.numbers)-1]
}

func (f File) String() string {
	s := make([]string, 0, len(f.numbers))
	for i, current := 0, f.start; i < len(f.numbers); i++ {
		s = append(s, strconv.Itoa(current.val))
		current = current.next
		if current.prev.next != current || current.next.prev != current {
			panic(errors.New("Problem"))
		}
	}

	return "[" + strings.Join(s, "\n") + "]\n"
}

func part1(file *File) int {
	//fmt.Println(file)

	for _, current := range file.numbers {
		current.shift(current.val % (len(file.numbers) - 1))
	}

	r1 := file.start.lookup(1000)
	r2 := r1.lookup(1000)
	r3 := r2.lookup(1000)
	fmt.Println("s[1000]=", r1.val, "s[2000]=", r2.val, "s[3000]=", r3.val)
	//utils.StringToFile("day20/part1.txt", fmt.Sprint(file))
	return r1.val + r2.val + r3.val
}

func part2(file *File) int {
	for _, current := range file.numbers {
		current.val *= 811589153
	}
	for i := 0; i < 10; i++ {
		for _, current := range file.numbers {
			current.shift(current.val % (len(file.numbers) - 1))
		}
	}
	r1 := file.start.lookup(1000)
	r2 := r1.lookup(1000)
	r3 := r2.lookup(1000)
	fmt.Println("s[1000]=", r1.val, "s[2000]=", r2.val, "s[3000]=", r3.val)
	//utils.StringToFile("day20/part2.txt", fmt.Sprint(file))
	return r1.val + r2.val + r3.val
}
