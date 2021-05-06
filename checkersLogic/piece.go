package checkersLogic

import (
	"errors"
)

// Piece stores the position of each piece on the board and whether it is a king
type Piece struct {
	X, Y      int
	Owner     int
	direction int // stores 1 if pieces move down, else -1
	IsKing    int
}

// Img returns appropriate image file
func (p Piece) Player() int {
	return p.Owner
}

// move updates the position of the piece if such a move is possible, otherwise returns false
func (p Piece) ValidMove(newX, newY int) bool {
	//TODO write checker to see if move is possible and if a piece will be captured in the process
	if p.IsKing == 0 && newY*p.direction < p.Y*p.direction { // only kings can move in reverse
		return false
	}

	return true
}

// SameLoc takes in an x and y position, and returns true if these match what is stored in
// Piece, else returns false
func (p Piece) SameLoc(x, y int) bool {
	return p.X == x && p.Y == y
}

// CurrentLoc returns the current x and y positions for Piece struct
func (p Piece) CurrentLoc() (int, int) {
	return p.X, p.Y
}

// Moves sets the x and y positions for Piece struct
func (p *Piece) Move(newX, newY int) error {
	if !Board.emptySpace(newX, newY) {
		return errors.New("invalid move")
	}

	if !p.ValidMove(newX, newY) {
		return nil
	}
	oldX, oldY := p.X, p.Y
	// TODO insert logic for valid move checking

	p.X, p.Y = newX, newY
	Board.update(p, newX, newY)
	Board.update(nil, oldX, oldY)

	playerTurn ^= 1 // toggle player turn once move is complete
	return nil
}
