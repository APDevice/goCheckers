/** board.go contains logic for the rendering and initilization of the board struct */

package checkersLogic

import (
	"errors"
	_ "image/png"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Stores graphics and positional data for board
type board struct {
	Image  *ebiten.Image
	Render ebiten.DrawImageOptions
	Grid   [8][8]*Piece
}

func Board_init() *board {
	return &board{}
}

func (b board) GetImage() *ebiten.Image {
	return b.Image
}

// loads initial assets
func (b *board) Load(img string) (err error) {
	b.Image, _, err = ebitenutil.NewImageFromFile(img)
	if b.Image == nil {
		return errors.New("Image nil as start")
	}
	b.Render = ebiten.DrawImageOptions{}
	missingAsset(err)
	return
}

// Reset the pieces on the board to their inital count and position
func (b board) Reset() {
	// reset number of pieces remaining for each player
	Player1.remaining = 12
	Player2.remaining = 12

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
	var redPiece, blackPiece *Piece
	for y, row := range initlayout {
		for x, char := range row {
			switch char {
			case 'o':
				//initilize red piece if necesary
				if Player1.pieces[redIdx] == nil {
					Player1.pieces[redIdx] = &Piece{
						x:         x,
						y:         y,
						owner:     &Player1,
						direction: -1,
						isKing:    false,
					}
				}
				redPiece = Player1.pieces[redIdx]
				redIdx++

				redPiece.x = x
				redPiece.y = y
				redPiece.isKing = false
				b.Grid[x][y] = redPiece
			case 'x':
				//initilize black piece if necesary
				if Player2.pieces[blackIdx] == nil {
					Player2.pieces[blackIdx] = &Piece{
						x:         x,
						y:         y,
						direction: 1,
						isKing:    false,
					}
				}
				blackPiece = Player2.pieces[blackIdx]
				blackIdx++

				blackPiece.x = x
				blackPiece.y = y
				blackPiece.isKing = false
				b.Grid[x][y] = blackPiece
			default:
				b.Grid[x][y] = nil
			}
		}
	}
	log.Println(b.Grid)
}

func (b board) update(piece *Piece, x, y int) {
	b.Grid[x][y] = piece
}

func (b board) GetPiece(x, y int) *Piece {
	return b.Grid[x][y]
}

// emptySpace returns whether the space is currently occupied
func (b board) emptySpace(x, y int) bool {
	return b.Grid[x][y] == nil
}
