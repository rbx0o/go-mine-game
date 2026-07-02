package domain

import (
	"testing"
	"time"
)

/*
Test_GetMinerConfigsReturnAllConfigs
проверяет количество вернувшихся конфигов
*/
func Test_GetMinerConfigsReturnAllConfigs(t *testing.T) {
	got := len(GetMinerConfigs())
	want := len(minerConfigs)

	if got != want {
		t.Fatalf("got len = %d, want = %d", got, want)
	}
}

/*
Test_GetMinerConfigsReturnCorrectValues
проверяет что функция вернула корректные значения
*/
func Test_GetMinerConfigsReturnCorrectValues(t *testing.T) {
	gotConfigs := GetMinerConfigs()

	tests := []struct {
		name      string
		minerType MinerType
		want      MinerConfig
	}{
		{
			name:      "small miner config",
			minerType: SmallMinerType,
			want: MinerConfig{
				Salary:    5,
				Energy:    30,
				CoalCount: 1,
				Timeout:   3 * time.Second,
			},
		},
		{
			name:      "normal miner config",
			minerType: NormalMinerType,
			want: MinerConfig{
				Salary:    50,
				Energy:    45,
				CoalCount: 3,
				Timeout:   2 * time.Second,
			},
		},
		{
			name:      "strong miner config",
			minerType: StrongMinerType,
			want: MinerConfig{
				Salary:    450,
				Energy:    60,
				CoalCount: 10,
				Timeout:   1 * time.Second,
				Progress:  3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := gotConfigs[tt.minerType]
			if !ok {
				t.Fatalf("config for miner type %v not found", tt.minerType)
			}
			if got != tt.want {
				t.Errorf("config = %+v; want %+v", got, tt.want)
			}
		})
	}
}

/*
Test_GetMinerConfigsReturnCopy
проверяет что функция вернула копию
*/
func Test_GetMinerConfigsReturnCopy(t *testing.T) {
	configs := GetMinerConfigs()

	delete(configs, SmallMinerType)

	configs[NormalMinerType] = MinerConfig{
		Salary:    999,
		Energy:    999,
		CoalCount: 999,
		Timeout:   999 * time.Second,
		Progress:  999,
	}

	freshConfigs := GetMinerConfigs()

	smallConfig, ok := freshConfigs[SmallMinerType]
	if !ok {
		t.Fatal("SmallMinerType was deleted from original configs")
	}

	wantSmallConfig := MinerConfig{
		Salary:    5,
		Energy:    30,
		CoalCount: 1,
		Timeout:   3 * time.Second,
	}

	if smallConfig != wantSmallConfig {
		t.Errorf("small config = %+v; want %+v", smallConfig, wantSmallConfig)
	}

	wantNormalConfig := MinerConfig{
		Salary:    50,
		Energy:    45,
		CoalCount: 3,
		Timeout:   2 * time.Second,
	}

	if freshConfigs[NormalMinerType] != wantNormalConfig {
		t.Errorf("normal config = %+v; want %+v", freshConfigs[NormalMinerType], wantNormalConfig)
	}
}
