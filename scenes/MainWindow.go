package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainWindow struct {
	app    fyne.App
	window fyne.Window
}

func NewMainWindow(title string, size fyne.Size) *MainWindow {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow(title)
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)
	mainWindow.Resize(size)
	return &MainWindow{
		app:    mainApp,
		window: mainWindow,
	}
}

func (w *MainWindow) OpenMenu() {
	menuScene := NewMenuScene(w.window, w.app)
	menuScene.Menu()
	w.window.ShowAndRun()
}
