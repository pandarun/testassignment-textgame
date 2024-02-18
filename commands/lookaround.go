package commands

import "text_game/game"

func (c LookAroundCommand) Execute(g *game.State) *game.State {

	g.Message = g.CurrentLocation.LookAround(g.Step)

	return g
}
