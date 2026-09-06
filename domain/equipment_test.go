package domain

import "testing"

/*
Test_InitPickaxeReturnCorrectValues
тестирует функцию инициализации кирки и проверяет, что
все поля корректно инициализированы
*/
func Test_InitPickaxeReturnCorrectValues(t *testing.T) {
	eq := InitPickaxe()

	wantCost := Coal(3000)
	hasCost := eq.Cost()
	if wantCost != hasCost {
		t.Errorf("want cost: %v, hasCost: %v", wantCost, hasCost)
	}

	wantIsBought := false
	hasIsBought := eq.IsBought()
	if wantIsBought != hasIsBought {
		t.Errorf("want is bought: %v, has is bought: %v", wantIsBought, hasIsBought)
	}
}

/*
Test_InitVentilationReturnCorrectValues
тестирует функцию инициализации вентиляции и проверяет, что
все поля корректно инициализированы
*/
func Test_InitVentilationReturnCorrectValues(t *testing.T) {
	eq := InitVentilation()

	wantCost := Coal(15000)
	hasCost := eq.Cost()
	if wantCost != hasCost {
		t.Errorf("want cost: %v, hasCost: %v", wantCost, hasCost)
	}

	wantIsBought := false
	hasIsBought := eq.IsBought()
	if wantIsBought != hasIsBought {
		t.Errorf("want is bought: %v, has is bought: %v", wantIsBought, hasIsBought)
	}
}

/*
Test_InitTrolleysReturnCorrectValues
тестирует функцию инициализации вагонеток и проверяет, что
все поля корректно инициализированы
*/
func Test_InitTrolleysReturnCorrectValues(t *testing.T) {
	eq := InitTrolleys()

	wantCost := Coal(50000)
	hasCost := eq.Cost()
	if wantCost != hasCost {
		t.Errorf("want cost: %v, hasCost: %v", wantCost, hasCost)
	}

	wantIsBought := false
	hasIsBought := eq.IsBought()
	if wantIsBought != hasIsBought {
		t.Errorf("want is bought: %v, has is bought: %v", wantIsBought, hasIsBought)
	}
}

/*
Test_BuyEquipment
тестирует функцию покупки оборудования
*/
func Test_BuyEquipment(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		equipment Equipment
	}{
		{
			name:      "pickaxe",
			want:      true,
			equipment: InitPickaxe(),
		},
		{
			name:      "ventilation",
			want:      true,
			equipment: InitVentilation(),
		},
		{
			name:      "trolleys",
			want:      true,
			equipment: InitTrolleys(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.equipment.Buy()
			has := tt.equipment.IsBought()
			if has != tt.want {
				t.Errorf("want is bought: %v, has is bought: %v", tt.want, has)
			}
		})
	}
}

/*
Test_CostEquipment
тестирует функцию, возвращающую стоимость оборудования
*/
func Test_CostEquipment(t *testing.T) {
	tests := []struct {
		name      string
		wantCost  Coal
		equipment Equipment
	}{
		{
			name:      "pickaxe",
			wantCost:  3_000,
			equipment: InitPickaxe(),
		},
		{
			name:      "ventilation",
			wantCost:  15_000,
			equipment: InitVentilation(),
		},
		{
			name:      "trolleys",
			wantCost:  50_000,
			equipment: InitTrolleys(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasCost := tt.equipment.Cost()
			if hasCost != tt.wantCost {
				t.Errorf("want cost: %v, has cost: %v", tt.wantCost, hasCost)
			}
		})
	}
}

/*
Test_GetInfo
тестирует функцию получения информации
*/
func Test_GetInfo(t *testing.T) {
	tests := []struct {
		name      string
		wantName  string
		wantCost  Coal
		equipment Equipment
	}{
		{
			name:      "pickaxe",
			wantName:  "Кирка",
			wantCost:  3_000,
			equipment: InitPickaxe(),
		},
		{
			name:      "ventilation",
			wantName:  "Вентиляция",
			wantCost:  15_000,
			equipment: InitVentilation(),
		},
		{
			name:      "trolleys",
			wantName:  "Вагонетка",
			wantCost:  50_000,
			equipment: InitTrolleys(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := tt.equipment.GetInfo()
			hasName := info.Name
			hasCost := info.Cost
			if hasName != tt.wantName {
				t.Errorf("want name info: %v, has name info: %v", tt.wantName, hasName)
			}
			if hasCost != tt.wantCost {
				t.Errorf("want cost info: %v, has cost info: %v", tt.wantCost, hasCost)
			}
		})
	}
}
