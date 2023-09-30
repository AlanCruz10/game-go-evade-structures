package main

import (
	"fyne.io/fyne/v2"
	"ganego/scenes"
)

func main() {
	mainWindow := scenes.NewMainWindow("EVACIÓN EXTREMA", fyne.NewSize(1240, 720))
	mainWindow.OpenMenu()
}
