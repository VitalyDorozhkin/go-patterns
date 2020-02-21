package abstractvisitor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/auditor"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/customer"
)

const rate = 5

func TestVisitor(t *testing.T) {

	visitor1 := customer.NewCustomer()
	auditor2 := auditor.NewAuditor(rate)

	pizzeria := placepizzaria.NewPizzeria()
	burgerBar := placeburgerbar.NewBurgerBar()
	sushiBar := placesushibar.NewSushiBar()

	t.Run("Customer in SushiBar", func(t *testing.T) {
		assert.Equal(t, sushiBar.Accept(visitor1), "Buy sushi")
	})
	t.Run("Customer in BurgerBar", func(t *testing.T) {
		assert.Equal(t, burgerBar.Accept(visitor1), "Buy burger")
	})
	t.Run("Customer in Pizzeria", func(t *testing.T) {
		assert.Equal(t, pizzeria.Accept(visitor1), "Buy pizza")
	})
	t.Run("Auditor in SushiBar", func(t *testing.T) {
		assert.Equal(t, sushiBar.Accept(auditor2), fmt.Sprintf("SushiBar was rated %d", rate))
	})
	t.Run("Auditor in BurgerBar", func(t *testing.T) {
		assert.Equal(t, burgerBar.Accept(auditor2), fmt.Sprintf("BurgerBar was rated %d", rate))
	})
	t.Run("Auditor in Pizzeria", func(t *testing.T) {
		assert.Equal(t, pizzeria.Accept(auditor2), fmt.Sprintf("Pizzeria was rated %d", rate))
	})
}
