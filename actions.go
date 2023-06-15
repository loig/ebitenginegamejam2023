package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ActionMoveRight int = iota
	ActionMoveLeft
	ActionJump
	ActionImproveJump
	ActionNumber
)

func GetActionsFromKeyboard() (actions [ActionNumber]bool) {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		actions[ActionMoveLeft] = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		actions[ActionMoveRight] = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		actions[ActionImproveJump] = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		actions[ActionJump] = true
	}

	return
}
