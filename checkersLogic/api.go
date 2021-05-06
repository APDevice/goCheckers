package checkersLogic

import (
	"errors"
	"log"
)

// Start initizizes game logic
func Start() error {
	// TODO Seperate start and reset logic
	return nil
}

// Reset returns all pieces to their initial position and status
func Reset() {
	Board.Reset([]string{
		"x-x-x-x-",
		"-x-x-x-x",
		"x-x-x-x-",
		"--------",
		"--------",
		"-o-o-o-o",
		"o-o-o-o-",
		"-o-o-o-o",
	})
}

/* Reload resets the board with the board layout provided to it. Board layout consists of an slice of eight strings of exactly eight characters each, with the following symbols used to represent different values:
"x": player 0 pawn,
"X": player 0 king,
"o": player 1 pawn,
"O": player 1 king,
"-": empty space
*/
func Reload(board []string) error {
	Board.Reset(board)
	return nil
}

// GetAllPieces returns the x and y positions of each peice, its owner, and whether it is a king
func GetAllPieces() []*Piece {
	// returns list allPieces. If pieces is captured, will be nil
	return allPieces
}

func GetWinner() (int, error) {

	return -1, errors.New("game still in progress")
}

func GetPiece(x, y int) (bool, error) {
	log.Println(x, y)
	p := Board.Grid[x][y]

	if p != nil {
		if p.Owner != playerTurn {
			return false, errors.New("not this players turn")
		}
		if p != selected {
			selected = p
			log.Println("selected")
			return true, nil
		}
		log.Println("unselected")
		selected = nil
		return false, nil
	}
	if selected != nil {
		selected.Move(x, y)
		selected = nil
	}

	return false, nil
}
