package placeburgerbar

type place = interface {
	Accept(v visitor) string
}

type sushiBar = interface {
	place
	BuySushi() string
}

type pizzeria = interface {
	place
	BuyPizza() string
}

type visitor = interface {
	VisitSushiBar(p sushiBar) string
	VisitPizzeria(p pizzeria) string
	VisitBurgerBar(p BurgerBar) string
}

//BurgerBar
type BurgerBar = interface {
	place
	BuyBurger() string
}

type burgerBar struct {
}


func (b *burgerBar) Accept(v visitor) string {
	return v.VisitBurgerBar(b)
}

func (b *burgerBar) BuyBurger() string {
	return "Buy burger..."
}

func NewBurgerBar() BurgerBar {
	return &burgerBar{}
}
