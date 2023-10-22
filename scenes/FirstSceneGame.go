package scenes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"ganego/models"
	"ganego/utilities"
	"math/rand"
)

type FirstSceneGame struct {
	window          fyne.Window
	app             fyne.App
	container       *fyne.Container
	containerStreet *fyne.Container
	containerView   *fyne.Container
	score           *widget.Label
}

func NewFirstSceneGame(w fyne.Window, app fyne.App) *FirstSceneGame {
	return &FirstSceneGame{
		window:          w,
		app:             app,
		container:       container.NewWithoutLayout(),
		containerStreet: container.NewWithoutLayout(),
		containerView:   container.NewWithoutLayout(),
		score:           widget.NewLabel(""),
	}
}

func (f *FirstSceneGame) InitScene() {
	streetRandom := rand.Intn(2) + 0
	street := models.NewStreet()
	street.GetImage()[streetRandom].Resize(fyne.NewSize(street.GetSize().Width, street.GetSize().Height))
	street.GetImage()[streetRandom].Move(street.GetPosition())
	f.containerStreet.Add(street.GetImage()[streetRandom])

	street.GetImage()[0].Resize(street.GetSize())
	street.GetImage()[0].Move(street.GetPosition())

	street.GetImage()[1].Resize(street.GetSize())
	street.GetImage()[1].Move(street.GetPosition())

	street.GetImage()[2].Resize(street.GetSize())
	street.GetImage()[2].Move(street.GetPosition())

	onScreen := street.GetImage()[0]

	//view := models.NewStreet()
	//view.GetImageView()[0].Resize(fyne.NewSize(495, 720))
	//view.GetImageView()[0].Move(fyne.NewPos(725, 0))
	//view.GetImageView()[2].Resize(fyne.NewSize(495, 720))
	//view.GetImageView()[2].Move(fyne.NewPos(0, 0))
	//f.containerView.Add(street.GetImageView()[0])
	//f.containerView.Add(street.GetImageView()[2])

	imgO, imageR, imageL, x, y := utilities.RandomSideAndSkinCharacter()
	character := models.NewCharacter(imgO, imageR, imageL, x, y)
	if character.GetImageOrientation() == "RIGHT" {
		character.GetImageRight().Resize(fyne.NewSize(125, 125))
		character.GetImageRight().Move(character.GetPosition())
		f.container.Add(character.GetImageRight())
	} else {
		character.GetImageLeft().Resize(fyne.NewSize(125, 125))
		character.GetImageLeft().Move(character.GetPosition())
		f.container.Add(character.GetImageLeft())
	}

	side, img, x, y := utilities.RandomSideAndSkinObstacle()
	obstacle := models.NewObstacle(x, y, img, true, side)
	obstacle.GetImage().Resize(obstacle.GetSize())
	obstacle.GetImage().Move(obstacle.GetPosition())
	f.container.Add(obstacle.GetImage())

	score := widget.NewLabel(fmt.Sprintf("Score: %d", character.GetScore()))
	score.TextStyle = fyne.TextStyle{Bold: true, TabWidth: 0}
	score.Resize(fyne.NewSize(200, 400))
	score.Move(fyne.NewPos(200, 50))
	f.score = score
	f.container.Add(f.score)

	f.window.SetContent(container.NewWithoutLayout(f.containerStreet, f.container))

	go street.AddStreetMovement(f.containerStreet, character, obstacle, onScreen)

	//street.AddMoveViews(f.containerStreet, character, obstacle)

	go obstacle.ObstaclesGeneratorAndMove(character, f.container, f.score)

	f.window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if character.GetLife() {
			switch e.Name {
			case fyne.KeyLeft:
				character.MoveToLeft(f.container)
			case fyne.KeyRight:
				character.MoveToRight(f.container)
			default:
				println("same side")
			}
		}
	})

	go ValidateCrash(character, obstacle, f.window, f.app)

}

func ValidateCrash(character *models.Character, obstacle *models.Obstacle, window fyne.Window, app fyne.App) {
	for character.GetLife() && obstacle.GetStatus() {
		if character.GetPosition().X == obstacle.GetPosition().X && obstacle.GetPosition().Y >= 470 {
			obstacle.SetStatus(false)
			character.SetLife(false)
		}
	}
	endSceneGame := NewEndSceneGame(window, app)
	endSceneGame.EndGameScene(character, obstacle)
}
