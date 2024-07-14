package interfaces

import "fmt"

type Shape interface {
	getArea() float64
	printArea()
}

type triangle struct {
	base float64
	height float64
}

type rectangle struct {
	width float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) printArea() {
	fmt.Println(r.getArea())
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}