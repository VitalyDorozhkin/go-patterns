package placesushibar

type place = interface {
	Accept(v visitor) string
}

type burgerBar =interface {
	place
	BuyBurger() string
}

type pizzeria = interface {
	place
	BuyPizza() string
}

type visitor = interface {
	VisitSushiBar(p SushiBar) string
	VisitPizzeria(p pizzeria) string
	VisitBurgerBar(p burgerBar) string
}
//SushiBar
type SushiBar interface {
	place
	BuySushi() string
}

type sushiBar struct {
}  


func (s *sushiBar) Accept(v visitor) string {
	return v.VisitSushiBar(s)
}

func (s *sushiBar) BuySushi() string {
	return "Buy sushi..."
}

func NewSushiBar() SushiBar {
	return &sushiBar{}
}
