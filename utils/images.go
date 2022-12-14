package utils

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func RunesToBlocks(runes [][]rune, colors map[rune]color.Color, bHeight int, bWidth int) *image.Paletted {
	rows, cols := GridSizes(runes)
	width, height := cols*bWidth, rows*bHeight
	palette := make(color.Palette, 0, len(colors))

	for _, c := range colors {
		palette = append(palette, c)
	}

	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			row, col := y/bHeight, x/bWidth
			img.Set(x, y, colors[runes[row][col]])
		}
	}

	return img
}

func Animate(frames []*image.Paletted) gif.GIF {
	return gif.GIF{Delay: make([]int, len(frames)), Image: frames}
}

func SaveGif(g gif.GIF, path string) {
	f, err := os.Create(path)

	MaybePanic(err)
	MaybePanic(gif.EncodeAll(f, &g))
}
