package main

import (
	"fmt"
	_ "image/png"
	"log"

	logic "github.com/APDevice/goCheckers/checkersLogic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	BOARD      = "./assets/board.png"
	RED_KING   = "./assets/RKing.png"
	RED_PAWN   = "./assets/RPawn.png"
	BLACK_KING = "./assets/BKing.png"
	BLACK_PAWN = "./assets/BPawn.png"
	SELECTOR   = "./assets/selector.png"
)

// assets
var (
	boardImg *ebiten.Image
	selector *ebiten.Image
	pieces   = [2][2]*ebiten.Image{}
)

func missingAsset(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// load assets
func init() {
	var err error
	boardImg, _, err = ebitenutil.NewImageFromFile(BOARD)
	missingAsset(err)
	selector, _, err = ebitenutil.NewImageFromFile(SELECTOR)
	missingAsset(err)
	pieces[0][0], _, err = ebitenutil.NewImageFromFile(RED_PAWN)
	missingAsset(err)
	pieces[0][1], _, err = ebitenutil.NewImageFromFile(RED_KING)
	missingAsset(err)
	pieces[1][0], _, err = ebitenutil.NewImageFromFile(BLACK_PAWN)
	missingAsset(err)
	pieces[1][1], _, err = ebitenutil.NewImageFromFile(BLACK_KING)
	missingAsset(err)

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	if !update {
		return
	}
	update = false
	// Write your game's rendering.
	//log.Println("start render")
	screen.DrawImage(boardImg, &ebiten.DrawImageOptions{})
	//log.Println("board rendered")
	for _, p := range logic.GetAllPieces() {
		if p.Captured {
			continue
		}
		//log.Printf("piece %v", i)
		drawAt := &ebiten.DrawImageOptions{}
		drawAt.GeoM.Translate(float64(p.X)*SQUARESIZE, float64(p.Y)*SQUARESIZE)
		screen.DrawImage(pieces[p.Owner][p.IsKing], drawAt)
	}
	selectorX, selectorY := logic.GetSelector()

	if selectorX != -1 {
		drawAt := &ebiten.DrawImageOptions{}
		drawAt.GeoM.Translate(float64(selectorX)*SQUARESIZE, float64(selectorY)*SQUARESIZE)
		log.Println("selector at:", float64(selectorX)*SQUARESIZE, float64(selectorY)*SQUARESIZE)
		screen.DrawImage(selector, drawAt)
	}

	ebiten.SetWindowTitle(fmt.Sprintf("goCheckers %v", (ebiten.CurrentFPS())))
}
