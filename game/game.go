package game

import (
	"text_game/consts"
)

type Commander interface {
	Execute(*State) *State
}

type InventoryItemCollection interface {
	RemoveItem(item InventoryItem)
	AddItem(item InventoryItem)
	GetItems() []InventoryItem
	ItemsString() string
}

type PathCollection interface {
	AddPath(location Location)
	GetPaths() []Location
	NextPathString() string
}

type Location interface {
	TakeItem(item InventoryItem) (InventoryItem, string, error)
	LookAround(step int) string
	MoveTo(locationName consts.LocationName) (Location, string, error)
	GetOnEnterMessage() string

	GetName() consts.LocationName

	PathCollection
	InventoryItemCollection
}

type State struct {
	Player          Player
	CurrentLocation Location
	Message         string
	Step            int
}

func (g *State) Accept(command Commander) *State {

	g.Step++

	return command.Execute(g)
}

func (g *State) String() string {
	// здесь будет происходить отображение состояния игры
	return g.Message
}

type Inventory struct {
	Items []InventoryItem
}

type InventoryItem struct {
	Name consts.ItemType
}

type Player struct {
	Inventory Inventory
}

func (p *Player) AddItem(item InventoryItem) {
	p.Inventory.Items = append(p.Inventory.Items, item)
}
