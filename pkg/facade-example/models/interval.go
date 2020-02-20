package models

type Interval struct {
	Min int
	Max int
}

func (i *Interval) Include(c int) bool {
	return c >= i.Min && c <= i.Max
}
