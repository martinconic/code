package game

import (
	"math/rand"
	"small-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	return &Meteor{
		position: Vector{},
		sprite:   sprite,
	}
}
