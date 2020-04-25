package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	screen       tcell.Screen
	screenWidth  int
	screenHeigth int
	universe     []Cell
	ticker       *time.Ticker
	stop         bool
	event        chan Event
}

func (g Game) display() {
	var style tcell.Style
	g.screen.Clear()
	for index, cell := range g.universe {

		if cell.alive {
			style = tcell.StyleDefault.Background(tcell.GetColor("#403f3f"))
		} else {
			style = tcell.StyleDefault.Background(tcell.ColorBeige)
		}

		g.screen.SetContent(index%g.screenWidth, index/g.screenWidth, ' ', nil, style)
	}
	g.screen.Show()
}

func (g *Game) Loop() {
	g.randomuniverse()
	for g.stop != true {
		select {
		case event := <-g.event:
			switch event.Type {
			case "done":
				g.stop = true
			}
		case <-g.ticker.C:
			g.randomuniverse()
			g.display()
		}
	}
}

// func (g *Game) nextuniverse() {
// 	for y := 0; i < cap(g.universe); i++ {
// 		for x := 0; j < cap(g.universe[i]); j++ {

// }

func (g *Game) randomuniverse() {
	rand.Seed(time.Now().UnixNano())
	var index int
	for row := 0; row < g.screenHeigth; row++ {
		for column := 0; column < g.screenWidth; column++ {
			index = row*g.screenWidth + column
			// fmt.Println(index)
			g.universe[index].alive = rand.Float32() < 0.5
		}
	}
}
