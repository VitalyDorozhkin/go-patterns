package placeburgerbar

type visitor interface {
	VisitBurgerBar(p BurgerBar) string
}

// BurgerBar ...
type BurgerBar interface {
	Accept(v visitor) string
	BuyBurger() string
	SetRating(r int)
}

type burgerBar struct {
	rating int
}

// SetRating ...
func (b *burgerBar) SetRating(r int) {
	b.rating = r
}

// Accept ...
func (b *burgerBar) Accept(v visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger ...
func (b *burgerBar) BuyBurger() string {
	return "Buy burger"
}

// NewBurgerBar initialize burgerBar
func NewBurgerBar() BurgerBar {
	return &burgerBar{}
}
