package ui

import (
	"GoMaze/src/maze"
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2/canvas"
)

func Draw(maze *maze.Maze, updateChan <-chan [][]int) *canvas.Raster {
	raster := canvas.NewRaster(func(width, height int) image.Image {

		img := image.NewRGBA(image.Rect(0, 0, width, height))

		rows := len(maze.Grid)
		cols := len(maze.Grid[0])
		cellWidth := float64(width) / float64(cols)
		cellHeight := float64(height) / float64(rows)

		for y, row := range maze.Grid {
			for x, cell := range row {
				var c color.Color
				switch cell {
				case 0:
					c = color.White
				case 1:
					c = color.Black
				case 2:
					c = color.RGBA{G: 255, A: 255}
				}

				rect := image.Rect(
					int(float64(x)*cellWidth),
					int(float64(y)*cellHeight),
					int(float64(x+1)*cellWidth),
					int(float64(y+1)*cellHeight),
				)
				draw.Draw(img, rect, &image.Uniform{C: c}, image.Point{}, draw.Src)
			}
		}

		return img
	})

	go func() {
		for updatedMaze := range updateChan {

			maze.Grid = updatedMaze

			raster.Refresh()
		}
	}()

	return raster
}
