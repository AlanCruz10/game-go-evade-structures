package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
	"time"
)

type Street struct {
	position   fyne.Position
	image      [3]*canvas.Image
	imageView  [4]*canvas.Image
	indexImage int
	size       fyne.Size
}

func NewStreet() *Street {
	return &Street{
		position: fyne.NewPos(475, 0),
		image: [3]*canvas.Image{
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street1.png")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street2.png")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/street3.png")),
		},
		imageView: [4]*canvas.Image{
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/views11.jpg")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/views12.jpg")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/views22.jpg")),
			canvas.NewImageFromURI(storage.NewFileURI("assets/backgrounds/views23.jpg")),
		},
		indexImage: 0,
		size:       fyne.NewSize(250, 720),
	}
}

func (s *Street) SetImage(img [3]*canvas.Image) {
	s.image = img
}

func (s *Street) GetImage() [3]*canvas.Image {
	return s.image
}

func (s *Street) SetImageView(img [4]*canvas.Image) {
	s.imageView = img
}

func (s *Street) GetImageView() [4]*canvas.Image {
	return s.imageView
}

func (s *Street) GetPosition() fyne.Position {
	return s.position
}

func (s *Street) GetSize() fyne.Size {
	return s.size
}

func (s *Street) AddStreetMovement(container *fyne.Container, character *Character, obstacle *Obstacle, onScreen *canvas.Image) {

	number := 0
	for {
		if character.GetLife() && obstacle.GetStatus() {
			time.Sleep(120 * time.Millisecond)
			container.Remove(onScreen)
			onScreen = s.image[number]
			container.Add(onScreen)
			container.Refresh()
			number++
			if number == 2 {
				number = 0
			}
		}
	}

	//container.Add(onScreen)

}

func (s *Street) AddMoveViews(container *fyne.Container, character *Character, obstacle *Obstacle) {
	s.imageView[0].Resize(fyne.NewSize(495, 720))
	s.imageView[0].Move(fyne.NewPos(725, 0))

	s.imageView[1].Resize(fyne.NewSize(495, 720))
	s.imageView[1].Move(fyne.NewPos(725, 0))

	s.imageView[2].Resize(fyne.NewSize(495, 720))
	s.imageView[2].Move(fyne.NewPos(0, 0))

	s.imageView[3].Resize(fyne.NewSize(495, 720))
	s.imageView[3].Move(fyne.NewPos(0, 0))

	onScreenLeft := s.imageView[0]
	onScreenRight := s.imageView[2]

	numberLeft := 0
	numberRight := 2
	go func() {
		for {
			if character.GetLife() && obstacle.GetStatus() {
				time.Sleep(120 * time.Millisecond)
				container.Remove(onScreenLeft)
				container.Remove(onScreenRight)
				onScreenLeft = s.imageView[numberLeft]
				onScreenRight = s.imageView[numberRight]
				container.Add(onScreenLeft)
				container.Add(onScreenRight)
				container.Refresh()
				numberLeft++
				numberRight++
				if numberLeft == 1 && numberRight == 3 {
					numberLeft = 0
					numberRight = 2
				}
			}
		}
	}()

	container.Add(onScreenLeft)
	container.Add(onScreenRight)

}
