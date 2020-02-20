package visitor

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/moll"
	"testing"
)

func TestVisitor(t *testing.T) {

	expect := "Buy sushi...Buy pizza...Buy burger..."

	city := moll.NewMoll()

	city.Add(placesushibar.NewSushiBar())
	city.Add(placepizzaria.NewPizzeria())
	city.Add(placeburgerbar.NewBurgerBar())

	result := city.Accept(placesushibar.NewVisitor())

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
