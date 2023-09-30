package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Character struct {
	life             bool
	position         fyne.Position
	imageRight       *canvas.Image
	imageLeft        *canvas.Image
	imageOrientation string
	size             fyne.Size
	score            int
}

func NewCharacter(imgO string, imgR *canvas.Image, imgL *canvas.Image, posX float32, posY float32) *Character {
	return &Character{
		life:             true,
		imageOrientation: imgO,
		imageRight:       imgR,
		imageLeft:        imgL,
		position:         fyne.NewPos(posX, posY),
		size:             fyne.NewSize(125, 125),
		score:            0,
	}
}

func (c *Character) SetImageRight(imgR *canvas.Image) {
	c.imageRight = imgR
}

func (c *Character) GetImageRight() *canvas.Image {
	return c.imageRight
}

func (c *Character) SetImageLeft(imgL *canvas.Image) {
	c.imageLeft = imgL
}

func (c *Character) GetImageLeft() *canvas.Image {
	return c.imageLeft
}

func (c *Character) SetPosition(position fyne.Position) {
	c.position = position
}

func (c *Character) GetPosition() fyne.Position {
	return c.position
}

func (c *Character) SetLife(life bool) {
	c.life = life
}

func (c *Character) GetLife() bool {
	return c.life
}

func (c *Character) SetImageOrientation(imgO string) {
	c.imageOrientation = imgO
}

func (c *Character) GetImageOrientation() string {
	return c.imageOrientation
}

func (c *Character) SetScore(score int) {
	c.score = score
}

func (c *Character) GetScore() int {
	return c.score
}

func (c *Character) MoveToLeft(container *fyne.Container) {

	if c.GetPosition().X > 475 {
		container.Remove(c.GetImageRight())
		c.SetPosition(fyne.NewPos(475, 590))
		c.GetImageLeft().Resize(c.size)
		c.GetImageLeft().Move(c.GetPosition())
		container.Add(c.GetImageLeft())
		container.Refresh()
	}

}

func (c *Character) MoveToRight(container *fyne.Container) {

	if c.GetPosition().X < 600 {
		container.Remove(c.GetImageLeft())
		c.SetPosition(fyne.NewPos(600, 590))
		c.GetImageRight().Resize(c.size)
		c.GetImageRight().Move(c.GetPosition())
		container.Add(c.GetImageRight())
		container.Refresh()
	}

}
