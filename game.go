package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	screen   tcell.Screen
	width    int
	heigth   int
	universe []bool
	ticker   *time.Ticker
	stop     bool
	event    chan Event
}

func (g Game) display() {
	var style tcell.Style
	g.screen.Clear()
	for index, cell := range g.universe {

		if cell {
			style = tcell.StyleDefault.Background(tcell.GetColor("#403f3f"))
		} else {
			style = tcell.StyleDefault.Background(tcell.ColorBeige)
		}

		g.screen.SetContent(index%g.width, index/g.width, ' ', nil, style)
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
			g.computeNextGeneration()
			g.display()
		}
	}
}

func (g *Game) getIndex(row, col int) int {
	return row*g.width + col
}

func (g *Game) neighborCount(row, col int) int {
	var count int
	for _, deltaRow := range []int{g.heigth - 1, 0, 1} {
		for _, deltaCol := range []int{g.heigth - 1, 0, 1} {
			if deltaCol == 0 && deltaRow == 0 {
				continue
			}

			neighborRow := (row + deltaRow) % g.heigth
			neighborCol := (col + deltaCol) % g.width
			index := g.getIndex(neighborRow, neighborCol)

			if g.universe[index] {
				count++
			}
		}
	}
	return count
}

func (g *Game) computeNextGeneration() {
	nextGeneration := make([]bool, len(g.universe))

	var index int
	var cell bool
	var cellNextState bool
	var aliveNeighbors int

	for row := 0; row < g.heigth; row++ {
		for col := 0; col < g.width; col++ {
			index = g.getIndex(row, col)
			cell = g.universe[index]
			aliveNeighbors = g.neighborCount(row, col)

			switch {
			case cell && aliveNeighbors < 2:
				cellNextState = false
			case cell && (aliveNeighbors == 2 || aliveNeighbors == 3):
				cellNextState = true
			case cell && aliveNeighbors > 3:
				cellNextState = false
			case cell == false && aliveNeighbors == 3:
				cellNextState = true
			}

			nextGeneration[index] = cellNextState
		}
	}
	g.universe = nextGeneration
}

// func (g *Game) nextuniverse() {
// 	for y := 0; i < cap(g.universe); i++ {
// 		for x := 0; j < cap(g.universe[i]); j++ {

// }

func (g *Game) randomuniverse() {
	rand.Seed(time.Now().UnixNano())
	var index int
	for row := 0; row < g.heigth; row++ {
		for col := 0; col < g.width; col++ {
			index = g.getIndex(row, col)
			g.universe[index] = rand.Float32() < 0.5
		}
	}
}
