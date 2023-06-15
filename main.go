package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowTitle("Petits fantômes")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
