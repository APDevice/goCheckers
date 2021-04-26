package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// piece stores the position of each piece on the board and whether it is a king
type piece struct {
	x, y      int
	renderAt  ebiten.DrawImageOptions
	direction int // stores 1 if pieces move down, else -1
	isKing    bool
}

func (p piece) renderLoc() *ebiten.DrawImageOptions {
	return &p.renderAt
}

// move updates the position of the piece if such a move is possible, otherwise returns false
func (p piece) move(newX, newY int) bool {
	//TODO write checker to see if move is possible and if a piece will be captured in the process

	return true
}
