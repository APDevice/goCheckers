/* initilize board */
package checkersLogic

var (
	Board   _board
	players = [2]player{
		player_init(0, -1, false),
		player_init(1, 1, false),
	}
	allPieces  = make([]*Piece, 24)
	selected   *Piece
	playerTurn int
)

func init() {
	//set initial layout of pieces
	Reset()
}
