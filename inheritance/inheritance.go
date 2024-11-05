package main

import (
	"fmt"
	"reflect"
)

func main() {
	forms := []any{
		&Circle{
			Colorful: Colorful{"black"},
			radius:   5.0,
		},
		&Rectangle{
			Colorful: Colorful{"red"},
			width:    2.0,
			height:   3.0,
		},
	}

	for _, a := range forms {
		t := reflect.TypeOf(a)
		fmt.Printf("%s:\n", t.Elem().Name())

		if f, ok := a.(Form); ok {
			fmt.Printf("Color: %s\n", f.Color())
			fmt.Printf("Area: %.2f\n", f.Area())
		}
		switch v := a.(type) {
		case HasCorners:
			fmt.Printf("Corners: %d\n", v.CntCorners())
		default:
			fmt.Printf("No Corners\n")
		}
	}
}
