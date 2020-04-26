package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell"
)

type game struct {
	screen   tcell.Screen
	width    int
	heigth   int
	universe []bool
	ticker   *time.Ticker
	stop     bool
	event    chan Event
}

func (g game) display() {
	style := tcell.StyleDefault.Background(tcell.ColorBeige).Foreground(tcell.GetColor("#403f3f"))
	var symbol rune
	g.screen.Clear()
	for index, cell := range g.universe {

		if cell {
			symbol = '■'
		} else {
			symbol = ' '
		}

		g.screen.SetContent(index%g.width, index/g.width, symbol, nil, style)
	}
	g.screen.Show()
}

func (g *game) loop() {
	g.randomuniverse()
	for !g.stop {
		select {
		case event := <-g.event:
			switch event.Type {
			case "done":
				g.stop = true
			}
		case <-g.ticker.C:
			g.display()
			g.computeNextGeneration()
		}
	}
}

func (g *game) getIndex(row, col int) int {
	return row*g.width + col
}

func (g *game) neighborCount(row, col int) int {
	var count int
	for _, deltaRow := range []int{g.heigth - 1, 0, 1} {
		for _, deltaCol := range []int{g.width - 1, 0, 1} {

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

func (g *game) computeNextGeneration() {
	nextGeneration := make([]bool, len(g.universe))

	var index int
	var cellNextState bool
	var aliveNeighbors int

	for row := 0; row < g.heigth; row++ {
		for col := 0; col < g.width; col++ {
			index = g.getIndex(row, col)
			cell := g.universe[index]
			aliveNeighbors = g.neighborCount(row, col)

			switch {
			case cell && (aliveNeighbors < 2 || aliveNeighbors > 3):
				cellNextState = false
			case cell && (aliveNeighbors == 2 || aliveNeighbors == 3):
				cellNextState = true
			case !cell && aliveNeighbors == 3:
				cellNextState = true
			default:
				cellNextState = false
			}

			nextGeneration[index] = cellNextState
		}
	}
	g.universe = nextGeneration
}

func (g *game) randomuniverse() {
	rand.Seed(time.Now().UnixNano())
	var index int
	for row := 0; row < g.heigth; row++ {
		for col := 0; col < g.width; col++ {
			index = g.getIndex(row, col)
			g.universe[index] = rand.Float32() < 0.5
		}
	}
}

func (g *game) tests() {
	type Vertex struct {
		row int
		col int
	}

	points := []Vertex{{10, 41}, {10, 42}, {10, 43}, {11, 40}, {11, 41}, {11, 42}}

	for _, point := range points {
		index := g.getIndex(point.row, point.col)
		g.universe[index] = true
	}
}
