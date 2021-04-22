package main

import "github.com/hajimehoshi/ebiten/v2"

func validMove() {

}

var (
	board                  _board
	redPlayer, blackPlayer player
)

type piece struct {
	x, y     int
	renderAt *ebiten.DrawImageOptions
	isKing   bool
}

type player struct {
	pawn   *ebiten.Image
	king   *ebiten.Image
	pieces [12]*piece
}

type _board struct {
	image  *ebiten.Image
	Render ebiten.DrawImageOptions
	Grid   [8][8]*piece
}

//resetPieces resets all pieces on the board to their inital position
func resetPieces() {
	initlayout := []string{
		"x-x-x-x-",
		"-x-x-x-x",
		"x-x-x-x-",
		"--------",
		"--------",
		"-o-o-o-o",
		"o-o-o-o-",
		"-o-o-o-o",
	}
	var redIdx, blackIdx int
	var redPiece, blackPiece *piece
	for y, row := range initlayout {
		for x, char := range row {
			switch char {
			case 'o':
				//initilize red piece if necesary
				if redPlayer.pieces[redIdx] == nil {
					redPlayer.pieces[redIdx] = &piece{
						x:        0,
						y:        0,
						renderAt: &ebiten.DrawImageOptions{},
						isKing:   false,
					}
				}
				redPiece = redPlayer.pieces[redIdx]
				redIdx++

				redPiece.x = x
				redPiece.y = y
				redPiece.renderAt.GeoM.Translate(float64(x*squareSize), float64(y*squareSize))
				redPiece.isKing = false
				board.Grid[x][y] = redPiece
			case 'x':
				//initilize black piece if necesary
				if blackPlayer.pieces[blackIdx] == nil {
					blackPlayer.pieces[blackIdx] = &piece{
						x:        0,
						y:        0,
						renderAt: &ebiten.DrawImageOptions{},
						isKing:   false,
					}
				}
				blackPiece = blackPlayer.pieces[blackIdx]
				blackIdx++

				blackPiece.x = x
				blackPiece.y = y
				blackPiece.renderAt.GeoM.Translate(float64(x*squareSize), float64(y*squareSize))
				blackPiece.isKing = false
				board.Grid[x][y] = blackPiece
			default:
				board.Grid[x][y] = nil
			}
		}
	}
}
