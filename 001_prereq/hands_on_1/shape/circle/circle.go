package circle

import "math"

func GetInstance(radius float64) Circle {
	return Circle{radius}
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
