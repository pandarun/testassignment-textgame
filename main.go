package main

import (
	"bufio"
	"fmt"
	"os"
	"text_game/consts"
	"text_game/game"
	"text_game/locations"
	"text_game/parsers"
)

var gameState game.State

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/

	initGame()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Print 'q:' to exit")

	for scanner.Scan() && scanner.Text() != "q:" {
		fmt.Println(handleCommand(scanner.Text()))
	}

}

type GameBuilder struct {
	Rooms           []game.Location
	Command         []interface{}
	InitialLocation game.Location
}

func (g *GameBuilder) Build() game.State {
	return game.State{Player: game.Player{Inventory: game.Inventory{}}, CurrentLocation: g.InitialLocation}
}

func (g *GameBuilder) SetInitialLocation(location game.Location) *GameBuilder {
	g.InitialLocation = location
	return g
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все локации
		если что-то было - оно корректно перезатирается
	*/
	gameBuilder := new(GameBuilder)

	kitchen := locations.NewLocation(consts.Kitchen)
	corridor := locations.NewLocation(consts.Corridor)

	room := locations.NewLocation(consts.Room)
	room.AddItem(game.InventoryItem{Name: consts.Backpack})
	room.AddItem(game.InventoryItem{Name: consts.Keys})
	room.AddItem(game.InventoryItem{Name: consts.Notes})

	street := locations.NewLocation(consts.Street)

	kitchen.AddPath(corridor)

	corridor.AddPath(kitchen)
	corridor.AddPath(room)
	corridor.AddPath(street)

	room.AddPath(corridor)
	street.AddPath(corridor)

	gameBuilder.SetInitialLocation(kitchen)

	gameState = gameBuilder.Build()

}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	parser := parsers.NewCommandParser()

	gameCommand, err := parser.ParseCommand(command)
	if err != nil {
		return err.Error()
	}

	newState := gameState.Accept(gameCommand)

	return newState.String()
}
