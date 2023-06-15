package main

type Game struct {
	Player Character
}

func MakeGame() (g Game) {
	g.Player = InitPlayer()
	return
}
