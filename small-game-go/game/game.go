package game

import (
	"small-game/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	PlayerPosition Vector
	AttackTimer    *engine.Timer
	Player         *Player
	Meteors        []*Meteor
}

func (g *Game) Update() error {
	//speed := float64(300 / ebiten.TPS())

	// speed := 5.0
	// var delta Vector

	// if ebiten.IsKeyPressed(ebiten.KeyDown) {
	// 	delta.Y = speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyUp) {
	// 	delta.Y = -speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	// 	delta.X = -speed
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
	// 	delta.X = speed
	// }

	// if delta.X != 0 && delta.Y != 0 {
	// 	factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
	// 	delta.X *= factor
	// 	delta.Y *= factor
	// }

	// g.PlayerPosition.X += delta.X
	// g.PlayerPosition.Y += delta.Y

	g.Player.Update()

	g.AttackTimer.Update()
	if g.AttackTimer.IsReady() {
		g.AttackTimer.Reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
	// op := &ebiten.DrawImageOptions{}
	//op := &colorm.DrawImageOptions{}
	//cm := colorm.ColorM{}
	//cm.Scale(1.0, 1.0, 1.0, 0.5)
	// cm.Translate(1.0, 1.0, 1.0, 0.0)
	// op.GeoM.Translate(g.PlayerPosition.X, g.PlayerPosition.Y)
	// screen.DrawImage(assets.PlayerSprite, op)
	//colorm.DrawImage(screen, PlayerSprite, cm, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
