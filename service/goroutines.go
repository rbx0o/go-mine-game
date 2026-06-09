package service

import (
	"context"
	"time"
)

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны отдельные функции, которые будут запускаться
в виде горутин.

==================================================*/

/*
StartPassiveIncome запускает пассивный заработок угля на
предприятии в горутине
*/
func (g *GameService) StartPassiveIncome(ctx context.Context) {

	go func() {
		select {
		case <-ctx.Done():
			return

		case <-time.After(1 * time.Second):
			g.enterprise.Mtx.Lock()
			g.enterprise.Balance += g.enterprise.PassiveIncome
			g.enterprise.Mtx.Unlock()
		}
	}()

}
