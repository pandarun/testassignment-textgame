package locations

import (
	"testing"
	"text_game/consts"
	"text_game/game"
)

func TestCorridor_LookAround(t *testing.T) {
	corridor := corridor{
		locationState: locationState{
			name:      consts.Corridor,
			items:     make(map[consts.ItemType]game.InventoryItem),
			nextPaths: []game.Location{NewLocation(consts.Room), NewLocation(consts.Kitchen), NewLocation(consts.Street)},
		},
	}

	expected := "ничего интересного. можно пройти - комната, кухня, улица"
	actual := corridor.LookAround(2)

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestCorridor_MoveTo(t *testing.T) {
	room := &room{locationState: locationState{name: consts.Room, items: make(map[consts.ItemType]game.InventoryItem)}}
	corridor := corridor{
		locationState: locationState{
			name:      consts.Corridor,
			items:     make(map[consts.ItemType]game.InventoryItem),
			nextPaths: []game.Location{room},
		},
	}
	room.nextPaths = []game.Location{&corridor}

	expected := "ты в своей комнате. можно пройти - коридор"
	_, actual, _ := corridor.MoveTo(consts.Room)

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}

}

func TestRoom_LookAround(t *testing.T) {
	room := room{
		locationState: locationState{
			name:      consts.Room,
			items:     make(map[consts.ItemType]game.InventoryItem),
			nextPaths: []game.Location{NewLocation(consts.Corridor)},
		},
	}

	expected := "пустая комната. можно пройти - коридор"
	actual := room.LookAround(1)

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestRoom_MoveTo(t *testing.T) {
	corridor := &corridor{locationState: locationState{name: consts.Corridor, items: make(map[consts.ItemType]game.InventoryItem)}}
	room := room{
		locationState: locationState{
			name:      consts.Room,
			items:     make(map[consts.ItemType]game.InventoryItem),
			nextPaths: []game.Location{corridor},
		},
	}
	corridor.nextPaths = []game.Location{&room}

	expected := "ничего интересного. можно пройти - комната"
	_, actual, _ := room.MoveTo(consts.Corridor)

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestRoom_TakeItem(t *testing.T) {
	items := make(map[consts.ItemType]game.InventoryItem)
	items[consts.Backpack] = game.InventoryItem{Name: consts.Backpack}
	room := room{
		locationState: locationState{
			name:      consts.Room,
			items:     items,
			nextPaths: []game.Location{NewLocation(consts.Corridor)},
		},
	}

	expected := "предмет добавлен в инвентарь: рюкзак"
	_, actual, _ := room.TakeItem(game.InventoryItem{Name: consts.Backpack})

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestKitchen_LookAround(t *testing.T) {
	kitchen := kitchen{
		locationState: locationState{
			name:      consts.Kitchen,
			items:     make(map[consts.ItemType]game.InventoryItem),
			nextPaths: []game.Location{NewLocation(consts.Corridor)},
		},
	}

	expected := "ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	actual := kitchen.LookAround(1)

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
