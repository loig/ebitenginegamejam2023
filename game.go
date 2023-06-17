package main

type Game struct {
	Player     Character
	BigGhost   Character
	SmallGhost Character
	TallGhost  Character
}

func MakeGame() (g Game) {
	g.Player = InitPlayer()
	g.BigGhost = InitBigGhost()
	g.SmallGhost = InitSmallGhost()
	g.TallGhost = InitTallGhost()
	return
}
