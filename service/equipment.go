package service

import (
	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с оборудованием.

==================================================*/

/*
BuyEquipment
совершает покупку оборудования
*/
func (g *GameService) BuyEquipment(equipment domain.EquipmentType) (error, *domain.EquipmentInfo) {
	g.mtx.Lock()
	g.enterprise.Mtx.Lock()

	switch g.state {
	case Finished:
		g.mtx.Unlock()
		g.enterprise.Mtx.Unlock()
		return GameAlreadyFinished, nil
	case Created:
		g.mtx.Unlock()
		g.enterprise.Mtx.Unlock()
		return GameNotRunningYet, nil
	}

	if equipment != domain.PickaxeType &&
		equipment != domain.VentilationType &&
		equipment != domain.TrolleysType {
		g.mtx.Unlock()
		g.enterprise.Mtx.Unlock()
		return EquipmentTypeNotFound, nil
	}
	if g.enterprise.AllEquipment[equipment].IsBought() {
		g.mtx.Unlock()
		g.enterprise.Mtx.Unlock()
		return EquipmentAlreadyBought, nil
	}
	if (g.enterprise.Balance - g.enterprise.AllEquipment[equipment].Cost()) < 0 {
		g.mtx.Unlock()
		g.enterprise.Mtx.Unlock()
		return NotEnoughCoal, nil
	}

	g.enterprise.AllEquipment[equipment].Buy()
	g.enterprise.Balance -= g.enterprise.AllEquipment[equipment].Cost()

	check := g.CheckAllEquipmentIsBought()

	g.mtx.Unlock()
	g.enterprise.Mtx.Unlock()

	if check {
		g.finishGame(true)
	}

	eqInfo := g.enterprise.AllEquipment[equipment].GetInfo()
	return nil, &eqInfo
}

/*
GetEquipmentTypesInfo
возвращает информацию о всех типах доступного оборудования
*/
func (g *GameService) GetEquipmentTypesInfo() map[domain.EquipmentType]domain.EquipmentInfo {
	equipmentTypeMap := make(map[domain.EquipmentType]domain.EquipmentInfo, 3)

	equipmentTypeMap[domain.PickaxeType] = g.enterprise.AllEquipment[domain.PickaxeType].GetInfo()
	equipmentTypeMap[domain.VentilationType] = g.enterprise.AllEquipment[domain.VentilationType].GetInfo()
	equipmentTypeMap[domain.TrolleysType] = g.enterprise.AllEquipment[domain.TrolleysType].GetInfo()

	return equipmentTypeMap
}

/*
GetEquipmentInfo
возвращает информацию о том какое оборудование куплено/не куплено
*/
func (g *GameService) GetEquipmentInfo() (error, map[domain.EquipmentType]bool) {
	g.enterprise.Mtx.RLock()
	defer g.enterprise.Mtx.RUnlock()

	switch g.state {
	case Created:
		return GameNotRunningYet, nil
	case Finished:
		return GameAlreadyFinished, nil
	}

	result := make(map[domain.EquipmentType]bool, len(g.enterprise.AllEquipment))

	for key := range g.enterprise.AllEquipment {
		result[key] = g.enterprise.AllEquipment[key].IsBought()
	}

	return nil, result
}

/*
CheckAllEquipmentIsBought
проверяет куплено ли всё оборудование
*/
func (g *GameService) CheckAllEquipmentIsBought() bool {
	if g.enterprise.AllEquipment[domain.PickaxeType].IsBought() &&
		g.enterprise.AllEquipment[domain.VentilationType].IsBought() &&
		g.enterprise.AllEquipment[domain.TrolleysType].IsBought() {
		return true
	} else {
		return false
	}
}
