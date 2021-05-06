package main

import (
	"log"

	"github.com/APDevice/goCheckers/checkersLogic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	selected bool
	update   bool = true
)

// mouseInput checks location of mouse cursor
func mouseInput() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		x, y := mouseX/(SCREEN_WIDTH/8), mouseY/(SCREEN_HEIGHT/8)
		_, err := checkersLogic.GetPiece(x, y)

		if err != nil {
			log.Println(err)
		}
		update = true
		// 	if selected != nil {
		// 		//oldX, oldY := selected.CurrentLoc() // store original position for transformation
		// 		err := selected.Move(x, y)

		// 		if err != nil {
		// 			log.Println(err)
		// 			return // invalid move
		// 		}

		// 		selected = nil
		// 		log.Println("moved")
		// 	}

	}
}

func (g *Game) Update() error {
	mouseInput()
	return nil
}
