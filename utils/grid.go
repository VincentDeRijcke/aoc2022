package utils

import "errors"

// IsGrid is true if size of each row is the same
func IsGrid[T any](maybeGrid [][]T) bool {
	size := -1
	for _, row := range maybeGrid {
		if size == -1 {
			size = len(row)
		} else if len(row) != size {
			return false
		}
	}

	return true
}

func GridSizes[T any](grid [][]T) (rows int, cols int) {
	if !IsGrid(grid) {
		panic(errors.New("not a grid"))
	}
	if len(grid) == 0 {
		return 0, 0
	}
	return len(grid), len(grid[0])
}

func GridMap[I any, O any](grid [][]I, f func(I) O) [][]O {
	if !IsGrid(grid) {
		panic(errors.New("not a grid"))
	}

	result := make([][]O, len(grid))
	for i, _ := range grid {
		result[i] = SliceMap(grid[i], f)
	}

	return result
}

func GridCopy[I any](grid [][]I) [][]I {
	rows, cols := GridSizes(grid)

	result := make([][]I, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]I, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = grid[i][j]
		}
	}

	return result
}

func Transpose[I any](grid [][]I) [][]I {
	if grid == nil {
		return nil
	}
	rows, cols := GridSizes(grid)

	result := make([][]I, cols)
	for j := 0; j < cols; j++ {
		result[j] = make([]I, rows)
		for i := 0; i < rows; i++ {
			result[j][i] = grid[i][j]
		}
	}

	return result
}
