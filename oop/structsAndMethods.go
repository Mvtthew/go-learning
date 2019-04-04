package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) Area() int {
	return r.height * r.width
}

type triangle struct {
	a int
	h int
}

func (t triangle) Area() int {
	return (t.a * t.h) / 2
}

func main() {

	square := rect{
		width:  5,
		height: 10,
	}

	fmt.Println(square)
	fmt.Println(square.Area())

	tr := triangle{
		a: 5,
		h: 10,
	}

	fmt.Println()

	fmt.Println(tr)
	fmt.Println(tr.Area())

}
