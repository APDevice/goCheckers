package checkersLogic

import (
	"errors"
	"log"
)

// Piece stores the position of each piece on the board and whether it is a king
type Piece struct {
	X, Y      int
	Owner     int
	direction int // stores 1 if pieces move down, else -1
	IsKing    int
	Captured  bool
}

// Img returns appropriate image file
func (p Piece) Player() int {
	return p.Owner
}

func (p Piece) canCapture() bool {
	captureDiag1 := Board.pieceExists(p.X-1, p.Y+p.direction) && !Board.pieceExists(p.X-2, p.Y+(p.direction*2))
	if captureDiag1 {
		log.Println(p.X-1, p.Y+p.direction, "can be captured")
		return true
	}

	captureDiag2 := Board.pieceExists(p.X+1, p.Y+p.direction) && !Board.pieceExists(p.X+2, p.Y+(p.direction*2))
	if captureDiag2 {
		log.Println(p.X+1, p.Y+p.direction, "can be captured")
		return true
	}

	if p.IsKing == 1 { //additional checks for king
		captureDiag3 := Board.pieceExists(p.X-1, p.Y-p.direction) && !Board.pieceExists(p.X-2, p.Y-(p.direction*2))
		if captureDiag3 {
			log.Println(p.X-1, p.Y-p.direction, "can be captured")
			return true
		}

		captureDiag4 := Board.pieceExists(p.X+1, p.Y-p.direction) && !Board.pieceExists(p.X+2, p.Y-(p.direction*2))
		if captureDiag4 {
			log.Println(p.X+1, p.Y-p.direction, "can be captured")
			return true
		}
	}

	return false
}

func (p Piece) Capture(x, y int) bool {
	log.Println("capture at: ", x, y)
	capturedPiece, err := Board.GetPiece(x, y)

	if err != nil {
		log.Println(err, x, y)
	}
	if capturedPiece == nil {
		return false
	}
	capturedPiece.Captured = true
	Board.update(nil, x, y)

	return true
}

// SameLoc takes in an x and y position, and returns true if these match what is stored in
// Piece, else returns false
func (p *Piece) _move(newX, newY int) {
	oldX, oldY := p.X, p.Y
	p.X, p.Y = newX, newY
	Board.update(nil, oldX, oldY)
	Board.update(p, newX, newY)
}

func (p Piece) SameLoc(x, y int) bool {
	return p.X == x && p.Y == y
}

// CurrentLoc returns the current x and y positions for Piece struct
func (p Piece) CurrentLoc() (int, int) {
	return p.X, p.Y
}

// Moves sets the x and y positions for Piece struct
func (p *Piece) Move(newX, newY int) error {
	err := errors.New("invalid move")
	if !Board.emptySpace(newX, newY) || // piece must move onto empty space
		p.IsKing == 0 && newY*p.direction < p.Y*p.direction || // only kings can move in reverse
		abs(p.X-newX) != abs(p.Y-newY) { // pieces can only move on exact diagonals
		return err
	}
	// TODO insert logic for valid move checking
	switch abs(p.X - newX) {
	case 2:
		c := p.Capture((p.X+newX)/2, (p.Y+newY)/2)
		if !c {
			return err
		}
		p._move(newX, newY)
		if p.canCapture() {
			return nil
		}
	case 1:
		p._move(newX, newY)
	default:
		return err
	}

	playerTurn ^= 1

	// oldX, oldY := p.X, p.Y
	// p.X, p.Y = newX, newY
	// Board.update(nil, oldX, oldY)
	// Board.update(p, newX, newY)

	// toggle player turn once move is complete
	return nil
}
