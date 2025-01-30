package main

import (
	"GoMaze/src/ui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Maze Generator")

	myWindow.SetContent(ui.GetContent(myWindow))
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
