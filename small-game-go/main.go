package main

import (
	_ "image/png"
	"small-game/engine"
	"small-game/game"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &game.Game{
		PlayerPosition: game.Vector{X: 100, Y: 100},
		AttackTimer:    engine.NewTimer(5 * time.Second),
		Player:         game.NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}

}
