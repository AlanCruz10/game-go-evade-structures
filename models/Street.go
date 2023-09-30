package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Street struct {
	position fyne.Position
	image    [3]*canvas.Image
	size     fyne.Size
}

func NewStreet() *Street {
	return &Street{
		position: fyne.NewPos(475, 0),
		image: [3]*canvas.Image{
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street1.png")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street2.png")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street3.png")),
		},
		size: fyne.NewSize(250, 720),
	}
}

func (s *Street) SetImage(img [3]*canvas.Image) {
	s.image = img
}

func (s *Street) GetImage() [3]*canvas.Image {
	return s.image
}

func (s *Street) GetPosition() fyne.Position {
	return s.position
}

func (s *Street) GetSize() fyne.Size {
	return s.size
}

func (s *Street) AddStreetMovement(container *fyne.Container, InitialStreetImage int, character *Character, obstacle *Obstacle) {

	go func() {
		for character.GetLife() && obstacle.GetStatus() {
			container.Remove(s.image[InitialStreetImage])
			s.image[InitialStreetImage+1].Resize(s.GetSize())
			s.image[InitialStreetImage+1].Move(s.GetPosition())
			container.Add(s.image[InitialStreetImage+1])
			container.Refresh()
		}
	}()

}
