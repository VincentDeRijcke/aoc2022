package utils

import (
	"errors"
	"strconv"
)

// Grid wraps a runes assuming a coordinate system with
//
//		TopLeft (0,0) origin,
//		Top to Down R axis,
//		Left to Right C positive
//	 This is chosen because the usual loop, where the first iterator is the vertical axis
//			for c, row := range runes {
//				for c, eIJ := range row[c] {
//					...
//				}
//			}
//
//	 through rows with c then with c
//		0....mc (c →)
//		......
//		......
//		mr....mc,mr
//		(r ↓)
type Grid[I any] struct {
	grid   [][]I
	r, c   int
	mr, mc int
}

func NewGrid[I any](grid [][]I) Grid[I] {
	r, c := GridSizes(grid)
	return Grid[I]{grid: grid, mc: c, mr: r}
}

func (g *Grid[I]) Move(r int, c int) I {
	g.r = r
	g.c = c

	return g.Current()
}

func (g *Grid[I]) Current() I {
	return g.grid[g.r][g.c]
}

func (g *Grid[I]) Explore() Explorer[I] {
	return Explorer[I]{grid: g, r: g.r, c: g.c}
}

const (
	Up = iota
	Down
	Left
	Right
)

type Explorer[I any] struct {
	grid     *Grid[I]
	r, c     int
	overflow bool
	path     []int
}

func (e *Explorer[I]) Current() I {
	return e.grid.grid[e.r][e.c]
}

func (e *Explorer[I]) Overflow() bool {
	return e.overflow
}

func (e *Explorer[I]) Next(direction int) bool {
	switch direction {
	case Up:
		return e.Up()
	case Down:
		return e.Down()
	case Left:
		return e.Left()
	case Right:
		return e.Right()
	}
	panic(errors.New("invalid direction: " + strconv.Itoa(direction)))
}

func (e *Explorer[I]) Up() bool {
	e.overflow = e.r == 0
	if !e.overflow {
		e.r--
	}
	return !e.overflow
}

func (e *Explorer[I]) Down() bool {
	e.overflow = e.r == e.grid.mr-1
	if !e.overflow {
		e.r++
	}
	return !e.overflow
}

func (e *Explorer[I]) Left() bool {
	e.overflow = e.c == 0
	if !e.overflow {
		e.c--
	}
	return !e.overflow
}

func (e *Explorer[I]) Right() bool {
	e.overflow = e.c == e.grid.mc-1
	if !e.overflow {
		e.c++
	}
	return !e.overflow
}
