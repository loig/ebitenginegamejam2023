package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Character struct {
	Width  int
	Height int
	PosX   float64
	PosY   float64
	SpeedX float64
	SpeedY float64
}

func InitPlayer() (c Character) {
	c.Width = gPlayerWidth * gUnit
	c.Height = gPlayerHeight * gUnit
	c.PosX = float64(gScreenWidth*gUnit)/2 - float64(c.Width)/2
	c.PosY = float64(gScreenHeight*gUnit - c.Height)
	return
}

func (c Character) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(c.PosX), float32(c.PosY), float32(c.Width), float32(c.Height), color.RGBA{R: 155, B: 50, A: 255}, false)
}

func (c *Character) Update(actions []int) {
	askMove := false

	for _, a := range actions {
		switch a {
		case ActionMoveLeft:
			c.SpeedX = -10
			askMove = true
		case ActionMoveRight:
			c.SpeedX = 10
			askMove = true
		}
	}

	if !askMove {
		c.SpeedX = 0
	}

	c.PosX += c.SpeedX
}
