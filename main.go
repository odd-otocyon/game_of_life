package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func initGame() Game {
	screen := initScreen()
	state := [24][80]bool{{false}}
	ticker := time.NewTicker(100 * time.Millisecond)
	stop := false
	event := make(chan Event)

	game := Game{
		screen,
		state,
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
