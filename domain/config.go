package domain

import "time"

/*==================================================

Информация о шахтёрах используется в нескольких местах.
Эта информация может изменяться и должен быть для
нескольких мест один источник правды.
Поэтому создаётся данный конфигурационный файл, откуда
будут браться все данные.

==================================================*/

type MinerConfig struct {
	Salary    Coal
	Energy    int
	CoalCount Coal
	Timeout   time.Duration
	Progress  Coal
}

var minerConfigs = map[MinerType]MinerConfig{
	SmallMinerType: {
		Salary:    5,
		Energy:    30,
		CoalCount: 1,
		Timeout:   3 * time.Second,
	},
	NormalMinerType: {
		Salary:    50,
		Energy:    45,
		CoalCount: 3,
		Timeout:   2 * time.Second,
	},
	StrongMinerType: {
		Salary:    450,
		Energy:    60,
		CoalCount: 10,
		Timeout:   1 * time.Second,
		Progress:  3,
	},
}

func GetMinerConfigs() map[MinerType]MinerConfig {
	tempMinerConfigs := make(map[MinerType]MinerConfig, len(minerConfigs))

	for key, value := range minerConfigs {
		tempMinerConfigs[key] = value
	}

	return tempMinerConfigs
}
