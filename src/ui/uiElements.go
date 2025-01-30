package ui

import (
	"GoMaze/src/maze"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

const (
	defaultWidth  = 21
	defaultHeight = 21
)

var (
	widthEntry, heightEntry     *widget.Entry
	generateButton, startButton *widget.Button
	currentMaze                 *maze.Maze
)

func getLeftMenu(mazeContainer *fyne.Container, myWindow fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()

	widthEntry = widget.NewEntry()
	widthEntry.SetText(fmt.Sprintf("%d", defaultWidth))
	heightEntry = widget.NewEntry()
	heightEntry.SetText(fmt.Sprintf("%d", defaultHeight))

	startButton = widget.NewButton("Start Exploration", func() {
		if currentMaze == nil {
			return
		}

		currentMaze.Reset()
		startButton.Disable()
		generateButton.Disable()

		updateChan := make(chan [][]int)

		go func() {
			ExploreMaze(currentMaze, updateChan)
			startButton.Enable()
			generateButton.Enable()
			close(updateChan)
		}()

		mazeContainer.RemoveAll()
		mazeContainer.Add(Draw(currentMaze, updateChan))
		mazeContainer.Refresh()
	})

	generateButton = widget.NewButton("Generate Maze", func() {
		width, err := strconv.Atoi(widthEntry.Text)
		if err != nil || width < 5 || width > 100 {
			width = defaultWidth
			widthEntry.SetText(fmt.Sprintf("%d", defaultWidth))

			dialog.ShowInformation("Invalid Input", "Width must be between 5 and 100", myWindow)
		}
		height, err := strconv.Atoi(heightEntry.Text)
		if err != nil || height < 5 || height > 100 {
			height = defaultHeight
			heightEntry.SetText(fmt.Sprintf("%d", defaultHeight))

			dialog.ShowInformation("Invalid Input", "Height must be between 5 and 100", myWindow)
		}

		currentMaze = maze.NewMaze(width, height)
		currentMaze.Generate(1, 1)

		updateChan := make(chan [][]int)

		mazeContainer.RemoveAll()
		mazeContainer.Add(Draw(currentMaze, updateChan))
		mazeContainer.Refresh()
	})

	content.Add(widget.NewLabel("Width:"))
	content.Add(widthEntry)
	content.Add(widget.NewLabel("Height:"))
	content.Add(heightEntry)
	content.Add(startButton)
	content.Add(generateButton)

	return content
}

func GetContent(myWindow fyne.Window) *fyne.Container {
	mazeContainer := container.NewMax()

	leftMenu := getLeftMenu(mazeContainer, myWindow)

	content := container.NewBorder(
		nil,
		nil,
		leftMenu,
		nil,
		mazeContainer,
	)

	return content
}
