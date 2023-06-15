package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowTitle("Petits fantômes")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := MakeGame()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}

}
