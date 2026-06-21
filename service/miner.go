package service

import (
	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с шахтёрами.

==================================================*/

func (g *GameService) GetMinerTypesInfo() map[domain.MinerType]domain.MinerConfig {
	return domain.GetMinerConfigs()
}
