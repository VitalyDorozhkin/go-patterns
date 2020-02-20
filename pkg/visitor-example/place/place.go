package place

type sushiBar interface {
	Place
	BuySushi() string
}

type burgerBar interface {
	Place
	BuyBurger() string
}

type pizzeria interface {
	Place
	BuyPizza() string
}

type visitor interface {
	VisitSushiBar(p sushiBar) string
	VisitPizzeria(p pizzeria) string
	VisitBurgerBar(p burgerBar) string
}

// Place provides an interface for place that the visitor should visit.
type Place interface {
	Accept(v visitor) string
}
