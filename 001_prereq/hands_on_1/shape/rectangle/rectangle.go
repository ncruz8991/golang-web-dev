package rectangle

func GetInstance(width, height float64) Rectangle {
	return Rectangle{width: width, height: height}
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}