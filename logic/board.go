/** board.go contains logic for the rendering and initilization of the board struct */

package logic

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Stores graphics and positional data for board
type Board struct {
	image  *ebiten.Image
	Render ebiten.DrawImageOptions
	Grid   [8][8]*piece
}

// loads initial assets
func (b Board) Load(img string) {
	var err error
	b.image, _, err = ebitenutil.NewImageFromFile(img)
	missingAsset(err)
}

// Reset the pieces on the board to their inital count and position
func (b Board) Reset() {
	// reset number of pieces remaining for each player
	redPlayer.remaining = 12
	blackPlayer.remaining = 12

	initlayout := [8]string{
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
						x:         0,
						y:         0,
						renderAt:  ebiten.DrawImageOptions{},
						direction: -1,
						isKing:    false,
					}
				}
				redPiece = redPlayer.pieces[redIdx]
				redIdx++

				redPiece.x = x
				redPiece.y = y
				redPiece.renderAt.GeoM.Translate(float64(x*SQUARESIZE), float64(y*SQUARESIZE))
				redPiece.isKing = false
				b.Grid[x][y] = redPiece
			case 'x':
				//initilize black piece if necesary
				if blackPlayer.pieces[blackIdx] == nil {
					blackPlayer.pieces[blackIdx] = &piece{
						x:         0,
						y:         0,
						renderAt:  ebiten.DrawImageOptions{},
						direction: 1,
						isKing:    false,
					}
				}
				blackPiece = blackPlayer.pieces[blackIdx]
				blackIdx++

				blackPiece.x = x
				blackPiece.y = y
				blackPiece.renderAt.GeoM.Translate(float64(x*SQUARESIZE), float64(y*SQUARESIZE))
				blackPiece.isKing = false
				b.Grid[x][y] = blackPiece
			default:
				b.Grid[x][y] = nil
			}
		}
	}
}
