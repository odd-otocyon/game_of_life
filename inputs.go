package main

import (
	"github.com/gdamore/tcell"
)

func inputLoop(game *Game) {
	for {
		eventPoll := game.screen.PollEvent()
		switch eventType := eventPoll.(type) {
		case *tcell.EventKey:
			if eventType.Key() == tcell.KeyEsc {
				game.event <- Event{Type: "done"}
			}
		}
	}
}
