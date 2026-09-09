package service

import "github.com/rbx0o/go-mine-game/domain"

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с работой предприятия.

==================================================*/

type EnterpriseInfo struct {
	Balance domain.Coal

	ActiveMiners   map[domain.ID]domain.Miner
	InactiveMiners map[domain.ID]domain.Miner

	AllEquipment map[domain.EquipmentType]bool
}

/*
GetEnterpriseInfo
возвращает структуру данных с информацией о предприятии в данный момент
*/
func (g *GameService) GetEnterpriseInfo() EnterpriseInfo {
	defer g.enterprise.Mtx.RLock()
	g.enterprise.Mtx.RUnlock()

	// ActiveMiners
	activeMiners := make(map[domain.ID]domain.Miner, len(g.enterprise.ActiveMiners))
	for key, value := range g.enterprise.ActiveMiners {
		activeMiners[key] = value
	}

	// InactiveMiners
	inactiveMiners := make(map[domain.ID]domain.Miner, len(g.enterprise.InactiveMiners))
	for key, value := range g.enterprise.InactiveMiners {
		inactiveMiners[key] = value
	}

	// EquipmentInfo
	equipmentInfo := make(map[domain.EquipmentType]bool, len(g.enterprise.AllEquipment))
	for key := range g.enterprise.AllEquipment {
		equipmentInfo[key] = g.enterprise.AllEquipment[key].IsBought()
	}

	return EnterpriseInfo{
		Balance: g.enterprise.Balance,

		ActiveMiners:   activeMiners,
		InactiveMiners: inactiveMiners,

		AllEquipment: equipmentInfo,
	}
}
