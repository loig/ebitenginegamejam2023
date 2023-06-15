package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ActionMoveRight int = iota
	ActionMoveLeft
)

func GetActionsFromKeyboard() (actions []int) {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		actions = append(actions, ActionMoveLeft)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		actions = append(actions, ActionMoveRight)
	}

	return
}
