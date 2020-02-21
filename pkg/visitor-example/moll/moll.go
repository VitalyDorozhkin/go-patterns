package moll

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
)

type place = interface {
	Accept(v visitor) string
}

type visitor = interface {
	VisitSushiBar(p placesushibar.SushiBar) string
	VisitPizzeria(p placepizzaria.Pizzeria) string
	VisitBurgerBar(p placeburgerbar.BurgerBar) string
}


type Moll interface {
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
