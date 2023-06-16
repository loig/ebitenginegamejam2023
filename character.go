package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Character struct {
	Width       int
	Height      int
	PosX        float64
	PosY        float64
	SpeedX      float64
	SpeedY      float64
	OnFloor     bool
	Jumping     bool
	MovingLeft  bool
	MovingRight bool
	MaxSpeedX   float64
	MaxSpeedY   float64
	IncrSpeedX  float64
	IncrSpeedY  float64
}

func InitPlayer() (c Character) {
	c.Width = gPlayerWidth * gUnit
	c.Height = gPlayerHeight * gUnit
	c.PosX = float64(gScreenWidth*gUnit)/2 - float64(c.Width)/2
	c.PosY = 0
	c.MaxSpeedX = 15
	c.IncrSpeedX = 2
	c.MaxSpeedY = 20
	c.IncrSpeedY = 4
	return
}

func (c Character) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(c.PosX), float32(c.PosY-float64(c.Height)), float32(c.Width), float32(c.Height), color.RGBA{R: 255, B: 255, G: 255, A: 255}, false)
}

func (c *Character) Update(actions [ActionNumber]bool) {

	c.OnFloor = c.CheckFloorAndAdjust()

	c.UpdateLeftRightSpeed(actions[ActionMoveLeft], actions[ActionMoveRight])

	c.UpadteUpDownSpeed(actions[ActionJump], actions[ActionImproveJump])

	c.PosX += c.SpeedX
	c.PosY += c.SpeedY
}

func (c *Character) UpadteUpDownSpeed(askJump, askImproveJump bool) {

	if c.OnFloor {
		c.Jumping = false
		c.SpeedY = 0
	}

	if askJump && c.OnFloor {
		c.SpeedY -= c.IncrSpeedY
		c.Jumping = true
		return
	}

	if askImproveJump && c.Jumping && c.SpeedY >= -c.MaxSpeedY {
		c.SpeedY -= c.IncrSpeedY
		return
	}

	c.Jumping = false

	if !c.OnFloor {
		c.SpeedY += gGravity
	}
}

func (c *Character) UpdateLeftRightSpeed(askLeft, askRight bool) {
	if c.SpeedX >= c.MaxSpeedX && askRight {
		c.SpeedX = c.MaxSpeedX
		return
	}

	if c.SpeedX <= -c.MaxSpeedX && askLeft {
		c.SpeedX = -c.MaxSpeedX
		return
	}

	if askLeft && askRight {
		return
	}

	if askRight {
		c.SpeedX += c.IncrSpeedX
		if c.SpeedX > c.MaxSpeedX {
			c.SpeedX = c.MaxSpeedX
		}
		return
	}

	if askLeft {
		c.SpeedX -= c.IncrSpeedX
		if c.SpeedX < -c.MaxSpeedX {
			c.SpeedX = -c.MaxSpeedX
		}
		return
	}

	if c.SpeedX > 0 && c.OnFloor {
		c.SpeedX -= c.IncrSpeedX
		if c.SpeedX < 0 {
			c.SpeedX = 0
		}
		return
	}

	if c.SpeedX < 0 && c.OnFloor {
		c.SpeedX += c.IncrSpeedX
		if c.SpeedX > 0 {
			c.SpeedX = 0
		}
		return
	}
}

func (c *Character) CheckFloorAndAdjust() (onFloor bool) {
	onFloor = c.PosY >= float64(gScreenHeight*gUnit)

	if onFloor {
		c.PosY = float64(gScreenHeight * gUnit)
	}

	return onFloor
}
