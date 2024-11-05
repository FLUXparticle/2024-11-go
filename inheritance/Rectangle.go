package main

type Rectangle struct {
	Colorful
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) CntCorners() int {
	return 4
}
