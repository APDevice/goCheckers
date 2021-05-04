package checkersLogic

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// player interface is prototype for player

type player struct {
	name      string
	assets    interface{}
	pieces    [12]*Piece
	remaining int
	isAI      bool
}

func (p *player) AddAssets(asset interface{}) {
	p.assets = asset
}

// throws error if asset is missing
func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (p player) Img(king bool) *ebiten.Image {
	if king {
		return p.assets.(pieceAssets).king
	}

	return p.assets.(pieceAssets).pawn
}

func (p player) Pieces() []*Piece {
	return p.pieces[:]
}

func (p player) Draw(screen *ebiten.Image) {
	for _, piece := range p.Pieces() {
		if piece.isKing {
			screen.DrawImage(p.assets.(pieceAssets).pawn, piece.RenderAt())
			return
		}
		screen.DrawImage(p.assets.(pieceAssets).king, piece.RenderAt())
	}
}

func player_init(name, king, pawn string, dir int, ai bool) player {
	var err error
	kingImg, _, err := ebitenutil.NewImageFromFile(king)
	missingAsset(err)
	pawnImg, _, err := ebitenutil.NewImageFromFile(pawn)
	missingAsset(err)

	return player{
		name:      name,
		assets:    pieceAssets{kingImg, pawnImg},
		pieces:    [12]*Piece{},
		remaining: 12,
		isAI:      ai,
	}
}
