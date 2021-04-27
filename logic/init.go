/* initilize assets and board */

package logic

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	BOARD      = "./assets/board.png"
	RED_KING   = "./assets/RKing.png"
	RED_PAWN   = "./assets/RPawn.png"
	BLACK_KING = "./assets/BKing.png"
	BLACK_PAWN = "./assets/BPawn.png"
)

var (
	SQUARESIZE = 100
	Board      board
	Player1    = player_init(RED_KING, RED_PAWN, -1, false)
	Player2    = player_init(BLACK_KING, BLACK_PAWN, 1, true)
)

func init() {
	//set initial layout of pieces
	Board.Reset()
	ebiten.SetScreenClearedEveryFrame(false)

	Board.Load(BOARD)

}
