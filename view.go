package main

import "github.com/hajimehoshi/ebiten/v2"

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	bop := &ebiten.DrawImageOptions{}
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(board, bop)
	op.GeoM.Translate(x, y) //position of graphic on screen
	screen.DrawImage(rect, op)
}
