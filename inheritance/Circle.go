package main

import "math"

type Circle struct {
	Colorful
	radius float64
}

func (k *Circle) Area() float64 {
	return math.Pi * k.radius * k.radius
}
