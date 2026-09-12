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
func (g *GameService) BuyEquipment(equipment domain.EquipmentType) error {
	g.enterprise.Mtx.Lock()

	if equipment != domain.PickaxeType &&
		equipment != domain.VentilationType &&
		equipment != domain.TrolleysType {
		g.enterprise.Mtx.Unlock()
		return EquipmentTypeNotFound
	}
	if g.enterprise.AllEquipment[equipment].IsBought() {
		g.enterprise.Mtx.Unlock()
		return EquipmentAlreadyBought
	}
	if (g.enterprise.Balance - g.enterprise.AllEquipment[equipment].Cost()) < 0 {
		g.enterprise.Mtx.Unlock()
		return NotEnoughCoal
	}

	g.enterprise.AllEquipment[equipment].Buy()
	g.enterprise.Balance -= g.enterprise.AllEquipment[equipment].Cost()

	check := g.CheckAllEquipmentIsBought()

	g.enterprise.Mtx.Unlock()

	if check {
		g.mtx.Lock()
		g.gameResult.EndedAuto = true
		g.mtx.Unlock()

		g.StopGame()
	}

	return nil
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
func (g *GameService) GetEquipmentInfo() map[domain.EquipmentType]bool {
	g.enterprise.Mtx.RLock()
	defer g.enterprise.Mtx.RUnlock()

	result := make(map[domain.EquipmentType]bool, len(g.enterprise.AllEquipment))

	for key := range g.enterprise.AllEquipment {
		result[key] = g.enterprise.AllEquipment[key].IsBought()
	}

	return result
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
