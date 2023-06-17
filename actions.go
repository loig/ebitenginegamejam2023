package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ActionJump int = iota
	ActionImproveJump
	ActionSlide
	ActionKick
	ActionNumber
)

func GetActionsFromKeyboard() (actions [ActionNumber]bool) {

	actions[ActionImproveJump] = ebiten.IsKeyPressed(ebiten.KeyUp)
	actions[ActionJump] = inpututil.IsKeyJustPressed(ebiten.KeyUp)
	actions[ActionSlide] = inpututil.IsKeyJustPressed(ebiten.KeyDown)
	actions[ActionKick] = ebiten.IsKeyPressed(ebiten.KeyLeft)

	return
}
