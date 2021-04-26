/* initilize assets and board */

package logic

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// SCREEN_WIDTH  = 800
	// SCREEN_HEIGHT = 800
	// SQUARESIZE    = 100
	// assets
	BOARD      = "./assets/board.png"
	RED_KING   = "./assets/RKing.png"
	RED_PAWN   = "./assets/RPawn.png"
	BLACK_KING = "./assets/BKing.png"
	BLACK_PAWN = "./assets/BPawn.png"
)

var (
	board       Board
	redPlayer   ai
	blackPlayer human
)

func init() {
	//set initial layout of pieces
	board.Reset()
	ebiten.SetScreenClearedEveryFrame(false)

	board.Load(BOARD)
	redPlayer.load(RED_KING, RED_PAWN)
	blackPlayer.load(BLACK_KING, BLACK_PAWN)

}
