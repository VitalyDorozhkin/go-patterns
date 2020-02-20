package moll

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


type Moll = interface {
	Add(p place)
	Accept(v visitor) string
}

type moll struct {
	places []place
}


// Add appends Place to the collection.
func (c *moll) Add(p place) {
	c.places = append(c.places, p)
}

// Accept implements a visit to all places in the city.
func (c *moll) Accept(v visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

func NewMoll() Moll {
	return &moll{}
}
