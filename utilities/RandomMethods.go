package utilities

import (
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
)

func RandomSideAndSkinObstacle() (string, *canvas.Image, float32, float32) {
	side := ""
	x := 0
	y := 0
	numRandom := rand.Intn(6) + 1
	numRandomSide := rand.Intn(2) + 1
	urlImageObstacle := fmt.Sprintf("./assets/obstacles/obstacle%d.png", numRandom)
	obstacleImage := canvas.NewImageFromFile(urlImageObstacle)
	if numRandomSide == 1 {
		side = "LEFT"
		x = 475
		y = 0
		return side, obstacleImage, float32(x), float32(y)
	} else {
		side = "RIGHT"
		x = 600
		y = 0
		return side, obstacleImage, float32(x), float32(y)
	}
}

func RandomSideAndSkinCharacter() (string, *canvas.Image, *canvas.Image, float32, float32) {
	side := ""
	x := 0
	y := 0
	numRandom := rand.Intn(3) + 1
	numRandomSide := rand.Intn(2) + 1
	urlImageRight := fmt.Sprintf("./assets/characters/character%d/character%d.png", numRandom, numRandom)
	urlImageLeft := fmt.Sprintf("./assets/characters/character%d/character%d%d.png", numRandom, numRandom, numRandom)
	characterImageR := canvas.NewImageFromFile(urlImageRight)
	characterImageL := canvas.NewImageFromFile(urlImageLeft)
	if numRandomSide == 1 {
		side = "LEFT"
		x = 475
		y = 590
		return side, characterImageR, characterImageL, float32(x), float32(y)
	} else {
		side = "RIGHT"
		x = 600
		y = 590
		return side, characterImageR, characterImageL, float32(x), float32(y)
	}
}
