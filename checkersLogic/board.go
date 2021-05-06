/** board.go contains logic for the rendering and initilization of the board struct */

package checkersLogic

import (
	"log"
)

// Stores graphics and positional data for _board
type _board struct {
	Grid [8][8]*Piece
}

func Board_init() *_board {
	return &_board{}
}

// Reset the pieces on the board to their inital count and position
func (b *_board) Reset(initlayout []string) {
	// reset number of pieces remaining for each player
	players[0].remaining = 12
	players[1].remaining = 12

	var redIdx, blackIdx, allIdx int
	var redPiece, blackPiece *Piece
	for y, row := range initlayout {
		for x, char := range row {
			switch char {
			case 'o':
				redPiece = players[0].pieces[redIdx]
				redIdx++

				redPiece.X = x
				redPiece.Y = y
				redPiece.IsKing = 0
				b.Grid[x][y] = redPiece

				// store in allIdx
				allPieces[allIdx] = redPiece
				allIdx++
			case 'x':
				blackPiece = players[1].pieces[blackIdx]
				blackIdx++

				blackPiece.X = x
				blackPiece.Y = y
				blackPiece.IsKing = 0
				b.Grid[x][y] = blackPiece

				// store in allIdx
				allPieces[allIdx] = blackPiece
				allIdx++
			default:
				b.Grid[x][y] = nil
			}
		}
	}
	log.Println(b.Grid)
}

func (b *_board) update(piece *Piece, x, y int) error {
	b.Grid[x][y] = piece
	return nil
}

func (b _board) GetPiece(x, y int) *Piece {
	return b.Grid[x][y]
}

// emptySpace returns whether the space is currently occupied
func (b _board) emptySpace(x, y int) bool {
	return b.Grid[x][y] == nil
}