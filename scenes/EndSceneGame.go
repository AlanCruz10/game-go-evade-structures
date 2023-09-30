package scenes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"ganego/models"
)

type EndSceneGame struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

func NewEndSceneGame(w fyne.Window, app fyne.App) *EndSceneGame {
	return &EndSceneGame{
		window:    w,
		app:       app,
		container: container.NewWithoutLayout(),
	}
}

func (e *EndSceneGame) EndGameScene(character *models.Character, obstacle *models.Obstacle) {

	backgroundGameOver := canvas.NewImageFromURI(storage.NewFileURI("./assets/backgrounds/gameover.png"))
	backgroundGameOver.Resize(fyne.NewSize(1240, 720))
	backgroundGameOver.Move(fyne.NewPos(0, 75))
	e.container.Add(backgroundGameOver)

	score := widget.NewLabel(fmt.Sprintf("Score: %d", character.GetScore()))
	score.TextStyle = fyne.TextStyle{Bold: true, TabWidth: 0}
	score.Resize(fyne.NewSize(450, 400))
	score.Move(fyne.NewPos(585, 350))
	e.container.Add(score)

	menu := widget.NewButton("Volver al menú", func() {
		e.BackMenu(character, obstacle)
	})
	menu.Resize(fyne.NewSize(200, 50))
	menu.Move(fyne.NewPos(400, 450))
	e.container.Add(menu)

	restart := widget.NewButton("Reintentar", func() {
		e.Restart(character, obstacle)
	})
	restart.Resize(fyne.NewSize(200, 50))
	restart.Move(fyne.NewPos(650, 450))
	e.container.Add(restart)

	exit := widget.NewButton("Salir", e.Exit)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(525, 550))
	e.container.Add(exit)

	e.window.SetContent(e.container)

}

func (e *EndSceneGame) BackMenu(character *models.Character, obstacle *models.Obstacle) {
	character.SetLife(true)
	obstacle.SetStatus(true)
	scene := NewMenuScene(e.window, e.app)
	scene.Menu()
}

func (e *EndSceneGame) Restart(character *models.Character, obstacle *models.Obstacle) {
	character.SetLife(true)
	obstacle.SetStatus(true)
	game := NewFirstSceneGame(e.window, e.app)
	game.InitScene()
}

func (e *EndSceneGame) Exit() {
	e.window.Close()
}
