package main

import (
	"log"

	"github.com/APDevice/goCheckers/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	selected *logic.Piece
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
			selected = logic.Board.Grid[x][y]
			if selected != nil {
				log.Printf("selected %v\n", selected.Player())
			}
			return
		}
		if selected != nil && selected.SameLoc(x, y) {
			selected = nil
			log.Println("not selected")
		}

		if selected != nil {
			//oldX, oldY := selected.CurrentLoc() // store original position for transformation
			err := selected.Move(x, y)

			if err != nil {
				log.Println(err)
				return // invalid move
			}

			selected = nil
			log.Println("moved")
		}

	}
}

func (g *Game) Update() error {
	mouseInput()
	return nil
}
