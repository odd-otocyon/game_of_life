package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

const StyleDefault tcell.Style = 0

type State struct {
	screen tcell.Screen
	runes  []rune
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

func render(state *State) {
	state.screen.Clear()
	for x, r := range state.runes {
		state.screen.SetContent(x, 1, r, nil, StyleDefault)
	}
	state.screen.Show()
}

func main() {
	runes := []rune{'Ƿ', 'Ʈ', '˨', 'Ϩ'}
	screen := initScreen()

	state := State{screen, runes}

	render(&state)
	// screen.Fini()
}
