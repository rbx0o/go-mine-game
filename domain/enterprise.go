package domain

// Файл в котором будет описано предприятие и базовая логика

import "context"

/*==================================================

Есть структура, которая представляет из себя предприятие
Структура хранит в себе информацию о:
- текущем балансе
- пассивном доходе
- контексте, к которому будут привязаны все шахтёры
- списке активных шахтёров
- списке неактивных шахтёров

==================================================*/
type Enterprise struct {
	balance       Coal // Баланс угля
	passiveIncome Coal // Пассивный доход угля в секунду

	ctx       context.Context    // Контекст выполнения горутин
	ctxCancel context.CancelFunc // Функция завершения контекста

	activeMiners   map[int]Miner // Работающие в данный момент шахтёры
	inactiveMiners map[int]Miner // Шахтёры завершившие работу
}

func InitEnterprise() *Enterprise {
	tempCtx, tempCtxCancel := context.WithCancel(context.Background())

	return &Enterprise{
		balance:       0,
		passiveIncome: 1,

		ctx:       tempCtx,
		ctxCancel: tempCtxCancel,

		activeMiners:   make(map[int]Miner),
		inactiveMiners: make(map[int]Miner),
	}
}
