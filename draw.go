package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 55, G: 55, B: 55, A: 255})

	g.Player.Draw(screen)

}
