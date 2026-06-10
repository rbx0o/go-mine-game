package service

import "github.com/rbx0o/go-mine-game/domain"

/*==================================================

Для удобства один пакет разбит на несколько файлов.

Здесь описаны функции, связанные с оборудованием.

==================================================*/

func (g *GameService) BuyEquipment(equipment domain.EquipmentType) error {
	defer g.enterprise.Mtx.Unlock()
	g.enterprise.Mtx.Lock()

	switch equipment {
	case domain.PickaxeType:
		return nil
	case domain.VentilationType:
		return nil
	case domain.TrolleysType:
		return nil
	default:
		return EquipmentTypeNotFound
	}
}
