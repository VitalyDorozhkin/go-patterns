package visitor

import (
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concretevisitors/customer"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/place"
	"testing"
)

func TestVisitor(t *testing.T) {

	expect := "Buy sushi...Buy pizza...Buy burger..."

	var moll []place.Place
	v := customer.NewPeople()
	moll = append(moll, placesushibar.NewSushiBar())
	placepizzaria.NewPizzeria()
	placeburgerbar.NewBurgerBar()

	var result string
	for _, p := range moll {
		result += p.Accept(v)
	}


	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
