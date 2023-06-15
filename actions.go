package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ActionMoveRight int = iota
	ActionMoveLeft
	ActionJump
)

func GetActionsFromKeyboard() (actions []int) {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		actions = append(actions, ActionMoveLeft)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		actions = append(actions, ActionMoveRight)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		actions = append(actions, ActionJump)
	}

	return
}
