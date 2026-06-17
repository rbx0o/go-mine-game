package service

import "errors"

var EquipmentTypeNotFound = errors.New("Equipment type not found")
var NotEnoughCoal = errors.New("Not enough coal")
var EquipmentAlreadyBought = errors.New("The equipment has already been purchased")
