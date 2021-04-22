package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	selected *piece
)

func mouseInput() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		x, y := mouseX/(screenWidth/8), mouseY/(screenHeight/8)
		if selected == nil {
			selected = board.Grid[x][y]
			fmt.Println("selected")
			return
		}
		if selected != nil && selected.x == x && selected.y == y {
			selected = nil
			fmt.Println("not selected")
		}

		if selected != nil && board.Grid[x][y] == nil {
			oldX, oldY := selected.x, selected.y // store original position for transformation
			selected.x = x
			selected.y = y
			board.Grid[x][y] = selected
			board.Grid[oldX][oldY] = nil
			fmt.Println(float64(x*squareSize), float64(y*squareSize))
			selected.renderAt.GeoM.Translate(float64((x-oldX)*squareSize), float64(y-oldY)*squareSize)
			selected = nil
			fmt.Println("moved")
		}
	}

}

func (g *Game) Update() error {
	// Write your game's logical update.
	//g.keys = inpututil.PressedKeys()

	// moveBlockside(ebiten.KeyD, 1)
	// moveBlockside(ebiten.KeyA, -1)

	// moveBlockupdown(ebiten.KeyW, -1)
	// moveBlockupdown(ebiten.KeyS, 1)
	mouseInput()
	return nil
}
