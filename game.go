package main

import (
	"github.com/gdamore/tcell"
)

type Game struct {
	screen tcell.Screen
	state  [24][80]bool
}

func (g Game) display() {
	var foregroundColor tcell.Style
	for y, sliceY := range g.state {
		for x, cell := range sliceY {

			if cell == true {
				foregroundColor = tcell.StyleDefault.Foreground(tcell.ColorWhite)
			} else {
				foregroundColor = tcell.StyleDefault.Foreground(tcell.ColorBlack)
			}

			g.screen.SetContent(x, y, ' ', nil, foregroundColor)
		}
	}
	g.screen.Show()
}
