package placepizzaria

type visitor interface {
	VisitPizzeria(p Pizzeria) string
}

//Pizzeria
type Pizzeria interface {
	Accept(v visitor) string
	BuyPizza() string
	SetRating(r int)
}

type pizzeria struct {
	rating int
}

// SetRating ...
func (p *pizzeria) SetRating(r int) {
	p.rating = r
}

// Accept  ...
func (p *pizzeria) Accept(v visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza ...
func (p *pizzeria) BuyPizza() string {
	return "Buy pizza"
}

// NewPizzeria initialize pizzeria
func NewPizzeria() Pizzeria {
	return &pizzeria{}
}
