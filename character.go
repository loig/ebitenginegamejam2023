package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Character struct {
	Width          int
	Height         int
	PosY           float64
	SpeedY         float64
	MaxSpeedY      float64
	IncrSpeedY     float64
	OnFloor        bool
	Jumping        bool
	Sliding        bool
	SlideDuration  int
	SlideFrame     int
	SlidingWidth   int
	SlidingHeight  int
	StandingWidth  int
	StandingHeight int
	Kicking        bool
	Color          color.Color
}

func InitPlayer() (c Character) {
	c.StandingWidth = gPlayerWidth
	c.StandingHeight = gPlayerHeight
	c.SlidingWidth = gPlayerSlidingWidth
	c.SlidingHeight = gPlayerSlidingHeight
	c.Width = c.StandingWidth
	c.Height = c.StandingHeight
	c.PosY = 0
	c.IncrSpeedY = gPlayerIncrSpeedY
	c.MaxSpeedY = gPlayerMaxSpeedY
	c.SlideDuration = gPlayerSlideDuration
	return
}

func (c Character) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(gScreenWidth-c.Width)/2, float32(c.PosY-float64(c.Height)), float32(c.Width), float32(c.Height), c.Color, false)
}

func (c *Character) Update(actions [ActionNumber]bool) {

	c.CheckFloorAndAdjust()
	c.HandleJump(actions[ActionJump], actions[ActionImproveJump])
	c.HandleSlide(actions[ActionSlide])
	c.HandleKick(actions[ActionKick])
	c.SetImage()

}

func (c *Character) HandleKick(askKick bool) {

	c.Kicking = c.Sliding || (!c.OnFloor && askKick)

}

func (c *Character) HandleSlide(askSlide bool) {

	if askSlide && c.OnFloor && !c.Sliding {
		c.Sliding = true
		c.SlideFrame = 0
		return
	}

	if c.Sliding {
		c.SlideFrame++
		if c.SlideFrame >= c.SlideDuration {
			c.Sliding = false
		}
	}

}

func (c *Character) SetImage() {
	var red uint8 = 0
	var green uint8 = 0
	var blue uint8 = 0

	if c.OnFloor {
		red = 150
	}

	if c.Jumping {
		green = 150
	}

	if c.Sliding {
		blue = 150
	}

	if c.Kicking || c.Sliding {
		c.Width = c.SlidingWidth
		c.Height = c.SlidingHeight
	} else {
		c.Width = c.StandingWidth
		c.Height = c.StandingHeight
	}

	c.Color = color.RGBA{R: red, G: green, B: blue, A: 150}
}

func (c *Character) HandleJump(askJump, askImproveJump bool) {

	defer func() { c.PosY += c.SpeedY }()

	if c.OnFloor {
		c.Jumping = false
		c.SpeedY = 0
	}

	if askJump && c.OnFloor && !c.Sliding {
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

func (c *Character) CheckFloorAndAdjust() {
	c.OnFloor = c.PosY >= float64(gScreenHeight)

	if c.OnFloor {
		c.PosY = float64(gScreenHeight)
	}
}
