package main

type Colorful struct {
	color string
}

func (f *Colorful) Color() string {
	return f.color
}
