package commands

import "text_game/game"

func (c TakeItemCommand) Execute(g *game.State) *game.State {

	item, message, err := g.CurrentLocation.TakeItem(c.NewItem)

	if err != nil {
		g.Message = err.Error()
		return g
	}

	g.Player.AddItem(item)
	g.Message = message

	return g
}
