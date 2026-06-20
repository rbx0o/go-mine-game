package domain

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
