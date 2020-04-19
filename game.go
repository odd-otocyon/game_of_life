package main

import (
	"github.com/gdamore/tcell"
)

type Game struct {
	screen tcell.Screen
	state  [24][80]bool
}

func (g Game) display() {
	var style tcell.Style
	for y, sliceY := range g.state {
		for x, cell := range sliceY {

			if cell == true {
				style = tcell.StyleDefault.Background(tcell.ColorBeige)
			} else {
				style = tcell.StyleDefault.Background(tcell.ColorDarkGray)
			}

			g.screen.SetContent(x, y, ' ', nil, style)
		}
	}
	g.screen.Show()
}
