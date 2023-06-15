package main

func (g *Game) Update() error {

	g.Player.Update(GetActionsFromKeyboard())

	return nil
}
