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

type Blueprint struct{ index, oreOre, clayOre, obsOre, obsClay, geodeOre, geodeObs int }

func (bp *Blueprint) buildOre(c *Configuration, r *Resource) (*Configuration, *Resource) {
	if bp.oreOre <= r.ore && c.ore < bp.maxOre() {
		return &Configuration{c.ore + 1, c.clay, c.obs, c.geode},
			&Resource{r.ore - bp.oreOre, r.clay, r.obs, r.geode}
	}

	return nil, nil
}
func (bp *Blueprint) maxOre() int {
	return utils.Max(bp.oreOre, bp.clayOre, bp.obsOre, bp.geodeOre)
}

func (bp *Blueprint) buildClay(c *Configuration, r *Resource) (*Configuration, *Resource) {
	if bp.clayOre <= r.ore && c.clay < bp.maxClay() {
		return &Configuration{c.ore, c.clay + 1, c.obs, c.geode},
			&Resource{r.ore - bp.clayOre, r.clay, r.obs, r.geode}
	}

	return nil, nil
}
func (bp *Blueprint) maxClay() int {
	return bp.obsClay
}
func (bp *Blueprint) buildObs(c *Configuration, r *Resource) (*Configuration, *Resource) {
	if bp.obsOre <= r.ore && bp.obsClay <= r.clay && c.obs < bp.maxObs() {
		return &Configuration{c.ore, c.clay, c.obs + 1, c.geode},
			&Resource{r.ore - bp.obsOre, r.clay - bp.obsClay, r.obs, r.geode}
	}

	return nil, nil
}
func (bp *Blueprint) maxObs() int {
	return bp.geodeObs
}
func (bp *Blueprint) buildGeode(c *Configuration, r *Resource) (*Configuration, *Resource) {
	if bp.geodeOre <= r.ore && bp.geodeObs <= r.obs {
		return &Configuration{c.ore, c.clay, c.obs, c.geode + 1},
			&Resource{r.ore - bp.geodeOre, r.clay, r.obs - bp.geodeObs, r.geode}
	}

	return nil, nil
}

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

func solveBp(bp *Blueprint) int {
	max := move(1, &Configuration{1, 0, 0, 0}, &Resource{0, 0, 0, 0}, bp)
	fmt.Println("Result for ", *bp, " is ", max)
	return max
}
func move(time int, conf *Configuration, res *Resource, bp *Blueprint) int {
	if time == 24 {
		return conf.tick(res).geode
	}
	max := 0
	if c, r := bp.buildOre(conf, res); c != nil {
		max = utils.Max(max, move(time+1, c, conf.tick(r), bp))
	}
	if c, r := bp.buildClay(conf, res); c != nil {
		max = utils.Max(max, move(time+1, c, conf.tick(r), bp))
	}
	if c, r := bp.buildObs(conf, res); c != nil {
		max = utils.Max(max, move(time+1, c, conf.tick(r), bp))
	}
	if c, r := bp.buildGeode(conf, res); c != nil {
		max = utils.Max(max, move(time+1, c, conf.tick(r), bp))
	}
	max = utils.Max(max, move(time+1, conf, conf.tick(res), bp))

	return max
}
