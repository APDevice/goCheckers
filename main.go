package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
	squareSize   = 100
)

// Game implements ebiten.Game interface.
type Game struct {
	// keys []ebiten.Key
}

func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	//set initial layout of pieces
	resetPieces()
	ebiten.SetScreenClearedEveryFrame(false)
	//import graphical elements
	//board.Render = ebiten.DrawImageOptions{}
	var err error
	assets := map[string]**ebiten.Image{
		"./assets/board.png": &board.image,
		"./assets/RKing.png": &redPlayer.king,
		"./assets/RPawn.png": &redPlayer.pawn,
		"./assets/BKing.png": &blackPlayer.king,
		"./assets/BPawn.png": &blackPlayer.pawn,
	}

	for key, val := range assets {
		*val, _, err = ebitenutil.NewImageFromFile(key)
		missingAsset(err)
	}

}

func main() {
	game := &Game{}

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("goCheckers")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
