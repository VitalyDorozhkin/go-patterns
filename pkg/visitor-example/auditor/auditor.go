package auditor

import (
	"fmt"

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

// auditor implements the Visitor interface.
type auditor struct {
	rating int
}

// VisitSushiBar implements visit to SushiBar.
func (v *auditor) VisitSushiBar(s placesushibar.SushiBar) string {
	s.SetRating(v.rating)
	return fmt.Sprintf("SushiBar was rated %d", v.rating)
}

// VisitPizzeria implements visit to Pizzeria.
func (v *auditor) VisitPizzeria(p placepizzaria.Pizzeria) string {
	p.SetRating(v.rating)
	return fmt.Sprintf("Pizzeria was rated %d", v.rating)
}

// VisitBurgerBar implements visit to BurgerBar.
func (v *auditor) VisitBurgerBar(b placeburgerbar.BurgerBar) string {
	b.SetRating(v.rating)
	return fmt.Sprintf("BurgerBar was rated %d", v.rating)
}

// NewAuditor initialize auditor
func NewAuditor(r int) Visitor {
	return &auditor{
		rating: r,
	}
}
