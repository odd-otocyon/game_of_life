package main

import (
	"math/rand"
	"time"
)

func fillRandomState(game *Game) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(game.state); i++ {
		for j := 0; j < cap(game.state[i]); j++ {
			game.state[i][j] = rand.Float32() < 0.5
		}
	}
}
