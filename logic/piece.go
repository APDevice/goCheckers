package logic

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

// Piece stores the position of each piece on the board and whether it is a king
type Piece struct {
	x, y      int
	owner     *player
	direction int // stores 1 if pieces move down, else -1
	isKing    bool
}

// Img returns appropriate image file
func (p Piece) Img() *ebiten.Image {
	if p.isKing {
		return p.owner.Img(true)
	}
	return p.owner.Img(false)
}

// renderAt returns a biten.DrawImageOptions object with the proper location
func (p Piece) RenderAt() *ebiten.DrawImageOptions {
	render := ebiten.DrawImageOptions{}
	render.GeoM.Translate(float64(p.x*SQUARESIZE), float64(p.y*SQUARESIZE))

	return &render
}

// move updates the position of the piece if such a move is possible, otherwise returns false
func (p Piece) ValidMove(newX, newY int) bool {
	//TODO write checker to see if move is possible and if a piece will be captured in the process

	return false
}

// SameLoc takes in an x and y position, and returns true if these match what is stored in
// Piece, else returns false
func (p Piece) SameLoc(x, y int) bool {
	return p.x == x && p.y == y
}

// CurrentLoc returns the current x and y positions for Piece struct
func (p Piece) CurrentLoc() (int, int) {
	return p.x, p.y
}

// Moves sets the x and y positions for Piece struct
func (p *Piece) Move(newX, newY int) error {
	if !Board.emptySpace(newX, newY) {
		return errors.New("invalid move")
	}

	oldX, oldY := p.x, p.y
	// TODO insert logic for valid move checking

	p.x, p.y = newX, oldX
	Board.update(p, newX, newY)
	Board.update(nil, oldX, oldY)

	return nil
}
