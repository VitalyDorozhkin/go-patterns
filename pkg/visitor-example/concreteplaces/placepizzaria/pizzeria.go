package placepizzaria

type place = interface {
	Accept(v visitor) string
}

type sushiBar = interface {
	place
	BuySushi() string
}

type burgerBar = interface {
	place
	BuyBurger() string
}

type visitor interface {
	VisitSushiBar(p sushiBar) string
	VisitPizzeria(p Pizzeria) string
	VisitBurgerBar(p burgerBar) string
}

//Pizzeria
type Pizzeria = interface {
	place
	BuyPizza() string
}

type pizzeria struct {
}

func (b *pizzeria) Accept(v visitor) string {
	return v.VisitPizzeria(b)
}

func (b *pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

func NewPizzeria() Pizzeria {
	return &pizzeria{}
}
