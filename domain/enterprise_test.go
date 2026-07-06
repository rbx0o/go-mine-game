package domain

import "testing"

/*
Test_InitEnterpriseReturnCorrectValues
тестирует функцию инициализации предприятия и проверяет, что
все поля корректно инициализированы
*/
func Test_InitEnterpriseReturnCorrectValues(t *testing.T) {
	ent := InitEnterprise()

	wantBalance := Coal(0)
	hasBalance := ent.Balance
	if wantBalance != hasBalance {
		t.Errorf("want balance: %v, has balance: %v", wantBalance, hasBalance)
	}

	wantPassiveIncome := Coal(1)
	hasPassiveIncome := ent.PassiveIncome
	if wantPassiveIncome != hasPassiveIncome {
		t.Errorf("want passive income: %v, has passive income: %v", wantPassiveIncome, hasPassiveIncome)
	}

	if ent.Ctx == nil {
		t.Errorf("ctx is nil")
	}

	if ent.CtxCancel == nil {
		t.Errorf("ctx cancel func is nil")
	}

	if ent.ActiveMiners == nil {
		t.Errorf("active miners map is nil")
	}

	if ent.InactiveMiners == nil {
		t.Errorf("inactive miners map is nil")
	}

	if ent.AllEquipment == nil {
		t.Errorf("all equipment map is nil")
	}

	if len(ent.AllEquipment) != 3 {
		t.Errorf("all equipment map len isn`t equal 3")
	}

	if ent.AllEquipment[PickaxeType] == nil {
		t.Errorf("pickaxe is nil")
	}

	if ent.AllEquipment[VentilationType] == nil {
		t.Errorf("ventilation is nil")
	}

	if ent.AllEquipment[TrolleysType] == nil {
		t.Errorf("trolleys is nil")
	}
}

/*
Test_InitEnterprisePointsToDifferentMap
тестирует функцию инициализации предприятия и проверяет, что
поля ActiveMiners и InactiveMiners указывают на разны map
*/
func Test_InitEnterprisePointsToDifferentMap(t *testing.T) {
	ent1 := InitEnterprise()
	ent2 := InitEnterprise()

	id, _ := NewID()

	ent1.ActiveMiners[id], _ = InitSmallMiner()
	ent2.InactiveMiners[id], _ = InitSmallMiner()

	if val, ok := ent1.InactiveMiners[id]; ok {
		t.Errorf("ent1 has inactive miner: %v", val)
	}
	if val, ok := ent2.ActiveMiners[id]; ok {
		t.Errorf("ent2 has active miner: %v", val)
	}
}
