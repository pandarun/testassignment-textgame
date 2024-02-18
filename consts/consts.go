package consts

type CommandType string
type ItemType string
type LocationName string

type Messages string

const (
	UnknownCommand  Messages = "неизвестная команда"
	ItemAddedFormat Messages = "предмет добавлен в инвентарь: %s"

	NoSuchItem                Messages = "нет такого"
	MoveCommandSuccessFormat  Messages = "ничего интересного. можно пройти - %s"
	RoomOnEnterFormat         Messages = "ты в своей комнате. можно пройти - %s"
	RoomEmptyLookAroundFormat Messages = "пустая комната. можно пройти - %s"
	MoveCommandFailFormat     Messages = "нет пути в %s"

	NothingInterestingFormat  Messages = "ничего интересного. можно пройти - %s"
	FirstStepLookAroundFormat Messages = "ты находишься в %s, надо собрать рюкзак и идти в универ. можно пройти - %s"

	StreetOnEnterFormat     Messages = "на улице весна. можно пройти - домой"
	KitchenLookAroundFormat Messages = "ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - %s"
	KitchenOnEnterFormat    Messages = "кухня, ничего интересного. можно пройти - %s"

	RoomLookAroundFormat Messages = "на столе: %s. можно пройти - %s"
)

const (
	LookAroundCommand CommandType = "осмотреться"
	MoveCommand       CommandType = "идти"
	TakeItemCommand   CommandType = "взять"
)

const (
	Backpack ItemType = "рюкзак"
	Keys     ItemType = "ключи"
	Notes    ItemType = "конспекты"
)

const (
	Kitchen  LocationName = "кухня"
	Corridor LocationName = "коридор"
	Room     LocationName = "комната"
	Street   LocationName = "улица"
)
