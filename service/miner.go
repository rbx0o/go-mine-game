package service

import (
	"sync"

	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с шахтёрами.

==================================================*/

/*
GetMinerTypesInfo() возвращает map с информацией о всех типах шахтёрах, в виде ключ - значение:
MinerType - MinerConfig
*/
func (g *GameService) GetMinerTypesInfo() map[domain.MinerType]domain.MinerConfig {
	return domain.GetMinerConfigs()
}

/*
HireMiner() нанимает шахтёра на работу
*/
func (g *GameService) HireMiner(minerType domain.MinerType) error {
	var miner domain.Miner
	var chCoal <-chan domain.Coal
	var wg = &sync.WaitGroup{}
	var err error

	switch minerType {
	case domain.SmallMinerType:
		miner, err = domain.InitSmallMiner()
	case domain.NormalMinerType:
		miner, err = domain.InitNormalMiner()
	case domain.StrongMinerType:
		miner, err = domain.InitStrongMiner()
	default:
		err = MinerTypeNotFound
	}

	if err != nil {
		return err
	}

	wg.Add(1)
	chCoal = miner.Run(g.enterprise.Ctx, wg)
	g.enterprise.ActiveMiners[miner.GetInfo().ID] = miner

	go func() {
		g.enterprise.Balance += <-chCoal
		wg.Wait()
		g.enterprise.InactiveMiners[miner.GetInfo().ID] = miner
		delete(g.enterprise.ActiveMiners, miner.GetInfo().ID)
	}()

	return nil
}

/*
GetActiveMiners() возвращает копию map всех работающих в данный момент шахтёров
*/
func (g *GameService) GetActiveMiners() map[domain.ID]domain.Miner {
	result := make(map[domain.ID]domain.Miner, len(g.enterprise.ActiveMiners))

	for key, value := range g.enterprise.ActiveMiners {
		result[key] = value
	}

	return result
}

/*
GetInactiveMiners() возвращает копию map шахтёров закончивших работу
*/
func (g *GameService) GetInactiveMiners() map[domain.ID]domain.Miner {
	result := make(map[domain.ID]domain.Miner, len(g.enterprise.InactiveMiners))

	for key, value := range g.enterprise.InactiveMiners {
		result[key] = value
	}

	return result
}

/*
GetActiveMinersFilter() возвращает копию map всех работающих в данный момент шахтёров,
отфильтрованных по классу шахтёра
*/
func (g *GameService) GetActiveMinersFilter(minerType domain.MinerType) map[domain.ID]domain.Miner {
	result := make(map[domain.ID]domain.Miner, len(g.enterprise.ActiveMiners))

	for key, value := range g.enterprise.ActiveMiners {
		if value.GetInfo().MinerType == minerType {
			result[key] = value
		}
	}

	return result
}

/*
GetInactiveMinersFilter() возвращает копию map шахтёров закончивших работу,
отфильтрованных по классу шахтёра
*/
func (g *GameService) GetInactiveMinersFilter(minerType domain.MinerType) map[domain.ID]domain.Miner {
	result := make(map[domain.ID]domain.Miner, len(g.enterprise.InactiveMiners))

	for key, value := range g.enterprise.InactiveMiners {
		if value.GetInfo().MinerType == minerType {
			result[key] = value
		}
	}

	return result
}
