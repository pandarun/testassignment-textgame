package commands

import "text_game/game"

func (c MoveCommand) Execute(g *game.State) *game.State {

	newLocation, message, err := g.CurrentLocation.MoveTo(c.NewLocation)
	if err != nil {
		g.Message = err.Error()
		return g
	}

	g.CurrentLocation = newLocation
	g.Message = message

	return g
}
