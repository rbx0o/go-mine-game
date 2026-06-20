package service

import (
	"time"

	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с шахтёрами.

==================================================*/

func (g *GameService) GetMinerTypesInfo() map[domain.MinerType]domain.MinerTypeInfo {
	minerTypeMap := make(map[domain.MinerType]domain.MinerTypeInfo, 3)

	minerTypeMap[domain.SmallMinerType] = domain.MinerTypeInfo{
		Salary:    5,
		Energy:    30,
		CoalCount: 1,
		Timeout:   3 * time.Second,
	}
	minerTypeMap[domain.NormalMinerType] = domain.MinerTypeInfo{
		Salary:    50,
		Energy:    45,
		CoalCount: 3,
		Timeout:   2 * time.Second,
	}
	minerTypeMap[domain.StrongMinerType] = domain.MinerTypeInfo{
		Salary:    450,
		Energy:    60,
		CoalCount: 10,
		Timeout:   1 * time.Second,
	}

	return minerTypeMap
}
