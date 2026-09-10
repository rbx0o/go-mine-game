package service

import "errors"

var EquipmentTypeNotFound = errors.New("Equipment type not found")
var NotEnoughCoal = errors.New("Not enough coal")
var EquipmentAlreadyBought = errors.New("The equipment has already been purchased")
var MinerTypeNotFound = errors.New("Miner type not found")
var GameServiceCtxAlreadyCanceled = errors.New("Game service context has already been canceled")
var MinersCtxAlreadyCanceled = errors.New("Miners context has already been canceled")
var GameNotEnd = errors.New("The game is not yet complete")
