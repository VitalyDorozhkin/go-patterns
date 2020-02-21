package customer

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
)

// Visitor ...
type Visitor interface {
	VisitSushiBar(p placesushibar.SushiBar) string
	VisitPizzeria(p placepizzaria.Pizzeria) string
	VisitBurgerBar(p placeburgerbar.BurgerBar) string
}

// customer implements the Visitor interface.
type customer struct {
}

// VisitSushiBar implements visit to SushiBar.
func (v *customer) VisitSushiBar(s placesushibar.SushiBar) string {
	return s.BuySushi()
}

// VisitPizzeria implements visit to Pizzeria.
func (v *customer) VisitPizzeria(p placepizzaria.Pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *customer) VisitBurgerBar(b placeburgerbar.BurgerBar) string {
	return b.BuyBurger()
}

// NewAuditor initialize auditor
func NewCustomer() Visitor {
	return &customer{}
}
