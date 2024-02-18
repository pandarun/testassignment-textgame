package locations

import (
	"text_game/consts"
	"text_game/game"
)

func NewLocation(name consts.LocationName) game.Location {

	switch name {
	case consts.Room:
		return &room{locationState: locationState{name: name, items: make(map[consts.ItemType]game.InventoryItem)}}
	case consts.Kitchen:
		return &kitchen{locationState: locationState{name: name, items: make(map[consts.ItemType]game.InventoryItem)}}
	case consts.Corridor:
		return &corridor{locationState: locationState{name: name, items: make(map[consts.ItemType]game.InventoryItem)}}
	case consts.Street:
		return &street{locationState: locationState{name: name, items: make(map[consts.ItemType]game.InventoryItem)}}
	default:
		panic("Unknown location")
	}
}
