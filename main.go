/*
	Project: goCheckers
	Author: Dylan Luttrell
	License: GNU GPLv3
	Description: An immplementation of the classic game of checkers in the Go langauges, using the ebiten game engine.
				Details for ebiten can be found at ebiten.org.
*/

package goCheckers

import (
	_ "image/png"
	"log"

	//"github.com/pkg/profile"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 800
	SQUARESIZE    = 100
)

// Game implements ebiten.Game interface.
type Game struct {
	// keys []ebiten.Key
}

func main() {
	//defer profile.Start().Stop()
	game := &Game{}
	// configure
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetRunnableOnUnfocused(false)

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("goCheckers")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
