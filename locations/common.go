package locations

import (
	"fmt"
	"sort"
	"strings"
	"text_game/consts"
	"text_game/game"
)

type locationState struct {
	name      consts.LocationName
	items     map[consts.ItemType]game.InventoryItem
	nextPaths []game.Location
}

func (s *locationState) GetName() consts.LocationName {
	return s.name
}

func (s *locationState) GetPaths() []game.Location {
	return s.nextPaths
}

func (s *locationState) GetItems() []game.InventoryItem {
	var items []game.InventoryItem
	for _, item := range s.items {
		items = append(items, item)
	}

	return items
}

func (s *locationState) TakeItem(item game.InventoryItem) (game.InventoryItem, string, error) {

	if _, ok := s.items[item.Name]; !ok {
		return game.InventoryItem{}, "", fmt.Errorf(string(consts.NoSuchItem))
	}

	s.RemoveItem(item)
	return item, fmt.Sprintf(string(consts.ItemAddedFormat), string(item.Name)), nil
}

func (s *locationState) ItemsString() string {
	var items []string
	for _, item := range s.GetItems() {
		items = append(items, string(item.Name))
	}

	sort.Slice(items, func(i, j int) bool { return items[i] < items[j] })

	nextPathsString := strings.Join(items, ", ")
	return nextPathsString
}

func (s *locationState) NextPathString() string {
	var nextPaths []string
	for _, item := range s.GetPaths() {
		nextPaths = append(nextPaths, string(item.GetName()))
	}

	nextPathsString := strings.Join(nextPaths, ", ")
	return nextPathsString
}

func (s *locationState) GetOnEnterMessage() string {
	return fmt.Sprintf(string(consts.MoveCommandSuccessFormat), s.NextPathString())
}

func (s *locationState) LookAround(step int) string {

	if step == 1 {
		return fmt.Sprintf(string(consts.FirstStepLookAroundFormat), s.name, s.NextPathString())
	}

	return fmt.Sprintf(string(consts.NothingInterestingFormat), s.NextPathString())
}

func (s *locationState) MoveTo(location consts.LocationName) (game.Location, string, error) {
	for _, path := range s.nextPaths {
		if path.GetName() == location {

			return path, path.GetOnEnterMessage(), nil
		}
	}

	return nil, "", fmt.Errorf(string(consts.MoveCommandFailFormat), location)
}

func (s *locationState) AddItem(item game.InventoryItem) {
	s.items[item.Name] = item
}

func (s *locationState) RemoveItem(item game.InventoryItem) {
	delete(s.items, item.Name)
}

func (s *locationState) AddPath(room game.Location) {
	s.nextPaths = append(s.nextPaths, room)
}
