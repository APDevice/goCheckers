/* initilize board */
package checkersLogic

var (
	Board   _board
	players = [2]_player{
		player_init(0, -1, false),
		player_init(1, 1, false),
	}
	allPieces  = make([]*_piece, 24)
	selected   *_piece
	playerTurn int
)

func init() {
	//set initial layout of pieces
	Reset()
}
