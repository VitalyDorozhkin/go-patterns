package visitor

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
)

type Visitor interface {
	VisitSushiBar(p placesushibar.SushiBar) string
	VisitPizzeria(p placepizzaria.Pizzeria) string
	VisitBurgerBar(p placeburgerbar.BurgerBar) string
}
