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
	stop   bool
	event  chan Event
}

func (g Game) display() {
	var style tcell.Style
	g.screen.Clear()
	for y, sliceY := range g.state {
		for x, cell := range sliceY {

			if cell {
				style = tcell.StyleDefault.Background(tcell.ColorBeige)
			} else {
				style = tcell.StyleDefault.Background(tcell.GetColor("#403f3f"))
			}

			g.screen.SetContent(x, y, ' ', nil, style)
		}
	}
	g.screen.Show()
}

func (g *Game) Loop() {
	g.randomState()
	for g.stop != true {
		select {
		case event := <-g.event:
			switch event.Type {
			case "done":
				g.stop = true
			}
		case <-g.ticker.C:
			g.randomState()
			g.display()
		}
	}
}

func (g *Game) randomState() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(g.state); i++ {
		for j := 0; j < cap(g.state[i]); j++ {
			g.state[i][j] = rand.Float32() < 0.5
		}
	}
}
