package commands

import (
	"text_game/consts"
	"text_game/game"
)

type LookAroundCommand struct {
}

type MoveCommand struct {
	NewLocation consts.LocationName
}

type TakeItemCommand struct {
	NewItem game.InventoryItem
}
