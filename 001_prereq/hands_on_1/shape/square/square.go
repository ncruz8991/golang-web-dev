package square

import "github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_1/shape/rectangle"

func GetInstance(side float64) Square {
	return Square{rectangle.GetInstance(side, side)}
}

type Square struct {
	r rectangle.Rectangle
}

func (s Square) Area() float64 {
	return s.r.Area()
}