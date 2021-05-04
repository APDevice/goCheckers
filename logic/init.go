/* initilize assets and board */

package checkersLogic

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
	Player1    = player_init("red", RED_KING, RED_PAWN, -1, false)
	Player2    = player_init("black", BLACK_KING, BLACK_PAWN, 1, true)
)

type pieceAssets struct {
	king *ebiten.Image
	pawn *ebiten.Image
}

func init() {
	//set initial layout of pieces
	Board.Reset()
	ebiten.SetScreenClearedEveryFrame(false)

	Board.Load(BOARD)
	// kingImg, _, _ := ebitenutil.NewImageFromFile(RED_KING)
	// pawnImg, _, _ := ebitenutil.NewImageFromFile(RED_PAWN)
	// a := pieceAssets{kingImg, pawnImg}

	// Player1.AddAssets(a)

	// kingImg, _, _ = ebitenutil.NewImageFromFile(BLACK_KING)
	// pawnImg, _, _ = ebitenutil.NewImageFromFile(BLACK_KING)
	// b := pieceAssets{kingImg, pawnImg}
	// Player2.AddAssets(b)

	// Player2.AddAssets(pawnImg)
}
