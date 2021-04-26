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
