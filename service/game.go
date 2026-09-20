package service

// Здесь будет описан весь сервис и процессы того, что пользователь может делать с игрой

import (
	"context"
	"sync"
	"time"

	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Есть структура, которая представляет из себя саму игру.
Данную игру можно инициализировать (при запуске игры),
запускать и завершать.

Сама игра содержит в себе предприятие.

Также есть структура, которая представляет из себя результат игры:
время, баланс, оборудования, нанятые шахтёры и завершилась ли игра досрочно

==================================================*/

type GameState string

const (
	Created  GameState = "game_created"
	Running  GameState = "game_running"
	Finished GameState = "game_finished"
)

//==================================================

type GameService struct {
	enterprise *domain.Enterprise

	ctx       context.Context
	ctxCancel context.CancelFunc

	minerCtx       context.Context
	minerCtxCancel context.CancelFunc

	gameResult *GameResult

	wg    *sync.WaitGroup
	mtx   sync.Mutex
	state GameState
}

/*
InitGameService
создаёт и инициализирует объект игры
*/
func InitGameService() *GameService {
	tempCtx, tempCtxCancel := context.WithCancel(context.Background())
	tempMinerCtx, tempMinerCtxCancel := context.WithCancel(tempCtx)

	return &GameService{
		enterprise: domain.InitEnterprise(),

		ctx:       tempCtx,
		ctxCancel: tempCtxCancel,

		minerCtx:       tempMinerCtx,
		minerCtxCancel: tempMinerCtxCancel,

		gameResult: InitGameResult(),

		wg:    &sync.WaitGroup{},
		mtx:   sync.Mutex{},
		state: Created,
	}
}

/*
Start
запускает игру
*/
func (g *GameService) Start() error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	switch g.state {
	case Running:
		return GameAlreadyRunning
	case Finished:
		return GameAlreadyFinished
	case Created:
		g.state = Running
		g.gameResult.StartTime = time.Now()
	}

	g.StartPassiveIncome(g.ctx)
	return nil
}

/*
StopGame
ручное отправление запроса о завершении игры
*/
func (g *GameService) StopGame() (error, *GameResult) {
	return g.finishGame(false)
}

/*
finishGame
завершает игру
*/
func (g *GameService) finishGame(endedAuto bool) (error, *GameResult) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	switch g.state {
	case Running:
		g.state = Finished
	case Finished:
		return GameAlreadyFinished, nil
	case Created:
		return GameNotRunningYet, nil
	}

	if err := g.ctx.Err(); err != nil {
		return GameServiceCtxAlreadyCanceled, nil
	} else {
		g.ctxCancel()

		g.wg.Wait()
		g.enterprise.Mtx.Lock()
		defer g.enterprise.Mtx.Unlock()

		g.gameResult.Balance = g.enterprise.Balance
		g.gameResult.EndTime = time.Now()
		g.gameResult.DurationTime = time.Since(g.gameResult.StartTime)

		resultEquipment := make(map[domain.EquipmentType]bool, len(g.enterprise.AllEquipment))
		for key := range g.enterprise.AllEquipment {
			resultEquipment[key] = g.enterprise.AllEquipment[key].IsBought()
		}
		g.gameResult.ResultEquipment = resultEquipment

		resultMiners := make(map[domain.ID]domain.Miner, len(g.enterprise.InactiveMiners))
		for key, value := range g.enterprise.InactiveMiners {
			resultMiners[key] = value
		}
		g.gameResult.ResultMiners = resultMiners

		g.gameResult.EndedAuto = endedAuto

		return nil, g.gameResult
	}
}

/*
StopMiners
останавливает только работу шахтёров
*/
func (g *GameService) StopMiners() error {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	if err := g.minerCtx.Err(); err != nil {
		return MinersCtxAlreadyCanceled
	} else {
		g.minerCtxCancel()
		return nil
	}
}

/*
GetGameResult
возвращает результат игры
*/
func (g *GameService) GetGameResult() (error, *GameResult) {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	if err := g.ctx.Err(); err == nil {
		return GameNotEnd, nil
	} else {
		return nil, g.gameResult
	}
}

/*
GetState возвращает состояние игры
*/
func (g *GameService) GetState() GameState {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.state
}

//==================================================

type GameResult struct {
	Balance         domain.Coal
	StartTime       time.Time
	EndTime         time.Time
	DurationTime    time.Duration
	EndedAuto       bool
	ResultEquipment map[domain.EquipmentType]bool
	ResultMiners    map[domain.ID]domain.Miner
}

func InitGameResult() *GameResult {
	return &GameResult{}
}
