package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*

u/korylprince

Go: https://github.com/korylprince/adventofcode/blob/master/2022/16/main.go
I originally solved today's problem with Python 3, but I wasn't happy that it took more than 1 second to solve,
so I decided to use Go and attempt to eke out as much efficiency as I could.
I got both parts including parsing the input file to under 100ms.
I wanted to share some things I learned along the way.
I started by finding the shortest path for all valves via floyd-warshall.
Then I converted all valves to a bitfield (luckly there were 15 non-zero flow valves plus the starting valve, so everything fit nicely into a uint16).
Then the flow rates and shortest paths were encoded in slices indexes.
Part 1 used DFS, keeping track of the max pressure along the way.
Part 2 used DFS again, returning all paths found (encoded in a bitfield). The paths were filtered to only include those with at least half the pressure from part 1, then compared all combinations to find the max pressure.
Some interesting efficiency things I learned along the way:
using bitfields makes things a lot faster and easier
adding to a path or set is just path | newnode
nodes become indexes to slices and can be OR'd for combined lookup: slice[v1|v2]
encoding things in slices (e.g. slice index -> value) is much faster than map lookups. This was the biggest gains I got
Raw iteration for combinations (i in range(0, max); j in range(i+1, max)) was much faster than using gonum's combination.

*/

var valveregexp = regexp.MustCompile(`([A-Z]{2}).*=(\d+);.*?((?:[A-Z]{2}(?:, )?)+)`)

type Valve struct {
	Name  string
	Rate  uint16
	Edges string
}

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
	valves := parse(input)

	graph := floydWarshall(valves)

	// pick valves with flow and starting point
	var mvps []*Valve
	for _, v := range valves {
		if v.Rate > 0 || v.Name == "AA" {
			mvps = append(mvps, v)
		}
	}

	// assign bits
	bitfield := make(map[*Valve]uint16)
	for idx, v := range mvps {
		bitfield[v] = 1 << idx
	}

	// find start
	var start uint16
	for _, v := range mvps {
		if v.Name == "AA" {
			start = bitfield[v]
			break
		}
	}

	// create slice for fast edge lookup
	fastDist := make([]uint16, 0xffff)
	for _, v1 := range mvps {
		for _, v2 := range mvps {
			fastDist[bitfield[v1]|bitfield[v2]] = graph[v1][v2]
		}
	}

	// create slice for fast node lookup
	fastMVP := make([][2]uint16, len(mvps))
	for idx, v := range mvps {
		fastMVP[idx] = [2]uint16{bitfield[v], v.Rate}
	}

	// part 1
	var dfs func(target, pressure, minute, on, node uint16) uint16
	dfs = func(target, pressure, minute, on, node uint16) uint16 {
		max := pressure
		for _, w := range fastMVP {
			if node == w[0] || w[0] == start || w[0]&on != 0 {
				continue
			}
			l := fastDist[node|w[0]] + 1
			if minute+l > target {
				continue
			}
			if next := dfs(target, pressure+(target-minute-l)*w[1], minute+l, on|w[0], w[0]); next > max {
				max = next
			}
		}
		return max
	}

	part1 := dfs(30, 0, 0, 0, start)

	// part 2
	var dfspaths func(target, pressure, minute, on, node, path uint16) [][2]uint16
	dfspaths = func(target, pressure, minute, on, node, path uint16) [][2]uint16 {
		paths := [][2]uint16{{pressure, path}}
		for _, w := range fastMVP {
			if w[0] == node || w[0] == start || w[0]&on != 0 {
				continue
			}
			l := fastDist[node|w[0]] + 1
			if minute+l > target {
				continue
			}
			paths = append(paths, dfspaths(target, pressure+(target-minute-l)*w[1], minute+l, on|w[0], w[0], path|w[0])...)
		}
		return paths
	}

	allpaths := dfspaths(26, 0, 0, 0, start, 0)

	// reduce paths (presumably, both paths are at least half of part 1)
	var trimpaths [][2]uint16
	for _, p := range allpaths {
		if p[0] > part1/2 {
			trimpaths = append(trimpaths, p)
		}
	}

	// compare all paths to find max
	var max uint16 = 0
	for idx := 0; idx < len(trimpaths); idx += 1 {
		for jdx := idx + 1; jdx < len(trimpaths); jdx += 1 {
			if trimpaths[idx][1]&trimpaths[jdx][1] != 0 {
				continue
			}
			if m := trimpaths[idx][0] + trimpaths[jdx][0]; m > max {
				max = m
			}
		}
	}

	return int(part1), int(max)
}

func parse(input string) []*Valve {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	valves := make([]*Valve, len(lines))
	for idx, line := range lines {
		m := valveregexp.FindStringSubmatch(line)
		i, _ := strconv.Atoi(m[2])
		valves[idx] = &Valve{Name: m[1], Rate: uint16(i), Edges: m[3]}
	}
	return valves
}

func floydWarshall(valves []*Valve) map[*Valve]map[*Valve]uint16 {
	graph := make(map[*Valve]map[*Valve]uint16)
	for _, v1 := range valves {
		graph[v1] = make(map[*Valve]uint16)
		for _, v2 := range valves {
			if v1 == v2 {
				graph[v1][v2] = 0
			} else if strings.Contains(v1.Edges, v2.Name) {
				graph[v1][v2] = 1
			} else {
				graph[v1][v2] = 0xff
			}
		}
	}

	for _, k := range valves {
		for _, i := range valves {
			for _, j := range valves {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}

	return graph
}
