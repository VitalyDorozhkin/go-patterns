package customer

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

type pizzeria = interface {
	place
	BuyPizza() string
}

type visitor = interface {
	VisitSushiBar(p sushiBar) string
	VisitPizzeria(p pizzeria) string
	VisitBurgerBar(p burgerBar) string
}

type People = interface {
	visitor
}

// People implements the Visitor interface.
type people struct {
}

// VisitSushiBar implements visit to SushiBar.
func (v *people) VisitSushiBar(s sushiBar) string {
	return s.BuySushi()
}

// VisitPizzeria implements visit to Pizzeria.
func (v *people) VisitPizzeria(p pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *people) VisitBurgerBar(b burgerBar) string {
	return b.BuyBurger()
}

func NewPeople() People {
	return &people{}
}
