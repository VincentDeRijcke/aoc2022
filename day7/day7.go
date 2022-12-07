package main

import (
	"aoc_go22/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type file struct {
	name     string
	dir      bool
	size     int
	parent   *file
	children map[string]*file
}

var fs = []*file{{name: "/", dir: true, children: map[string]*file{}}}
var root = fs[0]

func (c *file) cd(name string) *file {
	if name == ".." {
		return c.parent
	} else {
		return c.children[name]
	}
}
func (c *file) mkdir(name string) {
	fs = append(fs, &file{name: name, dir: true, parent: c, children: map[string]*file{}})
	c.children[name] = fs[len(fs)-1]
}
func (c *file) fallocate(name string, size int) {
	fs = append(fs, &file{name: name, parent: c, size: size})
	c.children[name] = fs[len(fs)-1]

	for d := c; d != nil; d = d.parent {
		d.size += size
	}
}
func parseTerminal(lines []string) {
	var current *file
	for _, line := range lines {
		if line == "$ cd /" {
			current = root
		} else if line == "$ ls" {
			// noop
		} else if strings.HasPrefix(line, "$ cd ") {
			current = current.cd(strings.TrimPrefix(line, "$ cd "))
		} else if strings.HasPrefix(line, "dir ") {
			current.mkdir(strings.TrimPrefix(line, "dir "))
		} else if line != "" {
			fields := strings.Fields(line)
			size, err := strconv.Atoi(fields[0])
			utils.MaybePanic(err)
			name := fields[1]
			current.fallocate(name, size)
		}
	}
}
func resolve(input string) (resultPart1 int, resultPart2 int) {
	lines := utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace)

	parseTerminal(lines)

	capacity, target := 70_000_000, 30_000_000
	free := capacity - root.size
	needed := target - free

	resultPart2 = capacity
	for _, f := range fs {
		if f.dir {
			if f.size <= 100000 {
				resultPart1 += f.size
			}
			if f.size >= needed && f.size < resultPart2 {
				resultPart2 = f.size
			}
		}
	}

	return
}
func main() {
	var content, err = os.ReadFile("./day7/input.txt")
	utils.MaybePanic(err)

	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(string(content))

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}
