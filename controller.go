package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	selected *piece
)

//TODO add movement restrictions
//TODO capture enemy pieces
//move logic to model

// mouseInput checks location of mouse cursor
func mouseInput() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		x, y := mouseX/(SCREEN_WIDTH/8), mouseY/(SCREEN_HEIGHT/8)
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
			fmt.Println(float64(x*SQUARESIZE), float64(y*SQUARESIZE))
			selected.renderAt.GeoM.Translate(float64((x-oldX)*SQUARESIZE), float64(y-oldY)*SQUARESIZE)
			selected = nil
			fmt.Println("moved")

		}

	}
}

func (g *Game) Update() error {
	mouseInput()
	return nil
}
