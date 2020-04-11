package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

const StyleDefault tcell.Style = 0

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
	screen.EnableMouse()

	return screen
}

func main() {

	var content string = "Print that"
	var combc string = "And that"
	combin := []rune(combc)

	screen := initScreen()
	screen.Clear()
	// width, height := screen.Size()

	for x, char := range content {
		screen.SetContent(x, 1, char, combin, StyleDefault)
	}
	screen.Show()
	// screen.Fini()
}
