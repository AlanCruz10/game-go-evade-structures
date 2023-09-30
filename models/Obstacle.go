package models

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"ganego/utilities"
	"time"
)

type Obstacle struct {
	size     fyne.Size
	position fyne.Position
	image    *canvas.Image
	status   bool
	side     string
}

func NewObstacle(posX float32, posY float32, img *canvas.Image, status bool, side string) *Obstacle {
	return &Obstacle{
		size:     fyne.NewSize(125, 125),
		position: fyne.NewPos(posX, posY),
		image:    img,
		status:   status,
		side:     side,
	}
}

func (o *Obstacle) SetImage(img *canvas.Image) {
	o.image = img
}

func (o *Obstacle) GetImage() *canvas.Image {
	return o.image
}

func (o *Obstacle) SetPosition(pos fyne.Position) {
	o.position = pos
}

func (o Obstacle) GetPosition() fyne.Position {
	return o.position
}

func (o Obstacle) SetStatus(status bool) {
	o.status = status
}

func (o *Obstacle) GetStatus() bool {
	return o.status
}

func (o *Obstacle) GetSize() fyne.Size {
	return o.size
}

func (o *Obstacle) SetSide(side string) {
	o.side = side
}

func (o *Obstacle) GetSide() string {
	return o.side
}

func (o *Obstacle) ObstaclesGeneratorAndMove(character *Character, container *fyne.Container, score *widget.Label) {
	count := 0
	speed := 50
	move := 10
	for o.status && character.GetLife() {
		o.position.Y += float32(move)
		o.SetPosition(fyne.NewPos(o.position.X, o.position.Y))
		o.image.Move(o.GetPosition())
		if o.position.Y >= 720 {
			container.Remove(o.GetImage())
			side, image, posX, posY := utilities.RandomSideAndSkinObstacle()
			obstacle := NewObstacle(posX, posY, image, true, side)
			obstacle.GetImage().Resize(obstacle.GetSize())
			obstacle.GetImage().Move(obstacle.GetPosition())
			character.SetScore(character.GetScore() + 1)
			score.SetText(fmt.Sprintf("Score: %d", character.GetScore()))
			if speed == 1 {
				speed = 1
			} else {
				speed--
			}
			if move < 15 {
				count++
				if count == 10 {
					move++
					count = 0
				}
			} else {
				move = 15
			}
			container.Add(obstacle.GetImage())
			container.Refresh()
			o.position.X = obstacle.GetPosition().X
			o.image = obstacle.GetImage()
			o.position.Y = obstacle.GetPosition().Y
		}
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}
