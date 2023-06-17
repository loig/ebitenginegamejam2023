package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 55, G: 55, B: 55, A: 255})

	g.Player.Draw(screen, DrawPositionTopLeft)
	g.BigGhost.Draw(screen, DrawPositionBottomLeft)
	g.SmallGhost.Draw(screen, DrawPositionTopRight)
	g.TallGhost.Draw(screen, DrawPositionBottomRight)

	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.ActualFPS(), ":", ebiten.ActualTPS()))

}
