package main

import (
	"fmt"

	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/auditor"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placeburgerbar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placepizzaria"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/concreteplaces/placesushibar"
	"github.com/VitalyDorozhkin/go-patterns/pkg/visitor-example/customer"
)

func main() {
	c := customer.NewCustomer()
	a := auditor.NewAuditor(5)

	p := placepizzaria.NewPizzeria()
	b := placeburgerbar.NewBurgerBar()
	s := placesushibar.NewSushiBar()

	fmt.Println(s.Accept(c))
	fmt.Println(p.Accept(c))
	fmt.Println(b.Accept(c))

	fmt.Println(s.Accept(a))
	fmt.Println(p.Accept(a))
	fmt.Println(b.Accept(a))
}
