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
	c.PosY = 0
	return
}

func (c Character) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(c.PosX), float32(c.PosY-float64(c.Height)), float32(c.Width), float32(c.Height), color.RGBA{R: 255, B: 255, G: 255, A: 255}, false)
}

func (c *Character) Update(actions []int) {
	askMove := false
	onFloor := c.CheckFloorAndAdjust()
	canJump := onFloor

	for _, a := range actions {
		switch a {
		case ActionMoveLeft:
			c.SpeedX = -10
			askMove = true
		case ActionMoveRight:
			c.SpeedX = 10
			askMove = true
		case ActionJump:
			if canJump {
				c.SpeedY = -10
			}
		}
	}

	if c.SpeedY >= 0 && onFloor {
		c.SpeedY = 0
	}

	if !onFloor {
		c.SpeedY += gGravity
	}

	if !askMove && onFloor {
		c.SpeedX = 0
	}

	c.PosX += c.SpeedX
	c.PosY += c.SpeedY
}

func (c *Character) CheckFloorAndAdjust() (onFloor bool) {
	onFloor = c.PosY > float64(gScreenHeight*gUnit)

	if onFloor {
		c.PosY = float64(gScreenHeight * gUnit)
	}

	return onFloor
}
