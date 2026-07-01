package domain

import "github.com/google/uuid"

/*==================================================

Основная валюта всей игры - уголь. Используется и для
заработка, и для оплаты.
В будущем уголь может меняться/расширяться. Например
добавят типы угля.
Для этого создали свой именованный тип, чтобы в будущем
при надобности можно было бы переделать его в нужную
структуру.

Подобным образом создан тип данных для оборудования.
И реализованы константы (подобие enum).

==================================================*/

type Coal int

//==================================================

type EquipmentType string

const (
	PickaxeType     EquipmentType = "pickaxe"
	VentilationType EquipmentType = "ventilation"
	TrolleysType    EquipmentType = "trolleys"
)

//==================================================

type MinerType string

const (
	SmallMinerType  MinerType = "small_miner"
	NormalMinerType MinerType = "normal_miner"
	StrongMinerType MinerType = "strong_miner"
)

//==================================================

type ID uuid.UUID

/*
NewID() генерирует новый ID на основе UUID
*/
func NewID() (ID, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		var zeroID ID
		return zeroID, err
	}

	return ID(id), err
}

/*
ParseID() парсит ID из строки
*/
func ParseID(str string) (ID, error) {
	id, err := uuid.Parse(str)

	if err != nil {
		var zeroID ID
		return zeroID, err
	}

	return ID(id), err
}

/*
String() конвертирует ID в строку
*/
func (id ID) String() string {
	return uuid.UUID(id).String()
}

/*
IsZero() проверяет ID на нулевое значение
*/
func (id ID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}
