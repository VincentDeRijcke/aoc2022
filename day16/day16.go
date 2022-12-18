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

	valves := parse(lines)
	mvps := make(map[string]*Valve)
	for _, valve := range valves {
		if valve.flow > 0 {
			mvps[valve.label] = valve
		}
	}

	for _, mvp := range mvps {
		mvp.build(mvps)
	}
	for _, mvp := range mvps {
		mvp.buildMissing(mvps)
	}

	resultPart1 = part1(valves)

	return
}

func part1(mvps map[string]*Valve) int {
	start := mvps["AA"]
	visited := make(map[string]bool)

	for _, next := range start.paths {
		loop all combo
	}

	return 0
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
	v.paths = make(map[string]*Path)

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
	paths map[string]*Path
}
func (valve Valve) String() string {
	return fmt.Sprintf("%s(to:%v):%v\n", valve.label, utils.SliceMap(valve.to, func(v *Valve) string { return v.label }), valve.paths)
}
func (valve *Valve) build(mvps map[string]*Valve) {
	prev := make([]*Valve, 0, 30)
	current := valve
	d := 1
	valve.recurBuild(mvps, current, prev, nil, d)
}
func (valve *Valve) recurBuild(mvps map[string]*Valve, current *Valve, prev []*Valve, partial []*Valve, d int) {
	for _, to := range current.to {
		if mvps[to.label] != nil && to != valve && !contains(prev, to) && valve.paths[to.label] == nil {
			partial := append(make([]*Valve, 0, len(partial)+1), partial...)
			partial = append(partial, to)
			path := &Path{partial, d}
			valve.paths[to.label] = path
		}
	}
	for _, to := range current.to {
		if mvps[to.label] == nil && to != valve && !contains(prev, to) {
			partial := append(make([]*Valve, 0, len(partial)+1), partial...)
			partial = append(partial, to)
			current = to
			prev = append(prev, to)
			valve.recurBuild(mvps, current, prev, partial, d+1)
		}
	}
}

func (valve *Valve) buildMissing(mvps map[string]*Valve) {
	prev := make([]*Valve, 0, 30)
	current := valve
	d := 1
	valve.recurBuild(mvps, current, prev, nil, d)

	todo
}

type Path struct {
	valves []*Valve
	dist   int
}

func (path Path) String() string {
	return strings.Join(utils.SliceMap(path.valves, func(v *Valve) string { return v.label }), "->")
}

func contains(slice []*Valve, e *Valve) bool {
	for _, i := range slice {
		if i == e {
			return true
		}
	}

	return false
}
