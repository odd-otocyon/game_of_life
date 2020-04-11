package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

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
	screen := initScreen()
	width, height := screen.Size()
	fmt.Printf("Width: %v\nHeight: %v", width, height)
	// screen.Fini()
}
