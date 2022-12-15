package main

import (
	"aoc_go22/utils"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	start := time.Now()
	fmt.Println("Start", start)

	part1, part2 := resolve(input, 2000000)

	fmt.Println("Result Part 1:", part1)
	fmt.Println("Result Part 2:", part2)
	fmt.Println("Done - ", time.Now().Sub(start))
}

func resolve(input string, y int) (resultPart1 int, resultPart2 int) {
	lines := utils.SliceFilter(utils.SliceMap(strings.Split(input, "\n"), strings.TrimSpace), utils.IsNotEmpty)

	sensors := parse(lines)
	existingBeacons := make([]int, 0, len(sensors))
	maxx, maxy := 0, 0
	for _, sensor := range sensors {
		if sensor.beacon.y == y {
			existingBeacons = append(existingBeacons, sensor.beacon.x)
		}
		maxx = utils.Min(utils.Max(sensor.location.x, maxx), 4_000_000)
		maxy = utils.Min(utils.Max(sensor.location.y, maxy), 4_000_000)
	}
	// Part1
	covers, uncovers := coverage(sensors, y)
	for _, existingBeacon := range existingBeacons {
		covers = subtract(covers, existingBeacon)
	}
	for _, cover := range covers {
		resultPart1 += cover.width()
	}
	// Part2
	for y := 0; y <= maxy; y++ {
		covers, uncovers = coverage(sensors, y)
		x := gap(uncovers, maxx)
		if x >= 0 {
			//fmt.Println("Found", Point{x, y})
			return resultPart1, 4_000_000*x + y
		}
	}

	return
}

func parse(lines []string) []Sensor {
	return utils.SliceMap(lines, parseSensor)
}

func coverage(sensors []Sensor, y int) ([]Cover, []Cover) {
	all := make([]Cover, 0, len(sensors))
	for _, sensor := range sensors {
		cover := sensor.cover(y)
		if cover != NILCOVER {
			all = append(all, cover)
		}
	}

	return merge(all)
}

func merge(covers []Cover) ([]Cover, []Cover) {
	uncovered := make([]Cover, 0, len(covers))
	merged := make([]Cover, 0, len(covers))
	sort.Slice(covers, func(i, j int) bool { return covers[i].from < covers[j].from })
	if len(covers) > 0 {
		merged = append(merged, covers[0])
		prev := covers[0]
		for i := 1; i < len(covers); i++ {
			cur := covers[i]
			if prev.to < cur.from { // Distinct 1fr..1to..2fr..2to
				merged = append(merged, cur)
				uncovered = append(uncovered, newCover(prev.to+1, cur.from-1))
				prev = cur
			} else if cur.to > prev.to { // Overlap 1fr..2fr..1to..2to
				cur = newCover(prev.to+1, cur.to)
				merged = append(merged, cur)
				prev = cur
			} // Included 1fr..2fr..2to..1fr
		}
	}

	return merged, uncovered
}
func subtract(covers []Cover, x int) []Cover {
	subtracted := make([]Cover, 0, len(covers))
	sort.Slice(covers, func(i, j int) bool { return covers[i].from < covers[j].from })

	for _, cur := range covers {
		if x == cur.from {
			subtracted = append(subtracted, newCover(cur.from+1, cur.to))
		} else if x == cur.to {
			subtracted = append(subtracted, newCover(cur.from, cur.to-1))
		} else if x > cur.from && x < cur.to {
			subtracted = append(subtracted, newCover(cur.from, x-1), newCover(x+1, cur.to))
		} else {
			subtracted = append(subtracted, cur)
		}
	}

	return subtracted
}

func gap(uncovers []Cover, maxx int) int {
	for _, u := range uncovers {
		if u.from <= maxx && u.to >= 0 {
			return utils.Max(0, u.from)
		}
	}
	//Let's assume the gap is not before covers begins and after covers ends.
	return -1
}

func parseSensor(line string) Sensor {
	coords := utils.Atois(strings.FieldsFunc(line, utils.IsNotInt))
	return Sensor{Point{coords[0], coords[1]}, Point{coords[2], coords[3]}}
}

var NILCOVER = Cover{0, -1}

type Point struct{ x, y int }

func (p Point) dist(o Point) int {
	// eg: S(8,7) -> B(2,10) == 8-2 + 10-7 == 9
	return utils.Diff(p.x, o.x) + utils.Diff(p.y, o.y)
}

type Cover struct{ from, to int }

func (r Cover) width() int {
	if r == NILCOVER {
		return 0
	}
	return r.to - r.from + 1
}

func (r Cover) include(o Cover) bool {
	return o.from >= r.from && o.to <= r.to
}
func newCover(x0, x1 int) Cover {
	return Cover{utils.Min(x0, x1), utils.Max(x0, x1)}
}

type Sensor struct {
	location Point
	beacon   Point
}

func (s Sensor) dist() int {
	return s.location.dist(s.beacon)
}
func (s Sensor) cover(y int) Cover {
	// ####S#### S.dist(B) == 4
	//  #######
	//   B####
	//    #T#    <- Target, S.dist(T) == 3 => xDist == 4 - 3 == 1
	//     #
	max := s.dist()
	actual := s.location.dist(Point{s.location.x, y})
	xDist := max - actual
	if xDist >= 0 {
		return newCover(s.location.x-xDist, s.location.x+xDist)
	}
	return NILCOVER
}
