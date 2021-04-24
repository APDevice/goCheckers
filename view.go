package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.DrawImage(board.image, &board.Render)

	for _, piece := range redPlayer.pieces {
		if piece.isKing {
			screen.DrawImage(redPlayer.king, piece.renderLoc())
		} else {
			screen.DrawImage(redPlayer.pawn, piece.renderLoc())
		}
	}

	for _, piece := range blackPlayer.pieces {
		if piece.isKing {
			screen.DrawImage(blackPlayer.king, piece.renderLoc())
		} else {
			screen.DrawImage(blackPlayer.pawn, piece.renderLoc())
		}
	}

}
