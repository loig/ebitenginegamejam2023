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
	OffScreen      *ebiten.Image
}

func InitCharacter() (c Character) {
	c.PosY = 0
	c.OffScreen = ebiten.NewImage(gScreenWidth, gScreenHeight)
	return
}

func InitPlayer() (c Character) {
	c = InitCharacter()
	c.StandingWidth = gPlayerWidth
	c.StandingHeight = gPlayerHeight
	c.SlidingWidth = gPlayerSlidingWidth
	c.SlidingHeight = gPlayerSlidingHeight
	c.Width = c.StandingWidth
	c.Height = c.StandingHeight
	c.IncrSpeedY = gPlayerIncrSpeedY
	c.MaxSpeedY = gPlayerMaxSpeedY
	c.SlideDuration = gPlayerSlideDuration
	return
}

func InitBigGhost() (c Character) {
	c = InitCharacter()
	c.StandingWidth = gBigGhostWidth
	c.StandingHeight = gBigGhostHeight
	c.SlidingWidth = gBigGhostSlidingWidth
	c.SlidingHeight = gBigGhostSlidingHeight
	c.Width = c.StandingWidth
	c.Height = c.StandingHeight
	c.IncrSpeedY = gBigGhostIncrSpeedY
	c.MaxSpeedY = gBigGhostMaxSpeedY
	c.SlideDuration = gBigGhostSlideDuration
	return
}

func InitSmallGhost() (c Character) {
	c = InitCharacter()
	c.StandingWidth = gSmallGhostWidth
	c.StandingHeight = gSmallGhostHeight
	c.SlidingWidth = gSmallGhostSlidingWidth
	c.SlidingHeight = gSmallGhostSlidingHeight
	c.Width = c.StandingWidth
	c.Height = c.StandingHeight
	c.IncrSpeedY = gSmallGhostIncrSpeedY
	c.MaxSpeedY = gSmallGhostMaxSpeedY
	c.SlideDuration = gSmallGhostSlideDuration
	return
}

func InitTallGhost() (c Character) {
	c = InitCharacter()
	c.StandingWidth = gTallGhostWidth
	c.StandingHeight = gTallGhostHeight
	c.SlidingWidth = gTallGhostSlidingWidth
	c.SlidingHeight = gTallGhostSlidingHeight
	c.Width = c.StandingWidth
	c.Height = c.StandingHeight
	c.IncrSpeedY = gTallGhostIncrSpeedY
	c.MaxSpeedY = gTallGhostMaxSpeedY
	c.SlideDuration = gTallGhostSlideDuration
	return
}

func (c Character) Draw(screen *ebiten.Image, position int) {

	c.OffScreen.Clear()
	vector.DrawFilledRect(c.OffScreen, float32(gScreenWidth-c.Width)/2, float32(c.PosY-float64(c.Height)), float32(c.Width), float32(c.Height), c.Color, false)
	vector.DrawFilledRect(c.OffScreen, 0, 0, 20, 30, c.Color, false)

	options := ebiten.DrawImageOptions{}

	tx := 0.0
	ty := 0.0
	scalex := 1.0
	scaley := 1.0
	switch position {
	case DrawPositionBottom:
		tx = float64(gScreenWidth) / 4
		ty = float64(gScreenHeight) / 2
		scalex = 0.5
		scaley = -0.5
	case DrawPositionTop:
		tx = float64(gScreenWidth) / 4
		scalex = 0.5
		scaley = 0.5
	case DrawPositionTopLeft:
		scalex = 0.5
		scaley = 0.5
	case DrawPositionBottomLeft:
		ty = float64(gScreenHeight)
		scalex = 0.5
		scaley = -0.5
	case DrawPositionTopRight:
		tx = float64(gScreenWidth)
		scalex = -0.5
		scaley = 0.5
	case DrawPositionBottomRight:
		ty = float64(gScreenHeight)
		tx = float64(gScreenWidth)
		scalex = -0.5
		scaley = -0.5
	}
	options.GeoM.Scale(scalex, scaley)
	options.GeoM.Translate(tx, ty)

	screen.DrawImage(c.OffScreen, &options)
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
