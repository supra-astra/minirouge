package main

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

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

func NewGame() *Game {
	g := &Game{}
	return g
}

// update logic will called for each tic.
func (g *Game) Update() error {
	return nil
}

// draw method will be called each draw cycle and
// is where we will split
func (g *Game) Draw(screen *ebiten.Image) {
	//draw the image
	screen.DrawImage(PlayerSprite, nil)
}

// layout will return the screen dimensions.
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
