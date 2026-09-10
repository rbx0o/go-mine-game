package service

// Здесь будет описан весь сервис и процессы того, что пользователь может делать с игрой

import (
	"context"
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

type GameService struct {
	enterprise *domain.Enterprise

	ctx       context.Context
	ctxCancel context.CancelFunc

	minerCtx       context.Context
	minerCtxCancel context.CancelFunc

	gameResult *GameResult
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
	}
}

/*
Start
запускает игру
*/
func (g *GameService) Start() error {
	g.gameResult.startTime = time.Now()
	g.StartPassiveIncome(g.ctx)
	return nil
}

/*
StopGame
останавливает игру целиком
*/
func (g *GameService) StopGame() (error, *GameResult) {
	if err := g.ctx.Err(); err != nil {
		return GameServiceCtxAlreadyCanceled, nil
	} else {
		g.ctxCancel()

		g.gameResult.balance = g.enterprise.Balance
		g.gameResult.endTime = time.Now()
		g.gameResult.durationTime = time.Since(g.gameResult.startTime)
		g.gameResult.resultEquipment = g.GetEquipmentInfo()
		g.gameResult.resultMiners = g.GetInactiveMiners()

		return nil, g.gameResult
	}
}

/*
StopMiners
останавливает только работу шахтёров
*/
func (g *GameService) StopMiners() error {
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
	if err := g.ctx.Err(); err == nil {
		return GameNotEnd, nil
	} else {
		return nil, g.gameResult
	}
}

//==================================================

type GameResult struct {
	balance         domain.Coal
	startTime       time.Time
	endTime         time.Time
	durationTime    time.Duration
	endedAuto       bool
	resultEquipment map[domain.EquipmentType]bool
	resultMiners    map[domain.ID]domain.Miner
}

func InitGameResult() *GameResult {
	return &GameResult{}
}
