package visitor

type place interface {
	Accept(v Visitor) string
}

type sushiBar interface {
	place
	BuySushi() string
}

type burgerBar interface {
	place
	BuyBurger() string
}

type pizzeria interface {
	place
	BuyPizza() string
}


type Visitor interface {
	VisitSushiBar(p sushiBar) string
	VisitPizzeria(p pizzeria) string
	VisitBurgerBar(p burgerBar) string
}
