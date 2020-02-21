package placesushibar

type visitor interface {
	VisitSushiBar(p SushiBar) string
}

// SushiBar ...
type SushiBar interface {
	Accept(v visitor) string
	BuySushi() string
	SetRating(r int)
}

type sushiBar struct {
	rating int
}

// SetRating ...
func (s *sushiBar) SetRating(r int) {
	s.rating = r
}

// Accept ...
func (s *sushiBar) Accept(v visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi ...
func (s *sushiBar) BuySushi() string {
	return "Buy sushi"
}

// NewSushiBar initialize sushiBar
func NewSushiBar() SushiBar {
	return &sushiBar{}
}
