package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	screen tcell.Screen
	state  [24][80]bool
	ticker *time.Ticker
}

func (game Game) display() {
	var style tcell.Style
	game.screen.Clear()
	for y, sliceY := range game.state {
		for x, cell := range sliceY {

			if cell {
				style = tcell.StyleDefault.Background(tcell.ColorBeige)
			} else {
				style = tcell.StyleDefault.Background(tcell.ColorDarkGray)
			}

			game.screen.SetContent(x, y, ' ', nil, style)
		}
	}
	game.screen.Show()
}

func (game *Game) randomState() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(game.state); i++ {
		for j := 0; j < cap(game.state[i]); j++ {
			game.state[i][j] = rand.Float32() < 0.5
		}
	}
}

// func (game *Game) nextGeneration() {

// }
