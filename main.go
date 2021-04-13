package main

import (
	"image/color"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

type player struct {
	pawn *ebiten.Image
	king *ebiten.Image
}

var (
	rect                   *ebiten.Image
	board                  *ebiten.Image
	redPlayer, blackPlayer player
	x, y                   float64 = 20, 20
	move                   bool
)

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	var err error
	board, _, err = ebitenutil.NewImageFromFile("./assets/board.png")
	missingAsset(err)
}

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("goCheckers")

	rect = ebiten.NewImage(20, 20)
	rect.Fill(color.RGBA{255, 255, 255, 128})

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
