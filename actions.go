package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func GetActionsFromKeyboard() (actions [ActionNumber]bool) {

	actions[ActionImproveJump] = ebiten.IsKeyPressed(ebiten.KeyUp)
	actions[ActionJump] = inpututil.IsKeyJustPressed(ebiten.KeyUp)
	actions[ActionSlide] = inpututil.IsKeyJustPressed(ebiten.KeyDown)
	actions[ActionKick] = ebiten.IsKeyPressed(ebiten.KeyLeft)

	return
}
