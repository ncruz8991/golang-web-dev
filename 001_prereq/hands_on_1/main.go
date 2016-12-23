package main

import (
	"fmt"
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_1/shape/square"
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_1/shape/circle"
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_1/shape/rectangle"
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_1/shape"
)

func main() {
	shapeSlice := []shape.Shape{
		square.GetInstance(2),
		circle.GetInstance(1),
		rectangle.GetInstance(1, 3),
	}

	for _, s := range shapeSlice {
		fmt.Println(s.Area())
	}
}
