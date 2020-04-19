package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

func initGame() Game {
	screen := initScreen()
	state := [24][80]bool{{false}}

	game := Game{
		screen,
		state,
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
	game.display()
}
