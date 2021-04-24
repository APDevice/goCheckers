package main

import (
	"log"

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
						renderAt:  &ebiten.DrawImageOptions{},
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
						renderAt:  &ebiten.DrawImageOptions{},
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

// player interface is prototype for player
type player struct {
	pawn      *ebiten.Image
	king      *ebiten.Image
	pieces    [12]*piece
	remaining int
}

type human struct {
	player
}

type ai struct {
	player
}

// throws error if asset is missing
func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// loads the initial assets for each player
func (p player) load(king, pawn string) {
	var err error
	p.king, _, err = ebitenutil.NewImageFromFile(king)
	missingAsset(err)
	p.pawn, _, err = ebitenutil.NewImageFromFile(pawn)
	missingAsset(err)
}

// piece stores the position of each piece on the board and whether it is a king
type piece struct {
	x, y      int
	renderAt  ebiten.DrawImageOptions
	direction int // stores 1 if pieces move down, else -1
	isKing    bool
}

func (p piece) renderLoc() *ebiten.DrawImageOptions {
	return &p.renderAt
}

// move updates the position of the piece if such a move is possible, otherwise returns false
func (p piece) move(newX, newY int) bool {

	return true
}
