package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func moveBlockside(key ebiten.Key, dir int) {

	if inpututil.KeyPressDuration(key) > 0 {
		x += float64(dir*inpututil.KeyPressDuration(key)) / 30
	}

}

func moveBlockupdown(key ebiten.Key, dir int) {

	if inpututil.KeyPressDuration(key) > 0 {
		y += float64(dir*inpututil.KeyPressDuration(key)) / 30
	}

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	//g.keys = inpututil.PressedKeys()

	moveBlockside(ebiten.KeyD, 1)
	moveBlockside(ebiten.KeyA, -1)

	moveBlockupdown(ebiten.KeyW, -1)
	moveBlockupdown(ebiten.KeyS, 1)

	return nil
}
