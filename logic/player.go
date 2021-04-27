package logic

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// player interface is prototype for player

type player struct {
	pawn      *ebiten.Image
	king      *ebiten.Image
	pieces    [12]*Piece
	remaining int
	isAI      bool
}

// throws error if asset is missing
func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (p player) Img(king bool) *ebiten.Image {
	if king {
		return p.king
	}

	return p.pawn
}

func (p player) Pieces() []*Piece {
	return p.pieces[:]
}

func (p player) Draw(screen *ebiten.Image) {
	for _, piece := range p.Pieces() {
		if piece.isKing {
			screen.DrawImage(p.king, piece.RenderAt())
			return
		}
		screen.DrawImage(p.pawn, piece.RenderAt())
	}
}

func player_init(king, pawn string, dir int, ai bool) player {
	var err error
	kingImg, _, err := ebitenutil.NewImageFromFile(king)
	missingAsset(err)
	pawnImg, _, err := ebitenutil.NewImageFromFile(pawn)
	missingAsset(err)

	return player{
		pawn:      pawnImg,
		king:      kingImg,
		pieces:    [12]*Piece{},
		remaining: 12,
		isAI:      ai,
	}
}
