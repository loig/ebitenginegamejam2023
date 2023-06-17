package main

func (g *Game) Update() error {

	actions := GetActionsFromKeyboard()
	g.Player.Update(actions)
	g.BigGhost.Update(actions)
	g.SmallGhost.Update(actions)
	g.TallGhost.Update(actions)

	return nil
}
