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
	return EnterpriseInfo{
		Balance: g.enterprise.Balance,

		ActiveMiners:   g.GetActiveMiners(),
		InactiveMiners: g.GetInactiveMiners(),

		AllEquipment: g.GetEquipmentInfo(),
	}
}
