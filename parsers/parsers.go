package parsers

import (
	"fmt"
	"strings"
	"text_game/commands"
	"text_game/consts"
	"text_game/game"
)

type CommandParser interface {
	ParseCommand(command string) (game.Commander, error)
}

type ItemParser interface {
	parseItem(item string) string
}

type LocationParser interface {
	parseLocation(location string) string
}
type DefaultLocationParserImplementation struct {
}

func (p *DefaultLocationParserImplementation) parseLocation(location string) (consts.LocationName, error) {

	switch consts.LocationName(location) {
	case consts.Kitchen:
		fallthrough
	case consts.Corridor:
		fallthrough
	case consts.Room:
		fallthrough
	case consts.Street:
		return consts.LocationName(location), nil
	default:
		return "", fmt.Errorf("Unknown location")
	}
}

func NewCommandParser() CommandParser {
	return &DefaultCommandParserImplementation{}
}

type DefaultCommandParserImplementation struct {
}

type DefaultItemParserImplementation struct {
}

func (p *DefaultItemParserImplementation) parseItem(item string) (game.InventoryItem, error) {

	switch consts.ItemType(item) {
	case consts.Backpack:
		fallthrough
	case consts.Keys:
		fallthrough
	case consts.Notes:
		return game.InventoryItem{Name: consts.ItemType(item)}, nil
	default:
		return game.InventoryItem{}, fmt.Errorf(string(consts.NoSuchItem))
	}
}

func (p *DefaultCommandParserImplementation) ParseCommand(command string) (game.Commander, error) {

	commandParts := strings.Split(command, " ")

	if len(commandParts) < 1 {
		return nil, fmt.Errorf(string(consts.UnknownCommand))
	}

	switch consts.CommandType(commandParts[0]) {
	case consts.LookAroundCommand:
		return commands.LookAroundCommand{}, nil
	case consts.MoveCommand:

		locationParser := new(DefaultLocationParserImplementation)
		location, err := locationParser.parseLocation(commandParts[1])
		if err != nil {

			return nil, err
		}

		return commands.MoveCommand{NewLocation: location}, nil
	case consts.TakeItemCommand:
		itemParser := new(DefaultItemParserImplementation)
		item, err := itemParser.parseItem(commandParts[1])
		if err != nil {
			return nil, err
		}

		return commands.TakeItemCommand{NewItem: item}, nil
	default:
		return nil, fmt.Errorf(string(consts.UnknownCommand))
	}

}
