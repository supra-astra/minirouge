package main

import "github.com/hajimehoshi/ebiten/v2"

import "embed"
import "image"
import _ "image/png"

//go:embed assets/*
var assets embed.FS

//loads the assets whole directory

var PlayerSprite = mustLoadImage("assets/PNG/playerShip1_blue.png")

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type Game struct{}

// update logic for the game
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//draw the image
	screen.DrawImage(PlayerSprite, nil)
}

// render the basic ui
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}
