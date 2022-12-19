package main

// VERY heavily inspired by Go: https://github.com/korylprince/adventofcode/blob/master/2022/16/main.go
// See original folder
// Same algorithm adapted to parsing and data used in the WIP.
// Remove the bitmap optimisations for clarity but still runs in 8s.

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const START = "AA"

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

	valves := parse(lines)
	dist := floydWarshall(valves)

	mvps := make(map[string]*Valve)
	for _, valve := range valves {
		if valve.flow > 0 {
			mvps[valve.label] = valve
		}
	}
	mvps[START] = valves[START]

	resultPart1 = part1(mvps, dist)
	resultPart2 = part2(mvps, dist)

	return
}

func part1(mvps map[string]*Valve, dist map[*Valve]map[*Valve]int) int {
	start := mvps[START]
	var dfs func(target int, pressure int, minute int, prevs []*Valve, node *Valve) int
	dfs = func(target int, pressure int, minute int, prevs []*Valve, node *Valve) int {
		max := pressure
		for _, mvp := range mvps {
			if node == mvp || mvp == start || utils.Contains(prevs, mvp) {
				continue
			}
			d := dist[node][mvp] + 1
			if minute+d > target {
				continue
			}
			if next := dfs(target, pressure+(target-minute-d)*mvp.flow, minute+d, append(prevs, mvp), mvp); next > max {
				max = next
			}
		}
		return max
	}

	return dfs(30, 0, 0, make([]*Valve, 0, len(mvps)), start)
}

type PathValue struct {
	pressure int
	visited  []*Valve
}

func (p *PathValue) overlaps(o *PathValue) bool {
	for _, v := range o.visited {
		if utils.Contains(p.visited, v) {
			return true
		}
	}

	return false
}

func part2(mvps map[string]*Valve, dist map[*Valve]map[*Valve]int) int {
	start := mvps[START]

	var dfs func(target int, pressure int, minute int, prevs []*Valve, node *Valve, visited []*Valve) []*PathValue

	dfs = func(target int, pressure int, minute int, prevs []*Valve, node *Valve, visited []*Valve) []*PathValue {
		paths := []*PathValue{{pressure, visited}}
		for _, mvp := range mvps {
			if node == mvp || mvp == start || utils.Contains(prevs, mvp) {
				continue
			}
			d := dist[node][mvp] + 1
			if minute+d > target {
				continue
			}
			paths = append(paths, dfs(target, pressure+(target-minute-d)*mvp.flow, minute+d, append(prevs, mvp), mvp, append(visited, mvp))...)
		}
		return paths
	}

	allpaths := dfs(26, 0, 0, make([]*Valve, 0, len(mvps)), start, make([]*Valve, 0))

	// compare all paths to find max
	var max = 0
	for i := 0; i < len(allpaths); i += 1 {
		for j := i + 1; j < len(allpaths); j += 1 {
			me := allpaths[i]
			elephant := allpaths[j]
			if me.overlaps(elephant) {
				continue
			}
			if m := me.pressure + elephant.pressure; m > max {
				max = m
			}
		}
	}
	return max
}

func parse(lines []string) map[string]*Valve {
	valves := utils.SliceMap(lines, parseValve)
	valveMap := make(map[string]*Valve)

	for _, valve := range valves {
		valveMap[valve.label] = valve
	}

	for i, leads := range utils.SliceMap(lines, parseLeads) {
		for _, lead := range leads {
			to := valveMap[lead]
			valves[i].to = append(valves[i].to, to)
		}
	}

	return valveMap
}
func parseValve(line string) *Valve {
	v := new(Valve)
	fmt.Sscanf(line, "Valve %s has flow rate=%d;", &v.label, &v.flow)

	return v
}
func parseLeads(line string) []string {
	_, line, _ = strings.Cut(line, ";")
	line = strings.ReplaceAll(line, " tunnel leads to valve ", "")
	line = strings.ReplaceAll(line, " tunnels lead to valves ", "")
	line = strings.ReplaceAll(line, ", ", " ")

	return strings.Fields(line)
}

type Valve struct {
	label string
	flow  int
	to    []*Valve
}

func (v Valve) String() string {
	return fmt.Sprintf("%s(%d)", v.label, v.flow)
}

// floydWarshall calculate the min distance between every 2 pair of nodes in a directed graph.
// returns map[from][to] = distance
func floydWarshall(nodes map[string]*Valve) map[*Valve]map[*Valve]int {
	dist := make(map[*Valve]map[*Valve]int)
	for _, v1 := range nodes {
		dist[v1] = make(map[*Valve]int)
		for _, v2 := range nodes {
			if v1 == v2 {
				dist[v1][v2] = 0
			} else if utils.Contains(v1.to, v2) {
				dist[v1][v2] = 1 // Distance always 1 in this case, could be positive or negative
			} else {
				dist[v1][v2] = 0xff // Infinite Distance, large enough, but should not overflow
			}
		}
	}

	for _, k := range nodes {
		for _, i := range nodes {
			for _, j := range nodes {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
