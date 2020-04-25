package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func initGame() Game {
	screen := initScreen()
	screenWidth, screenHeigth := 80, 24
	universe := make([]Cell, screenWidth*screenHeigth)
	ticker := time.NewTicker(100 * time.Millisecond)
	stop := false
	event := make(chan Event)

	game := Game{
		screen,
		screenWidth,
		screenHeigth,
		universe,
		ticker,
		stop,
		event,
	}
	return game
}

func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	// screen.EnableMouse()

	return screen
}

func main() {
	game := initGame()
	go inputLoop(&game)
	game.Loop()
}
