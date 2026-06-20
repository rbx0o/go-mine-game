package service

import "github.com/rbx0o/go-mine-game/domain"

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с оборудованием.

==================================================*/

func (g *GameService) BuyEquipment(equipment domain.EquipmentType) error {
	defer g.enterprise.Mtx.Unlock()
	g.enterprise.Mtx.Lock()

	if (g.enterprise.Balance - g.enterprise.AllEquipment[equipment].Cost()) < 0 {
		return NotEnoughCoal
	}
	if g.enterprise.AllEquipment[equipment].IsBought() {
		return EquipmentAlreadyBought
	}
	g.enterprise.AllEquipment[equipment].Buy()
	return nil
}

/*
GetEquipmentTypes возвращает информацию о всех типах доступного оборудования
*/
func (g *GameService) GetEquipmentTypesInfo() map[domain.EquipmentType]domain.EquipmentInfo {
	equipmentTypeMap := make(map[domain.EquipmentType]domain.EquipmentInfo, 3)

	equipmentTypeMap[domain.PickaxeType] = g.enterprise.AllEquipment[domain.PickaxeType].GetInfo()
	equipmentTypeMap[domain.VentilationType] = g.enterprise.AllEquipment[domain.VentilationType].GetInfo()
	equipmentTypeMap[domain.TrolleysType] = g.enterprise.AllEquipment[domain.TrolleysType].GetInfo()

	return equipmentTypeMap
}
