package domain

// Файл в котором будет описано оборудования для предприятия и базовая логика

/*==================================================

Оборудование покупается не для повышения производительности,
а для того, чтобы пройти игру.
У нас есть 3 оборудования, которые необходимо купить:
- Кирки - 3000 угля
- Вентиляция - 15000 угля
- Вагонетки - 50000
Лучше сделать интерфейс и уже на его основе 3 структуры.

==================================================*/

type Equipment interface {
	Buy()
	Cost() Coal
	IsBought() bool
}

//==================================================

type Pickaxe struct {
	cost     Coal // Стоимость - 3000
	isBought bool // Куплено/не куплено
}

func InitPickaxe() *Pickaxe {
	return &Pickaxe{
		cost:     3000,
		isBought: false,
	}
}

func (p *Pickaxe) Buy() {
	p.isBought = true
}

func (p *Pickaxe) Cost() Coal {
	return p.cost
}

func (p *Pickaxe) IsBought() bool {
	return p.isBought
}

//==================================================

type Ventilation struct {
	cost     Coal // Стоимость - 15000
	isBought bool // Куплено/не куплено
}

func InitVentilation() *Ventilation {
	return &Ventilation{
		cost:     15000,
		isBought: false,
	}
}

func (v *Ventilation) Buy() {
	v.isBought = true
}

func (v *Ventilation) Cost() Coal {
	return v.cost
}

func (v *Ventilation) IsBought() bool {
	return v.isBought
}

//==================================================

type Trolleys struct {
	cost     Coal // Стоимость - 50000
	isBought bool // Куплено/не куплено
}

func InitTrolleys() *Trolleys {
	return &Trolleys{
		cost:     50000,
		isBought: false,
	}
}

func (t *Trolleys) Buy() {
	t.isBought = true
}

func (t *Trolleys) Cost() Coal {
	return t.cost
}

func (t *Trolleys) IsBought() bool {
	return t.isBought
}
