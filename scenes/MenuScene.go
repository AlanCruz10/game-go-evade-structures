package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window fyne.Window
	app    fyne.App
}

func NewMenuScene(w fyne.Window, a fyne.App) *MenuScene {
	return &MenuScene{
		window: w,
		app:    a,
	}
}

func (m *MenuScene) Menu() {

	title := widget.NewLabel("EVACIÓN EXTREMA")
	title.TextStyle = fyne.TextStyle{Bold: true, TabWidth: 0}
	title.Resize(fyne.NewSize(200, 400))
	title.Move(fyne.NewPos(540, 50))

	start := widget.NewButton("START", m.Start)
	start.Resize(fyne.NewSize(200, 50))
	start.Move(fyne.NewPos(150, 250))

	exit := widget.NewButton("EXIT", m.Exit)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(900, 250))

	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/backgrounds/background2.png"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	m.window.SetContent(container.NewWithoutLayout(backgroundImage, title, start, exit))

}

func (m *MenuScene) Start() {
	sceneGame := NewFirstSceneGame(m.window, m.app)
	sceneGame.InitScene()
}

func (m *MenuScene) Exit() {
	m.window.Close()
}
