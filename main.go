package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image
var (
	posX = 0.0
	posY = 0.0
	velX = 12.2
	velY = 12.12
)

var bgCol = 1
var prevSpacePressed = false

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("./src/images/SblackBall2.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	posX += velX
	posY += velY

	screenW, screenH := 1920.0, 1080.0
	scale := 0.1

	imgW, imgH := img.Size()
	scaledW := float64(imgW) * scale
	scaledH := float64(imgH) * scale

	if posX < 0 {
		posX = 0
		velX = -velX
	}
	if posX+scaledW > screenW {
		posX = screenW - scaledW
		velX = -velX
	}
	if posY < 0 {
		posY = 0
		velY = -velY
	}
	if posY+scaledH > screenH {
		posY = screenH - scaledH
		velY = -velY
	}

	spacePressed := ebiten.IsKeyPressed(ebiten.KeySpace)
	if spacePressed && !prevSpacePressed {
		switch bgCol {
		case 1:
			bgCol = 2
		case 2:
			bgCol = 3
		case 3:
			bgCol = 1
		}
	}
	prevSpacePressed = spacePressed
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Fatal("User ended game.")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch bgCol {
	case 1:
		screen.Fill(color.RGBA{246, 182, 215, 255})
	case 2:
		screen.Fill(color.RGBA{92, 20, 255, 255})
	case 3:
		screen.Fill(color.RGBA{12, 255, 70, 255})
	}

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(0.1, 0.1)
	options.GeoM.Translate(posX, posY)
	screen.DrawImage(img, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowTitle("BlackBall")
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
