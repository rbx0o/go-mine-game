package domain

import (
	"context"
	"sync"
	"time"
)

// Файл в котором будут описаны шахтеры и базовая логика

/*==================================================

Есть 3 типа шахтёров. Для них лучше создать интерфейс.
Шахтёры должны уметь начинать и завершать работу, а также
возвращать информацию о себе.
Шахтёры - структуры, которые хранят в себе информацию о том,
как они должны работать.
Нужно также уметь возвращать некий snapshot, состояния шахтёра
в данный момент. Для этого можно создать структуры MinerInfo,
которая будет хранить информацию о шахтёре в данный момент и
возвращать её в методе GetInfo(). По сути во вне нужно возвращать
только сколько энергии осталось у конкретного шахтёра (в теории
можно ещё ID и класс шахтёра)

Также у каждого класса шахтёра будет поле Mutex. Почему так?
Mutex контролирует данные внутри одного шахтёра.
Один Mutex для всех шахтёров не нужен, потому что у каждого
разные данные.
Mutex нужен лишь в рамках одного конкретного шахтёра, потому что
гонка данных может возникнуть, если работа с данными конкретного
экземпляра шахтёра будет производиться в нескольких функциях
одновременно.

Run() - запускает горутину с шахтёром
GetInfo() - возвращает информацию об одном конкретном шахтёре
в данный момент

==================================================*/

type Miner interface {
	/*
		Run() запускает работу шахтёра.
		Запускается добыча угля в горутине по заданным параметрам (энергия, добыча и т.п.).
		В аргументах принимается контекст работы данной горутины (от предприятия).
		Возвращается канал, в которой передаётся добытый уголь.
	*/
	Run(context.Context, *sync.WaitGroup) <-chan Coal

	/*
		GetInfo() возвращает информацию о шахтёре.
		По сути просто возвращает копию поля info шахтёра
	*/
	GetInfo() MinerInfo
}

//==================================================

type MinerInfo struct {
	ID         ID  // ID шахтёра
	EnergyLeft int // Сколько энергии осталось
}

func InitMinerInfo(id ID, energy int) MinerInfo {
	return MinerInfo{
		ID:         id,
		EnergyLeft: energy,
	}
}

//==================================================

type SmallMiner struct {
	ID        ID            // ID шахтёра
	salary    Coal          // Оплата труда - 5
	energy    int           // Энергия - 30
	coalCount Coal          // За одну добычу - 1
	timeout   time.Duration // Задержка между добычами - 3
	info      MinerInfo     // Информация о шахтёре в данный момент
	mtx       sync.Mutex    // Mutex для контроля гонки данных
}

func InitSmallMiner() (*SmallMiner, error) {
	id, err := NewID()

	if err != nil {
		return &SmallMiner{}, err
	}

	return &SmallMiner{
		ID:        id,
		salary:    minerConfigs[SmallMinerType].Salary,
		energy:    minerConfigs[SmallMinerType].Energy,
		coalCount: minerConfigs[SmallMinerType].CoalCount,
		timeout:   minerConfigs[SmallMinerType].Timeout,
		info:      InitMinerInfo(id, minerConfigs[SmallMinerType].Energy),
	}, nil
}

func (m *SmallMiner) GetInfo() MinerInfo {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.info
}

func (m *SmallMiner) Run(ctx context.Context, wg *sync.WaitGroup) <-chan Coal {
	ch := make(chan Coal)

	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()

		for i := 0; i < m.energy; i++ {
			select {
			case <-ctx.Done():
				return
			case <-time.After(m.timeout):
				ch <- m.coalCount

				m.mtx.Lock()
				m.info.EnergyLeft--
				m.mtx.Unlock()
			}
		}
	}()

	return ch
}

//==================================================

type NormalMiner struct {
	ID        ID            // ID шахтёра
	salary    Coal          // Оплата труда - 50
	energy    int           // Энергия - 45
	coalCount Coal          // За одну добычу - 3
	timeout   time.Duration // Задержка между добычами - 2
	info      MinerInfo     // Информация о шахтёре в данный момент
	mtx       sync.Mutex    // Mutex для контроля гонки данных
}

func InitNormalMiner() (*NormalMiner, error) {
	id, err := NewID()

	if err != nil {
		return &NormalMiner{}, err
	}

	return &NormalMiner{
		ID:        id,
		salary:    minerConfigs[NormalMinerType].Salary,
		energy:    minerConfigs[NormalMinerType].Energy,
		coalCount: minerConfigs[NormalMinerType].CoalCount,
		timeout:   minerConfigs[NormalMinerType].Timeout,
		info:      InitMinerInfo(id, minerConfigs[NormalMinerType].Energy),
	}, nil
}

func (m *NormalMiner) GetInfo() MinerInfo {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.info
}

func (m *NormalMiner) Run(ctx context.Context, wg *sync.WaitGroup) <-chan Coal {
	ch := make(chan Coal)

	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()

		for i := 0; i < m.energy; i++ {
			select {
			case <-ctx.Done():
				return
			case <-time.After(m.timeout):
				ch <- m.coalCount

				m.mtx.Lock()
				m.info.EnergyLeft--
				m.mtx.Unlock()
			}
		}
	}()

	return ch
}

//==================================================

type StrongMiner struct {
	ID        ID            // ID шахтёра
	salary    Coal          // Оплата труда - 450
	energy    int           // Энергия - 60
	coalCount Coal          // За одну добычу - 10
	timeout   time.Duration // Задержка между добычами - 1
	info      MinerInfo     // Информация о шахтёре в данный момент
	progress  Coal          // coalCount увеличивается на progress - 3
	mtx       sync.Mutex    // Mutex для контроля гонки данных
}

func InitStrongMiner() (*StrongMiner, error) {
	id, err := NewID()

	if err != nil {
		return &StrongMiner{}, err
	}

	return &StrongMiner{
		ID:        id,
		salary:    minerConfigs[StrongMinerType].Salary,
		energy:    minerConfigs[StrongMinerType].Energy,
		coalCount: minerConfigs[StrongMinerType].CoalCount,
		timeout:   minerConfigs[StrongMinerType].Timeout,
		info:      InitMinerInfo(id, minerConfigs[StrongMinerType].Energy),
		progress:  minerConfigs[StrongMinerType].Progress,
	}, nil
}

func (m *StrongMiner) GetInfo() MinerInfo {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.info
}

func (m *StrongMiner) Run(ctx context.Context, wg *sync.WaitGroup) <-chan Coal {
	ch := make(chan Coal)

	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()

		for i := 0; i < m.energy; i++ {
			select {
			case <-ctx.Done():
				return

			case <-time.After(m.timeout):
				ch <- m.coalCount

				m.mtx.Lock()
				m.info.EnergyLeft--
				m.coalCount += m.progress
				m.mtx.Unlock()
			}
		}
	}()

	return ch
}
