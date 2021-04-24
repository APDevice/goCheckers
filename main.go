package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	SQUARESIZE    = 100
	// assets
	BOARD      = "./assets/board.png"
	RED_KING   = "./assets/RKing.png"
	RED_PAWN   = "./assets/RPawn.png"
	BLACK_KING = "./assets/BKing.png"
	BLACK_PAWN = "./assets/BPawn.png"
)

var (
	board       Board
	redPlayer   ai
	blackPlayer human
)

// Game implements ebiten.Game interface.
type Game struct {
	// keys []ebiten.Key
}

func init() {
	//set initial layout of pieces
	board.Reset()
	ebiten.SetScreenClearedEveryFrame(false)

	board.Load(BOARD)
	redPlayer.load(RED_KING, RED_PAWN)
	blackPlayer.load(BLACK_KING, BLACK_PAWN)

}

func main() {
	game := &Game{}

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("goCheckers")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
