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
	bps := utils.SliceMap(lines, parse)

	resultPart1 = part1(bps)

	return
}

type Blueprint struct{ index, oreOre, clayOre, obsOre, obsClay, geodeOre, geodeObs, maxOre, maxClay, maxObs int }
type Configuration struct{ ore, clay, obs, geode int }

func (c Configuration) String() string {
	return fmt.Sprintf("(o:%d, c:%d, s:%d, g:%d)", c.ore, c.clay, c.obs, c.geode)
}
func (s *Configuration) tick(res *Resource) *Resource {
	return &Resource{res.ore + s.ore, res.clay + s.clay, res.obs + s.obs, res.geode + s.geode}
}

type Resource struct{ ore, clay, obs, geode int }

func (r Resource) String() string {
	return fmt.Sprintf("(o:%d, c:%d, s:%d, g:%d)", r.ore, r.clay, r.obs, r.geode)
}

func parse(line string) *Blueprint {
	bp := Blueprint{}

	fmt.Sscanf(line, "Blueprint %d: "+
		"Each ore robot costs %d ore. "+
		"Each clay robot costs %d ore. "+
		"Each obsidian robot costs %d ore and %d clay. "+
		"Each geode robot costs %d ore and %d obsidian.", &bp.index, &bp.oreOre, &bp.clayOre, &bp.obsOre, &bp.obsClay, &bp.geodeOre, &bp.geodeObs)

	bp.maxOre = utils.Max(bp.oreOre, bp.clayOre, bp.obsOre, bp.geodeOre)
	bp.maxClay = bp.obsClay
	bp.maxObs = bp.geodeObs

	return &bp
}

func part1(bps []*Blueprint) int {
	total := 0
	start := time.Now()
	for i, bp := range bps {
		total += bp.index * solveBp(bp)
		bpTime := int64(time.Now().Sub(start)) / int64(i+1)
		remaining := bpTime * int64(len(bps)-i-1)
		fmt.Println("Remaining: ", time.Duration(remaining))
	}

	return total
}

var iterations, trees = 0, 0

func solveBp(bp *Blueprint) int {
	iterations, trees = 0, 0
	max := move(24, &Configuration{1, 0, 0, 0}, &Resource{0, 0, 0, 0}, bp)
	fmt.Println("Result for ", *bp, " is ", max, "t:", trees, "i:", iterations)
	return max
}
func move(time int, conf *Configuration, res *Resource, bp *Blueprint) int {
	iterations++
	if time == 1 {
		trees++
		return conf.tick(res).geode
	}
	max := 0
	if bp.oreOre <= res.ore && conf.ore < bp.maxOre {
		res := conf.tick(res)
		res.ore -= bp.oreOre
		max = utils.Max(move(time-1, &Configuration{conf.ore + 1, conf.clay, conf.obs, conf.geode}, res, bp))
	}
	if bp.clayOre <= res.ore && conf.clay < bp.maxClay {
		res := conf.tick(res)
		res.ore -= bp.clayOre
		max = utils.Max(max, move(time-1, &Configuration{conf.ore, conf.clay + 1, conf.obs, conf.geode}, res, bp))
	}
	if bp.obsOre <= res.ore && bp.obsClay <= res.clay && conf.obs < bp.maxObs {
		res := conf.tick(res)
		res.ore -= bp.obsOre
		res.clay -= bp.obsClay
		max = utils.Max(max, move(time-1, &Configuration{conf.ore, conf.clay, conf.obs + 1, conf.geode}, res, bp))
	}
	if bp.geodeOre <= res.ore && bp.geodeObs <= res.obs {
		res := conf.tick(res)
		res.ore -= bp.geodeOre
		res.obs -= bp.geodeObs
		max = utils.Max(max, move(time-1, &Configuration{conf.ore, conf.clay, conf.obs, conf.geode + 1}, res, bp))
	}

	max = utils.Max(max, move(time-1, conf, conf.tick(res), bp))

	return max
}
