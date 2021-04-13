package main

//TODO make a box move
import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 340
	screenHeight = 240
)

var (
	rect *ebiten.Image
	x, y float64 = 20, 20
	move bool
)

// Game implements ebiten.Game interface.
type Game struct {
	keys []ebiten.Key
}

func moveBlockside(key ebiten.Key, dir int) {

	if inpututil.KeyPressDuration(key) > 0 {
		x += float64(dir*inpututil.KeyPressDuration(key)) / 30
	}

}

func moveBlockupdown(key ebiten.Key, dir int) {

	if inpututil.KeyPressDuration(key) > 0 {
		y += float64(dir*inpututil.KeyPressDuration(key)) / 30
	}

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	//g.keys = inpututil.PressedKeys()

	moveBlockside(ebiten.KeyD, 1)
	moveBlockside(ebiten.KeyA, -1)

	moveBlockupdown(ebiten.KeyW, -1)
	moveBlockupdown(ebiten.KeyS, 1)

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	op := &ebiten.DrawImageOptions{}

	//op.GeoM.Translate(x, y) //position of graphic on screen

	screen.DrawImage(rect, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("goCheckers")

	rect = ebiten.NewImage(20, 20)
	rect.Fill(color.RGBA{255, 255, 255, 128})

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
