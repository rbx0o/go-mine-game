package service

// Здесь будет описан весь сервис и процессы того, что пользователь может делать с игрой

import (
	"context"

	"github.com/rbx0o/go-mine-game/domain"
)

/*==================================================

Есть структура, которая представляет из себя саму игру.
Данную игру можно инициализировать (при запуске игры),
запускать, ставить на паузу и завершать.

Сама игра содержит в себе предприятие.

==================================================*/

type GameService struct {
	enterprise *domain.Enterprise

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func InitGameService() *GameService {
	tempCtx, tempCtxCancel := context.WithCancel(context.Background())

	return &GameService{
		enterprise: domain.InitEnterprise(),

		ctx:       tempCtx,
		ctxCancel: tempCtxCancel,
	}
}

func (g *GameService) Start() error {
	g.StartPassiveIncome(g.ctx)

	return nil
}

func (g *GameService) Stop() error {

	g.enterprise.CtxCancel()
	g.ctxCancel()

	return nil
}
