package domain

// Файл в котором будет описано предприятие и базовая логика

import (
	"context"
	"sync"
)

/*
==================================================

Есть структура, которая представляет из себя предприятие
Структура хранит в себе информацию о:
- текущем балансе
- пассивном доходе
- контексте, к которому будут привязаны все шахтёры
- списке активных шахтёров
- списке неактивных шахтёров

==================================================
*/
type Enterprise struct {
	Balance       Coal // Баланс угля
	PassiveIncome Coal // Пассивный доход угля в секунду

	Ctx       context.Context    // Контекст выполнения горутин
	CtxCancel context.CancelFunc // Функция завершения контекста

	ActiveMiners   map[int]Miner // Работающие в данный момент шахтёры
	InactiveMiners map[int]Miner // Шахтёры завершившие работу

	AllEquipment map[string]Equipment // Всё оборудование на предприятии

	Mtx sync.Mutex
}

func InitEnterprise() *Enterprise {
	tempCtx, tempCtxCancel := context.WithCancel(context.Background())

	tempAllEquipment := make(map[string]Equipment, 3)
	tempAllEquipment["pickaxe"] = InitPickaxe()
	tempAllEquipment["ventilation"] = InitVentilation()
	tempAllEquipment["trolleys"] = InitTrolleys()

	return &Enterprise{
		Balance:       0,
		PassiveIncome: 1,

		Ctx:       tempCtx,
		CtxCancel: tempCtxCancel,

		ActiveMiners:   make(map[int]Miner),
		InactiveMiners: make(map[int]Miner),

		AllEquipment: tempAllEquipment,
	}
}
